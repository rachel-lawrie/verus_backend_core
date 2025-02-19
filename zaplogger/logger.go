package zaplogger

import (
	"os"

	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once     sync.Once
	instance *zap.Logger
)

// GetLogger returns a singleton instance of the logger
func GetLogger() *zap.Logger {
	once.Do(func() {
		// Read the log level from the environment variable (default to "info")
		logLevel := os.Getenv("LOG_LEVEL")
		if logLevel == "" {
			logLevel = "info" // Default to info if not set
		}

		// Set the log level based on the environment variable
		var zapLevel zapcore.Level
		switch logLevel {
		case "debug":
			zapLevel = zapcore.DebugLevel
		case "info":
			zapLevel = zapcore.InfoLevel
		case "warn":
			zapLevel = zapcore.WarnLevel
		case "error":
			zapLevel = zapcore.ErrorLevel
		default:
			zapLevel = zapcore.InfoLevel
		}
		encoderCfg := zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			MessageKey:     "message",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder, // Loki-friendly timestamp
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)

		consoleCore := zapcore.NewCore(
			jsonEncoder,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // Direct to stdout without prefix
			zap.NewAtomicLevelAt(zapLevel),                          // Adjust log level as needed
		)

		instance = zap.New(consoleCore)

		// var err error
		// instance, err = zap.Config{
		// 	Encoding:         "json",
		// 	Level:            zapLevel,
		// 	OutputPaths:      []string{"stdout"}, // Logs to stdout for Promtail or Loki
		// 	ErrorOutputPaths: []string{"stderr"},
		// 	EncoderConfig: zapcore.EncoderConfig{
		// 		TimeKey:        "time",
		// 		LevelKey:       "level",
		// 		MessageKey:     "message",
		// 		CallerKey:      "caller",
		// 		StacktraceKey:  "stacktrace",
		// 		EncodeLevel:    zapcore.CapitalLevelEncoder,
		// 		EncodeTime:     zapcore.ISO8601TimeEncoder, // Loki-friendly timestamp
		// 		EncodeDuration: zapcore.StringDurationEncoder,
		// 		EncodeCaller:   zapcore.ShortCallerEncoder,
		// 	},
		// }.Build()

		// if err != nil {
		// 	panic("Failed to initialize logger")
		// }
	})
	return instance
}
