/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

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
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core/cdata"
	"github.com/intelsdi-x/snap/core/ctypes"
)

type HaproxyPluginSuite struct {
	suite.Suite
	cfg                    plugin.PluginConfigType
	dataNode               *cdata.ConfigDataNode
	ncInfoData, ncStatData []string
	MockSocket             *MockSocket
}

func (hps *HaproxyPluginSuite) SetupSuite() {
	hps.ncInfoData = []string{
		"Maxsock: 4444",
		"CurrConns: 4",
		"CumReq: 44",
		"Tasks: 7",
	}

	hps.ncStatData = []string{
		"# pxname,svname,qcur,qmax,scur,smax,slim,stot,bin,bout,dreq,dresp,ereq,econ,eresp,wretr,wredis,status,weight,act,bck,chkfail,chkdown,lastchg,downtime,qlimit,pid,iid,sid,throttle,lbtot,tracked,type,rate,rate_lim,rate_max,check_status,check_code,check_duration,hrsp_1xx,hrsp_2xx,hrsp_3xx,hrsp_4xx,hrsp_5xx,hrsp_other,hanafail,req_rate,req_rate_max,req_tot,cli_abrt,srv_abrt,comp_in,comp_out,comp_byp,comp_rsp,lastsess,last_chk,last_agt,qtime,ctime,rtime,ttime,",
		"LB,FRONTEND,,,0,0,2000,0,0,0,0,0,0,,,,,OPEN,,,,,,,,,1,2,0,,,,0,0,0,0,,,,,,,,,,,0,0,0,,,0,0,0,0,,,,,,,,",
		"LB,Server01,0,0,0,0,,0,0,0,,0,,0,0,0,0,UPe,1,1,0,0,0,3235,0,,1,3,1,,0,,2,0,,0,L4OK,,0,,,,,,,0,,,,0,0,,,,,-1,,,0,0,0,0,",
	}

	hps.dataNode = cdata.NewNode()
	hps.dataNode.AddItem("socket", ctypes.ConfigValueStr{Value: "fake.haproxy.sock"})
	hps.cfg = plugin.PluginConfigType{ConfigDataNode: hps.dataNode}

	hps.MockSocket = &MockSocket{}
	hps.MockSocket.On("Read", mock.AnythingOfType("string"), ncinfo).Return(hps.ncInfoData, nil)
	hps.MockSocket.On("Read", mock.AnythingOfType("string"), ncstat).Return(hps.ncStatData, nil)
}

func (hps *HaproxyPluginSuite) TestGetMetricTypes() {
	Convey("Given haproxy plugin is initialized", hps.T(), func() {
		haprx := haproxyPlugin{host: "localhost", socket: hps.MockSocket}

		Convey("When one wants to get list of available meterics", func() {
			mts, err := haprx.GetMetricTypes(hps.cfg)

			Convey("Then no error should be reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then proper list of available metrics should be returned", func() {
				So(len(mts), ShouldEqual, 124)

				namespace := []string{}
				for _, m := range mts {
					namespace = append(namespace, strings.Join(m.Namespace(), "/"))
				}

				So(namespace, ShouldContain, "intel/linux/haproxy/stat/FRONTEND/LB/qcur")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/FRONTEND/LB/qmax")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/FRONTEND/LB/rtime")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/FRONTEND/LB/ttime")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/Server01/LB/qcur")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/Server01/LB/qmax")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/Server01/LB/rtime")
				So(namespace, ShouldContain, "intel/linux/haproxy/stat/Server01/LB/ttime")
				So(namespace, ShouldContain, "intel/linux/haproxy/info/Maxsock")
				So(namespace, ShouldContain, "intel/linux/haproxy/info/Tasks")
				So(namespace, ShouldContain, "intel/linux/haproxy/info/CurrConns")
				So(namespace, ShouldContain, "intel/linux/haproxy/info/CumReq")
			})
		})
	})
}

func (hps *HaproxyPluginSuite) TestCollectMetrics() {
	Convey("Given haproxy plugin is initialized", hps.T(), func() {
		haprx := haproxyPlugin{host: "localhost", socket: hps.MockSocket}

		Convey("When one wants to get values for requested meterics", func() {
			mts := []plugin.PluginMetricType{
				plugin.PluginMetricType{
					Namespace_: []string{VENDOR, OS, PLUGIN, ncstat, "FRONTEND", "LB", "slim"},
					Config_:    hps.dataNode,
				},
				plugin.PluginMetricType{
					Namespace_: []string{VENDOR, OS, PLUGIN, ncstat, "FRONTEND", "LB", "status"},
					Config_:    hps.dataNode,
				},
				plugin.PluginMetricType{
					Namespace_: []string{VENDOR, OS, PLUGIN, ncstat, "Server01", "LB", "lastchg"},
					Config_:    hps.dataNode,
				},
				plugin.PluginMetricType{
					Namespace_: []string{VENDOR, OS, PLUGIN, ncstat, "Server01", "LB", "ttime"},
					Config_:    hps.dataNode,
				},
				plugin.PluginMetricType{
					Namespace_: []string{VENDOR, OS, PLUGIN, ncinfo, "Maxsock"},
					Config_:    hps.dataNode,
				},
				plugin.PluginMetricType{
					Namespace_: []string{VENDOR, OS, PLUGIN, ncinfo, "CurrConns"},
					Config_:    hps.dataNode,
				},
			}

			metrics, err := haprx.CollectMetrics(mts)

			Convey("Then no error should be reported", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then proper list of available metrics should be returned", func() {
				So(len(metrics), ShouldEqual, len(mts))

				stats := map[string]interface{}{}
				for _, m := range metrics {
					stat := strings.Join(m.Namespace()[3:], "/")
					stats[stat] = m.Data()
				}

				val, ok := stats["stat/FRONTEND/LB/slim"].(string)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, "2000")

				val, ok = stats["stat/FRONTEND/LB/status"].(string)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, "OPEN")

				val, ok = stats["stat/Server01/LB/lastchg"].(string)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, "3235")

				val, ok = stats["stat/Server01/LB/ttime"].(string)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, "0")

				val, ok = stats["info/Maxsock"].(string)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, "4444")

				val, ok = stats["info/CurrConns"].(string)
				So(ok, ShouldBeTrue)
				So(val, ShouldEqual, "4")
			})
		})
	})
}

func TestHaproxyPluginSuite(t *testing.T) {
	suite.Run(t, &HaproxyPluginSuite{})
}

type MockSocket struct {
	mock.Mock
}

func (ms *MockSocket) Read(name string, mode string) ([]string, error) {
	ret := ms.Mock.Called(name, mode)
	return ret.Get(0).([]string), ret.Error(1)
}
