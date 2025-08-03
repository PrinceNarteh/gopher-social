package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func NewLogger() {
	logger := zap.Must(zap.NewProduction())
	defer func(l zapcore.Level) bool {
		if err := logger.Sync(); err != nil {
			Logger.Error("Failed to sync logger", zap.Error(err))
			return false
		}
		return true
	}(logger.Level())
	Logger = logger.Sugar()
}
