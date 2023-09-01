package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
)

func New(c *viper.Viper) *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.ConsoleSeparator = " | "
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writer := getWriteSyncer(c.GetString("log.driver"), c.GetString("log.file_name"))
	level := convLevel(c.GetString("log.level"))
	core := zapcore.NewCore(encoder, writer, level)
	return zap.New(core)
}

func convLevel(s string) zapcore.LevelEnabler {
	switch strings.ToLower(s) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

func getWriteSyncer(driver string, fileName string) zapcore.WriteSyncer {
	switch driver {
	case "console":
		return zapcore.AddSync(consoleWriter())
	case "file":
		return zapcore.AddSync(fileWriter(fileName))
	default:
		return zapcore.AddSync(consoleWriter())
	}
}

func consoleWriter() io.Writer {
	return os.Stderr
}

func fileWriter(fileName string) io.Writer {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 0,
		LocalTime:  true,
		Compress:   false,
	}
}
