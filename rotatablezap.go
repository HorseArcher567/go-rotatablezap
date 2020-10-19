package rotatablezap

import (
	"io"
	"time"

	"github.com/artisan-yp/go-rotatefiles"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(filename string, options ...rotatefiles.Option) *zap.Logger {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "livel",
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		TimeKey:       "time",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		CallerKey:     "line",
		FunctionKey:   "func",
		StacktraceKey: "stack",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		ConsoleSeparator: "\t",
	})

	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel
	})

	infoWriter := getWriter(filename + "_info")
	errorWriter := getWriter(filename + "_error")

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	logger := zap.New(core,
		zap.AddCaller(), zap.AddCallerSkip(0),
	)

	return logger
}

func getWriter(filename string, options ...rotatefiles.Option) io.Writer {
	file, err := rotatefiles.New(filename, options...)
	if err != nil {
		panic(err)
	}

	return file
}
