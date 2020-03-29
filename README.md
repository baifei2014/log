# log

fast convenient and flexible

## Installation

`go get -u github.com/baifei2014/log`

Note that log only supports the two most recent minor versions of Go.

## Quick Start

```go
config := &log.Config{
	OutputDir: "",
	ErrorOutputDir: "",
}
log.Init(config)
defer log.Close()

log.Info("success...")
```