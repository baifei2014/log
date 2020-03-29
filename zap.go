package log

import (
	"context"
	"fmt"
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

// NewStdout create a stdout log handler
func NewZap(outputDir, errorOutputDir string) *ZapLogger {

	cfg := zap.Config{
        Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
        Development: true,
        Encoding:    "json",
        EncoderConfig: zapcore.EncoderConfig{
	        MessageKey: "msg",
	        LevelKey:   "level",
	        TimeKey:    "time",
	        //CallerKey:      "file",
	        CallerKey:     "caller",
	        StacktraceKey: "trace",
	        LineEnding:    zapcore.DefaultLineEnding,
	        EncodeLevel:   zapcore.LowercaseLevelEncoder,
	        //EncodeLevel:    zapcore.CapitalLevelEncoder,
	        EncodeCaller: zapcore.ShortCallerEncoder,
	        EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	            enc.AppendString(t.Format("2006-01-02 15:04:05"))
	        },
	        //EncodeDuration: zapcore.SecondsDurationEncoder,
	        EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	            enc.AppendInt64(int64(d) / 1000000)
	        },
	    },
        OutputPaths:      []string{outputDir},
        ErrorOutputPaths: []string{errorOutputDir},
    }

	logger, err := cfg.Build();
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return &ZapLogger{
		logger:    logger,
	}
}

// Log stdout loging, only for developing env.
func (h *ZapLogger) Log(ctx context.Context, lv Level, args ...D) {
	var fields []zap.Field
	for _, v := range args[1:] {
		item := zap.Any(v.Key, v.Value)
		fields = append(fields, item)
	}
	h.logger.Info(args[0].Value.(string), fields...)
}

// Close stdout loging
func (h *ZapLogger) Close() error {
	return nil
}

// SetFormat set stdout log output format
// %T time format at "15:04:05.999"
// %t time format at "15:04:05"
// %D data format at "2006/01/02"
// %d data format at "01/02"
// %L log level e.g. INFO WARN ERROR
// %f function name and line number e.g. model.Get:121
// %i instance id
// %e deploy env e.g. dev uat fat prod
// %z zone
// %S full file name and line number: /a/b/c/d.go:23
// %s final file name element and line number: d.go:23
// %M log message and additional fields: key=value this is log message
func (h *ZapLogger) SetFormat(format string) {
}
