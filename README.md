# snap collector plugin - haproxy
This plugin collects metrics from running HAProxy by reading statistics and informations from Unix socket.

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [golang 1.4+](https://golang.org/dl/)
* [HAProxy](http://haproxy.org) installed and running with 'stats socket' configuration

### Operating systems
All OSs currently supported by snap:
* Linux/amd64

### Installation
#### Download haproxy plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-haproxy  
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-haproxy.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/intel/linux/haproxy/info/Nbproc |
/intel/linux/haproxy/info/Hard_maxconn |
/intel/linux/haproxy/info/SessRateLimit |
/intel/linux/haproxy/info/CumReq |
/intel/linux/haproxy/info/Process_num |
/intel/linux/haproxy/info/Uptime |
/intel/linux/haproxy/info/CumConns |
/intel/linux/haproxy/info/MaxSessRate |
/intel/linux/haproxy/info/Version |
/intel/linux/haproxy/info/CurrConns |
/intel/linux/haproxy/info/MaxConnRate |
/intel/linux/haproxy/info/CompressBpsIn | 
/intel/linux/haproxy/info/Run_queue |
/intel/linux/haproxy/info/Uptime_sec |
/intel/linux/haproxy/info/Maxsock |
/intel/linux/haproxy/info/ConnRateLimit |
/intel/linux/haproxy/info/PipesFree |
/intel/linux/haproxy/info/Memmax_MB |
/intel/linux/haproxy/info/Ulimit-n |
/intel/linux/haproxy/info/Maxpipes | 
/intel/linux/haproxy/info/CompressBpsOut |
/intel/linux/haproxy/info/CompressBpsRateLim |
/intel/linux/haproxy/info/Idle_pct |
/intel/linux/haproxy/info/node |
/intel/linux/haproxy/info/Release_date |
/intel/linux/haproxy/info/ConnRate |
/intel/linux/haproxy/info/SessRate |
/intel/linux/haproxy/info/PipesUsed |
/intel/linux/haproxy/info/Tasks |
/intel/linux/haproxy/info/description |
/intel/linux/haproxy/info/Name |
/intel/linux/haproxy/info/Pid |
/intel/linux/haproxy/info/Maxconn |
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/qcur | current queued requests. For the backend this reports the number queued without a server assigned.
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/qmax | max value of qcur
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/scur | current sessions
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/smax | max sessions
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/slim | configured session limit
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/stot | cumulative number of connections
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/bin | bytes in
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/bout | bytes out
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/dreq | requests denied because of security concerns.
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/dresp | responses denied because of security concerns.
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/ereq | request errors
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/econ | number of requests that encountered an error trying to  connect to a backend server
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/eresp | response errors
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/wretr | number of times a connection to a server was retried.
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/wredis | number of times a request was redispatched to another  server
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/status | status (UP/DOWN/NOLB/MAINT/MAINT(via)...)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/weight | total weight (backend), server weight (server)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/act | number of active servers (backend), server is active (server)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/bck | number of backup servers (backend), server is backup (server)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/chkfail | number of failed checks
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/chkdown | number of UP->DOWN transitions
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/lastchg | number of seconds since the last UP<->DOWN transition
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/downtime | total downtime (in seconds)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/qlimit | configured maxqueue for the server, or nothing in the  value is 0 (default, meaning no limit)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/pid | process id (0 for first instance, 1 for second, ...)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/iid | unique proxy id
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/sid | server id (unique inside a proxy)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/throttle | current throttle percentage for the server, when slowstart is active, or no value if not in slowstart.
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/lbtot | total number of times a server was selected, either for new sessions, or when re-dispatching
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/tracked | id of proxy/server if tracking is enabled.
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/type | (0=frontend, 1=backend, 2=server, 3=socket/listener)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/rate | number of sessions per second over last elapsed second
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/rate_lim | configured limit on new sessions per second
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/rate_max | max number of new sessions per second
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/check_status | status of last health check
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/check_code | layer5-7 code, if available
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/check_duration | time in ms took to finish last health check
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_1xx | http responses with 1xx code
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_2xx | http responses with 2xx code
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_3xx | http responses with 3xx code
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_4xx | http responses with 4xx code
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_5xx | http responses with 5xx code
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_other | http responses with other codes (protocol error)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/hanafail | failed health checks details
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/req_rate | HTTP requests per second over last elapsed second
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/req_rate_max | max number of HTTP requests per second observed
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/req_tot | total number of HTTP requests received
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/cli_abrt | number of data transfers aborted by the client
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/srv_abrt | number of data transfers aborted by the server (inc. in eresp)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_in | number of HTTP response bytes fed to the compressor
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_out | number of HTTP response bytes emitted by the compressor
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_byp | number of bytes that bypassed the HTTP compressor (CPU/BW limit)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_rsp | number of HTTP responses that were compressed
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/lastsess | number of seconds since last session assigned to server/backend
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/last_chk | last health check contents or textual error
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/last_agt | last agent check contents or textual error
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/qtime | the average queue time in ms over the 1024 last requests
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/ctime | the average connect time in ms over the 1024 last requests
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/rtime | the average response time in ms over the 1024 last requests (0 for TCP)
/intel/linux/haproxy/stat/\<service_name\>/\<proxy_name\>/ttime | the average total session time in ms over the 1024 last requests

### Examples
Example running haproxy, passthru processor, and writing data to a file.

This is done from the snap directory.

In one terminal window, open the snap daemon (in this case with logging set to 1 and trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0
```

In another terminal window:
Load haproxy plugin
```
$ $SNAP_PATH/bin/snapctl plugin load snap-plugin-collector-haproxy
```
See available metrics for your system
```
$ $SNAP_PATH/bin/snapctl metric list
```

Create a task manifest file (e.g. `haproxy-file.json`):    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/linux/haproxy/info/CurrConns": {},
                "/intel/linux/haproxy/info/Run_queue": {}, 
                "/intel/linux/haproxy/stat/<service_name>/<proxy_name>/qlimit": {},
                "/intel/linux/haproxy/stat/<service_name>/<proxy_name>/chkdown": {} 
            },
            "config": {
                "/intel/mock": {
                    "password": "secret",
                    "user": "root"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {                         
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published_haproxy"
                            }
                        }
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```

Load passthru plugin for processing:
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-processor-passthru
Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:44:03 PST
```

Load file plugin for publishing:
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-publisher-file
Plugin loaded
Name: file
Version: 3
Type: publisher
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:41:39 PST
```

Create task:
```
$ $SNAP_PATH/bin/snapctl task create -t examples/tasks/haproxy-file.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

Stop task:
```
$ $SNAP_PATH/bin/snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@MarcinKrolik](https://github.com/marcin-krolik/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.