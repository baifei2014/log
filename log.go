package log

import (
	"fmt"
	"context"
)

type Config struct {
	OutputDir string
	ErrorOutputDir string
}

// D represents a map of entry level data used for structured logging.
// type D map[string]interface{}
type D struct {
	Key   string
	Value interface{}
}

// KV return a log kv for logging field.
func KV(key string, value interface{}) D {
	return D{
		Key:   key,
		Value: value,
	}
}

var (
	h Handler
)

func init() {
	h = newHandlers([]string{}, NewStdout())
}

func Init(c *Config) {
	h = newHandlers([]string{}, NewZap(c.OutputDir, c.ErrorOutputDir))
}

func Info(format string, args ...interface{}) {
	h.Log(context.Background(), _infoLevel, KV(_log, fmt.Sprintf(format, args...)))
}

// Close close resource.
func Close() (err error) {
	err = h.Close()
	h = _defaultStdout
	return
}