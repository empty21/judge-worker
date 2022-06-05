package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func createLogger(logName string, logLevel zapcore.Level, filePath string) *zap.SugaredLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	ll := lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    100, // 100mb
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}
	err := zap.RegisterSink(logName, func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &ll,
		}, nil
	})
	loggerConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(logLevel),
		Development:   logLevel != zap.DebugLevel,
		Encoding:      "console",
		EncoderConfig: encoderConfig,
		OutputPaths:   []string{"stdout", fmt.Sprintf("%s:%s", logName, filePath)},
	}

	_logger, err := loggerConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("Build zap logger from config error: %v", err))
	}
	return _logger.Sugar()
}
