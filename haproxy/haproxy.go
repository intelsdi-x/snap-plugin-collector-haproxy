// +build linux

/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015-2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package haproxy

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/snap-plugin-utilities/config"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core/serror"
)

const (
	// VENDOR namespace part
	VENDOR = "intel"
	// PLUGIN name namespace part
	PLUGIN = "haproxy"
	// VERSION of haproxy plugin
	VERSION = 2

	ncinfo    = "info"
	ncstat    = "stat"
	ncinfosep = ": "
	ncstatsep = ","
)

var (
	ncModes             = []string{ncinfo, ncstat}
	ncinfoStringMetrics = map[string]bool{"Name": true, "Release_date": true, "Uptime": true, "Version": true, "description": true, "node": true}
	ncstatStringMetrics = map[string]bool{"status": true, "check_status": true, "hanafail": true, "last_chk": true, "last_agt": true}
)

// GetMetricTypes returns list of available metric types
// It returns error in case retrieval was not successful
func (ha *haproxyPlugin) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	metricTypes := []plugin.PluginMetricType{}
	sckt, err := config.GetConfigItem(cfg, "socket")
	if err != nil {
		return nil, err
	}

	for _, ncMode := range ncModes {
		data, err := ha.socket.Read(sckt.(string), ncMode)
		if err != nil {
			return nil, err
		}

		switch ncMode {
		case ncinfo:
			stats, err := parseInfo(data, ncinfosep)
			if err != nil {
				return nil, err
			}
			for stat := range stats {
				namespace := []string{VENDOR, PLUGIN, ncinfo, stat}
				metricTypes = append(metricTypes, plugin.PluginMetricType{Namespace_: namespace})
			}
		case ncstat:
			all, err := parseStats(data, ncstatsep)
			if err != nil {
				return nil, err
			}
			for _, stats := range all {
				svname, ok := stats["svname"]
				if !ok {
					return nil, fmt.Errorf("SVNAME not available in statistics")
				}

				pxname, ok := stats["pxname"]
				if !ok {
					return nil, fmt.Errorf("PXNAME not available in statistics")
				}

				// svname and pxname will be included in namespace for each stat
				delete(stats, "svname")
				delete(stats, "pxname")

				for stat := range stats {
					namespace := []string{VENDOR, PLUGIN, ncstat, svname, pxname, stat}
					metricTypes = append(metricTypes, plugin.PluginMetricType{Namespace_: namespace})
				}
			}
		}
	}

	return metricTypes, nil
}

// CollectMetrics returns list of requested metric values
// It returns error in case retrieval was not successful
func (ha *haproxyPlugin) CollectMetrics(metricTypes []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	metrics := []plugin.PluginMetricType{}
	sckt, err := config.GetConfigItem(metricTypes[0], "socket")
	if err != nil {
		return nil, err
	}

	ncInfoData, err := ha.socket.Read(sckt.(string), ncinfo)
	if err != nil {
		return nil, err
	}

	ncStatData, err := ha.socket.Read(sckt.(string), ncstat)
	if err != nil {
		return nil, err
	}

	for _, metricType := range metricTypes {
		namespace := metricType.Namespace()
		if len(namespace) < 4 {
			return nil, fmt.Errorf("Namespace length is incorrect %d", len(namespace))
		}
		mode := namespace[2]
		switch mode {
		case ncinfo:
			stats, err := parseInfo(ncInfoData, ncinfosep)
			if err != nil {
				return nil, err
			}
			stat := namespace[3]
			val, ok := stats[stat]
			if !ok {
				return nil, fmt.Errorf("Requested metric is not found {%s}", strings.Join(namespace, "/"))
			}

			valConverted := setMetricType(namespace, val, ncinfoStringMetrics)
			metrics = append(metrics, plugin.PluginMetricType{
				Namespace_: []string{VENDOR, PLUGIN, ncinfo, stat},
				Data_:      valConverted,
				Version_:   VERSION,
				Timestamp_: time.Now(),
				Source_:    ha.host,
			})

		case ncstat:
			svname := namespace[3]
			pxname := namespace[4]

			all, err := parseStats(ncStatData, ncstatsep)
			if err != nil {
				return nil, err
			}

			for _, stats := range all {
				stat := namespace[5]
				sv, _ := stats["svname"]
				px, _ := stats["pxname"]
				val, ok := stats[stat]
				if sv == svname && px == pxname && ok {
					valConverted := setMetricType(namespace, val, ncstatStringMetrics)
					metrics = append(metrics, plugin.PluginMetricType{
						Namespace_: namespace,
						Data_:      valConverted,
						Version_:   VERSION,
						Timestamp_: time.Now(),
						Source_:    ha.host,
					})
					break
				}
			}
		}
	}
	return metrics, nil
}

