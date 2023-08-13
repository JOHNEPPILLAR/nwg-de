// Package utility - wrapper for zap logger
package utility

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger -
type Logger struct {
	zap *zap.Logger
}

// NewLogger - Creates a new logger instance
func NewLogger() (*Logger, error) {
	var cfg zap.Config
	environment := os.Getenv("ENVIRONMENT")
	if environment == "Development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.TimeKey = zapcore.OmitKey
	} else {
		rawJSON := []byte(`{
			"level": "info",
			"encoding": "json",
			"outputPaths": ["stdout"],
			"errorOutputPaths": ["stderr"],
			"encoderConfig": {
			  "messageKey": "message",
			  "levelKey": "level",
			  "levelEncoder": "lowercase"
			}
		  }`)
		if err := json.Unmarshal(rawJSON, &cfg); err != nil {
			return nil, err
		}
	}

	logger, err := cfg.Build(zap.AddCaller())
	defer logger.Sync()
	if err != nil {
		return nil, err
	}
	return &Logger{zap: logger}, err
}

// Debug - Log with debug
func (l Logger) Debug(msg string, fields ...zap.Field) {
	l.writer().Debug(msg, fields...)
}

// Info - Log with info
func (l Logger) Info(msg string, fields ...zap.Field) {
	l.writer().Info(msg, fields...)
}

// Warn - Log with warn
func (l Logger) Warn(msg string, fields ...zap.Field) {
	l.writer().Warn(msg, fields...)
}

// Error - Log with error
func (l Logger) Error(msg string, fields ...zap.Field) {
	l.writer().Error(msg, fields...)
}

// Fatal - Log with fatal
func (l Logger) Fatal(msg string, fields ...zap.Field) {
	l.writer().Fatal(msg, fields...)
}

var noOpLogger = zap.NewNop()

func (l Logger) writer() *zap.Logger {
	if l.zap == nil {
		return noOpLogger
	}
	return l.zap
}
