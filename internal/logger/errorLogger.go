package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ErrorLogger *zap.Logger

func InitLogger() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"stdout",
		"logs/errors.log", // Output to file
	}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	ErrorLogger, err = config.Build()
	if err != nil {
		panic("Failed to initialize errorLogger: " + err.Error())
	}
}
