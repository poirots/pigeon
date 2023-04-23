package log

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger(path string, level string) error {
	Logger, _ = zap.NewProduction()
	return nil
}
