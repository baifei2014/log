package log

import (
	"context"
	"fmt"
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

// NewStdout create a stdout log handler
func NewZap(outputDir, errorOutputDir string) *ZapLogger {
	// sugar := zap.NewExample().Sugar()

	// cfg := zap.Config{
	// 	Encoding: "json",
	// 	OutputPaths: []string {"/Users/jianglonghao/go/src/music/zap"},
	// }

	cfg := zap.Config{
        Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
        Development: true,
        Encoding:    "json",
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "t",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "msg",
            StacktraceKey:  "trace",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.LowercaseLevelEncoder,
            EncodeTime:     formatEncodeTime,
            EncodeDuration: zapcore.SecondsDurationEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
        OutputPaths:      []string{outputDir},
        ErrorOutputPaths: []string{errorOutputDir},
        InitialFields: map[string]interface{}{
            "app": "test",
        },
    }

	logger, err := cfg.Build();
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	defer sugar.Sync()
	return &ZapLogger{
		logger:    sugar,
	}
}

// Log stdout loging, only for developing env.
func (h *ZapLogger) Log(ctx context.Context, lv Level, args ...D) {
	h.logger.Info(args)
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