// GetConfigPolicy returns config policy
// It returns error in case retrieval was not successful
func (ha *haproxyPlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	return cpolicy.New(), nil
}

// New creates instance of haproxy plugin
func New() *haproxyPlugin {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	ha := &haproxyPlugin{host: host, socket: &uSocket{}}

	return ha
}

type haproxyPlugin struct {
	host   string
	socket socketer
}

type socketer interface {
	Read(socket string, mode string) ([]string, error)
}

type uSocket struct{}

// Read communicates with socket to retrieve haproxy statistics from it.
// socketName is path to Unix-Domain socket
// mode can be "info" or "stat" depending on type of statistics needed
// It returns list of coma-separated entries or error in case of failure
func (sock *uSocket) Read(socketName string, mode string) ([]string, error) {
	var err error
	var reader io.ReadCloser
	cmd1 := exec.Command("echo", "show", mode)
	cmd2 := exec.Command("nc", "-U", socketName)

	cmd2.Stdin, err = cmd1.StdoutPipe()

	if err != nil {
		return nil, err
	}

	reader, err = cmd2.StdoutPipe()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(reader)

	if err = cmd1.Start(); err != nil {
		return nil, err
	}

	if err = cmd2.Start(); err != nil {
		return nil, err
	}

	done := make(chan bool)
	csv := []string{}

	go func() {
		for scanner.Scan() {
			csv = append(csv, scanner.Text())
		}
		done <- true
	}()

	<-done

	if err = cmd1.Wait(); err != nil {
		return nil, err
	}
	if err = cmd2.Wait(); err != nil {
		return nil, err
	}

	return csv, nil
}

func parseStats(csv []string, sep string) ([]map[string]string, error) {
	if len(csv) < 2 {
		return nil, fmt.Errorf("Return stats seems to short len(csv) = %v", len(csv))
	}

	rows := make([]map[string]string, 0)
	header := strings.Split(strings.Trim(csv[0], "# "), sep)

	for _, line := range csv[1:] {
		if line == "" {
			continue
		}

		stats := map[string]string{}
		data := strings.Split(line, sep)

		if len(data) != len(header) {
			return nil, fmt.Errorf("Wrong format of data! %v != %v\n", len(header), len(data))
		}

		for i := 0; i < len(header); i++ {
			name := header[i]
			if name == "" {
				continue
			}
			stats[name] = data[i]
		}

		rows = append(rows, stats)
	}
	return rows, nil
}

func parseInfo(data []string, sep string) (map[string]string, error) {
	stats := map[string]string{}
	for _, d := range data {
		if d == "" {
			continue
		}

		stat := strings.Split(d, sep)

		if len(stat) != 2 {
			return nil, fmt.Errorf("Wrong format of input data len() = %d\n", len(stat))
		}

		name := stat[0]
		stats[name] = stat[1]
	}

	return stats, nil
}

func setMetricType(nc []string, val interface{}, stringMetricsTab map[string]bool) interface{} {
	if stringMetricsTab[nc[len(nc)-1]] {
		return val
	}
	parsedVal, err := strconv.ParseInt(val.(string), 10, 64)
	if err != nil {
		f := map[string]interface{}{
			"namespace": strings.Join(nc, "/"),
			"val":       val,
			"parsedVal": parsedVal,
		}
		se := serror.New(err, f)
		log.WithFields(se.Fields()).Warn("Cannot parse metric value to number, metric value saved as -1, ", se.String())
		parsedVal = -1
	}
	return parsedVal
}
