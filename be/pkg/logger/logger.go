package logger

import (
	"os"

	"flash-card.juslt.click/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	encoder := GetEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileFormat,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	logLevel := GetLogLevel(config)
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		logLevel)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

func GetEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func GetLogLevel(config setting.LoggerSetting) zapcore.Level {
	logLevel := config.LogLevel

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	return level
}
