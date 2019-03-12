package logCustom

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logPath string, logLevel string) *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    1024,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       //days
		Compress:   true,    // 是否压缩 disabled by default
	}
	w := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewDevelopmentEncoderConfig() //returns an opinionated EncoderConfig for development environments.
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level)
	logger := zap.New(core)
	logger.Info("Init default logger")
	return logger
}
