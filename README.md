# Snap collector plugin - haproxy
This plugin collects metrics from running HAProxy by reading statistics and informations from Unix socket.

It's used in the [Snap framework](https://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating systems](#operating-systems)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Snap's Global Config](#snaps-global-config)
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
All OSs currently supported by Snap:
* Linux/amd64

### Installation
#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/releases) page. Download the plugin from the latest release and load it into `snapd` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-haproxy
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-haproxy.git
```

Build the Snap haproxy plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap#getting-started).
* Create HAProxy global configuration and launch HAProxy, see [HAProxy documentation](http://www.haproxy.org/download/1.6/doc/).
* Create Global Config, see description in [Snap's Global Config] (https://github.com/intelsdi-x/snap-plugin-collector-haproxy#snaps-global-config).
* Load the plugin and create a task, see example in [Examples](https://github.com/intelsdi-x/snap-plugin-collector-haproxy#examples).

## Documentation

### Collected Metrics
List of collected metrics is described in [METRICS.md](METRICS.md).

### Snap's Global Config
Global configuration files are described in [Snap's documentation](https://github.com/intelsdi-x/snap/blob/master/docs/SNAPD_CONFIGURATION.md). You have to add section "haproxy" in "collector" section and then specify following options:
- `"socket"` -  indicates stats socket which is set in HAProxy global configuration, for more information read [HAProxy
Configuration Manual](http://www.haproxy.org/download/1.6/doc/configuration.txt).

See example Global Config in [example cfg] (examples/cfg/snap-global-cfg.json).

### Examples
Example running haproxy and writing data to a file.

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

Download an [example Snap global config](examples/cfg/snap-global-cfg.json) file.
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-haproxy/master/examples/cfg/snap-global-cfg.json
```

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap) with provided configuration file:
* command line: `snapd -l 1 -t 0 --config cfg.json&`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-haproxy/latest/linux/x86_64/snap-plugin-collector-haproxy
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snapctl plugin load snap-plugin-collector-haproxy
$ snapctl plugin load snap-plugin-publisher-file
```

See available metrics for your system
```
$ snapctl metric list
```

Download an [example task file](examples/tasks/haproxy-file.json) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-haproxy/master/examples/tasks/haproxy-file.json
$ snapctl task create -t haproxy-file.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See realtime output from `snapctl task watch <task_id>` (CTRL+C to exit)
```
$ snapctl task watch 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

This data is published to a file `/tmp/published_haproxy.log` per task specification

Stop task:
```
$ snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850

Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-haproxy/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [Marcin Krolik](https://github.com/marcin-krolik)
