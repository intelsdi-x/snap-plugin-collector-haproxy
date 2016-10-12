# snap collector plugin - haproxy
This plugin collects metrics from running HAProxy by reading statistics and informations from Unix socket.

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating systems](#operating-systems)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [snap's Global Config](#snaps-global-config)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [golang 1.5+](https://golang.org/dl/) - needed only for building
* [HAProxy](http://haproxy.org) installed and running with 'stats socket' configuration

### Operating systems
All OSs currently supported by snap:
* Linux/amd64

### Installation
#### Download haproxy plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page. Download the plugins package from the latest release, unzip and store in a path you want `snapd` to access.

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
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started).
* Create HAProxy global configuration and launch HAProxy, see [HAProxy documentation](http://www.haproxy.org/download/1.6/doc/).
* Create Global Config, see description in [snap's Global Config] (https://github.com/intelsdi-x/snap-plugin-collector-haproxy/blob/master/README.md#snaps-global-config).
* Load the plugin and create a task, see example in [Examples](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/blob/master/README.md#examples).

## Documentation

### Collected Metrics
List of collected metrics is described in [METRICS.md](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/blob/master/METRICS.md).

### snap's Global Config
Global configuration files are described in [snap's documentation](https://github.com/intelsdi-x/snap/blob/master/docs/SNAPD_CONFIGURATION.md). You have to add section "haproxy" in "collector" section and then specify following options:
- `"socket"` -  indicates stats socket which is set in HAProxy global configuration, for more information read [HAProxy
Configuration Manual](http://www.haproxy.org/download/1.6/doc/configuration.txt).

See example Global Config in [example/cfg/] (https://github.com/intelsdi-x/snap-plugin-collector-haproxy/blob/master/example/cfg/).

### Examples
Example running haproxy and writing data to a file.

Make sure that your `$SNAP_PATH` is set, if not:
```
$ export SNAP_PATH=<snapDirectoryPath>/build
```
Other paths to files should be set according to your configuration, using a file you should indicate where it is located.

Download and install HAProxy, the latest version of HAProxy in available [here](http://www.haproxy.org/).

Create HAProxy configuration and save it in config.cfg, example HAProxy configuration:
```
global
    stats socket /var/run/haproxy.sock mode 600 level admin
    stats timeout 2m
frontend LB
    bind *:81
    default_backend LB
backend LB
    server Server01 127.0.0.1:8080 check
    server Server02 127.0.0.1:8081 check
```
As a root launch HAProxy with configuration:
```
$ haproxy -f config.cfg
```
Create Global Config, see example in [examples/cfg/] (examples/cfg/).

In one terminal window, open the snap daemon (in this case with logging set to 1,  trust disabled and global configuration saved in cfg.json):
```
$ snapd -l 1 -t 0 --config examples/cfg/snap-global-cfg.json
```
In another terminal window in snap-plugin-collector-haproxy directory:
Load haproxy plugin:
```
$ snapctl plugin load build/rootfs/snap-plugin-collector-haproxy
```
See available metrics for your system
```
$ snapctl metric list
```
Create a task manifest file (exemplary file in [examples/tasks/] (examples/tasks/)):
```
{
  "version": 1,
  "schedule": {
    "type": "simple",
    "interval": "1s"
  },
  "workflow": {
    "collect": {
      "metrics": {
        "/intel/haproxy/*": {}
      },
      "config": null,
      "publish": [
        {
          "plugin_name": "file",
          "config": {
            "file": "/tmp/published_haproxy.log"
          }
        }
      ]
    }
  }
}
```

Get mock file plugin for publishing, appropriate for Linux or Darwin:
```
wget  http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest_build/linux/x86_64/snap-plugin-publisher-file
```
or 
```
wget  http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest_build/darwin/x86_64/snap-plugin-publisher-file
```

Load mock file plugin for publishing:
```
$ snapctl plugin load snap-plugin-publisher-file
```
Create a task:
```
$ snapctl task create -t examples/tasks/haproxy-file.json

Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```
Stop previously created task:
```
$ snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850

Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. The full project is at http://github.com/intelsdi-x/snap.
To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [snap Gitter channel](https://gitter.im/intelsdi-x/snap).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [Marcin Krolik](https://github.com/marcin-krolik)
