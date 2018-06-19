AnomLog
===============

General text log anomaly detection engine.

Prerequisite
----------

- Go >= 1.10.2
- [dep](https://github.com/golang/dep) >= 0.4.1 (for development)

Also, environment variables `$GOPATH` is required to set. See [official document](https://github.com/golang/go/wiki/SettingGOPATH) for detail.


Setup
----------

As a CLI tool.

```
$ go get github.com/m-mizutani/anomlog
```

Usage
-----------

### Basic usage

```
$ anomlog -t 1000 your_log_file.log
```

`-t` option specifies training data size (basically line number). The example command trains with head of 1,000 lines in `your_log_file.log`. After training, the process starts anomaly detection with after the 1,001th line log.



