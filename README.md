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

Namespace | Data Type | Description
----------|-----------|-----------------------
/intel/haproxy/info/CompressBpsIn | int64 |
/intel/haproxy/info/CompressBpsOut | int64 |
/intel/haproxy/info/CompressBpsRateLim | int64 |
/intel/haproxy/info/ConnRate | int64 |
/intel/haproxy/info/ConnRateLimit | int64 |
/intel/haproxy/info/CumConns | int64 |
/intel/haproxy/info/CumReq | int64 |
/intel/haproxy/info/CumSslConns | int64 |
/intel/haproxy/info/CurrConns | int64 |
/intel/haproxy/info/CurrSslConns | int64 |
/intel/haproxy/info/Hard_maxconn | int64 |
/intel/haproxy/info/Idle_pct | int64 |
/intel/haproxy/info/MaxConnRate | int64 |
/intel/haproxy/info/MaxSessRate | int64 |
/intel/haproxy/info/MaxSslConns | int64 |
/intel/haproxy/info/MaxSslRate | int64 |
/intel/haproxy/info/MaxZlibMemUsage | int64 |
/intel/haproxy/info/Maxconn |  int64
/intel/haproxy/info/Maxpipes | int64 |
/intel/haproxy/info/Maxsock | int64 |
/intel/haproxy/info/Memmax_MB | int64 |
/intel/haproxy/info/Name | string |
/intel/haproxy/info/Nbproc | int64 |
/intel/haproxy/info/Pid | int64 |
/intel/haproxy/info/PipesFree | int64 |
/intel/haproxy/info/PipesUsed | int64 |
/intel/haproxy/info/Process_num | int64 |
/intel/haproxy/info/Release_date | string |
/intel/haproxy/info/Run_queue | int64 |
/intel/haproxy/info/SessRate | int64 |
/intel/haproxy/info/SessRateLimit | int64 |
/intel/haproxy/info/SslBackendKeyRate | int64 |
/intel/haproxy/info/SslBackendMaxKeyRate | int64 |
/intel/haproxy/info/SslCacheLookups | int64 |
/intel/haproxy/info/SslCacheMisses | int64 |
/intel/haproxy/info/SslFrontendKeyRate | int64 |
/intel/haproxy/info/SslFrontendMaxKeyRate | int64 |
/intel/haproxy/info/SslFrontendSessionReuse_pct | int64 |
/intel/haproxy/info/SslRate | int64 |
/intel/haproxy/info/SslRateLimit | int64 |
/intel/haproxy/info/Tasks | int64 |
/intel/haproxy/info/Ulimit-n | int64 |
/intel/haproxy/info/Uptime | string |
/intel/haproxy/info/Uptime_sec | int64 |
/intel/haproxy/info/Version | string |
/intel/haproxy/info/ZlibMemUsage | int64 |
/intel/haproxy/info/description | string |
/intel/haproxy/info/node | string |
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qcur | int64 | current queued requests. For the backend this reports the number queued without a server assigned.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qmax | int64 | max value of qcur
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/scur | int64 |current sessions
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/smax | int64 | max sessions
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/slim |  int64 | configured session limit
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/stot | int64 |cumulative number of connections
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/bin | int64| bytes in
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/bout | int64 | bytes out
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/dreq | int64 | requests denied because of security concerns.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/dresp | int64 | responses denied because of security concerns.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/ereq | int64 | request errors
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/econ | int64 | number of requests that encountered an error trying to  connect to a backend server
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/eresp | int64| response errors
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/wretr | int64 | number of times a connection to a server was retried.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/wredis | int64 | number of times a request was redispatched to another  server
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/status | string | status (UP/DOWN/NOLB/MAINT/MAINT(via)...)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/weight | int64 | total weight (backend), server weight (server)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/act | int64 | number of active servers (backend), server is active (server)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/bck | int64 | number of backup servers (backend), server is backup (server)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/chkfail | int64 |number of failed checks
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/chkdown | int64 | number of UP->DOWN transitions
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/lastchg | int64 |number of seconds since the last UP<->DOWN transition
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/downtime |  int64 | total downtime (in seconds)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qlimit | int64 | configured maxqueue for the server, or nothing in the  value is 0 (default, meaning no limit)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/pid | int64 | process id (0 for first instance, 1 for second, ...)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/iid | in64 | unique proxy id
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/sid | int64 | server id (unique inside a proxy)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/throttle | int64 | current throttle percentage for the server, when slowstart is active, or no value if not in slowstart.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/lbtot | int64 | total number of times a server was selected, either for new sessions, or when re-dispatching
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/tracked | int64 | id of proxy/server if tracking is enabled.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/type | int64 | (0=frontend, 1=backend, 2=server, 3=socket/listener)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rate | int64 | number of sessions per second over last elapsed second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rate_lim | int64 | configured limit on new sessions per second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rate_max | int64 | max number of new sessions per second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/check_status | string | status of last health check
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/check_code | int64 | layer5-7 code, if available
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/check_duration | int64 | time in ms took to finish last health check
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_1xx | int64 | http responses with 1xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_2xx | int64 | http responses with 2xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_3xx | int64 | http responses with 3xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_4xx | int64 | http responses with 4xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_5xx | int64 | http responses with 5xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_other | int64 | http responses with other codes (protocol error)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hanafail | string |failed health checks details
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/req_rate | int64 | HTTP requests per second over last elapsed second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/req_rate_max | int64 | max number of HTTP requests per second observed
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/req_tot | int64 | total number of HTTP requests received
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/cli_abrt | int64 | number of data transfers aborted by the client
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/srv_abrt | int64| number of data transfers aborted by the server (inc. in eresp)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_in | int64 | number of HTTP response bytes fed to the compressor
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_out | int64 | number of HTTP response bytes emitted by the compressor
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_byp | int64 | number of bytes that bypassed the HTTP compressor (CPU/BW limit)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_rsp | int64 | number of HTTP responses that were compressed
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/lastsess | int64 | number of seconds since last session assigned to server/backend
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/last_chk | string | last health check contents or textual error
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/last_agt | string | last agent check contents or textual error
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qtime | int64 | the average queue time in ms over the 1024 last requests
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/ctime | int64 | the average connect time in ms over the 1024 last requests
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rtime | int64 | the average response time in ms over the 1024 last requests (0 for TCP)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/ttime | int64 | the average total session time in ms over the 1024 last requests

**Important Note: Above metric list is a full list. Actual metrics are dependant on HAProxy service configuration.** 

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