package loggers

import (
	"go.uber.org/zap"
)

// Bg stands for background, to get logger with no hassle
func Bg() *zap.Logger {
	return loggers
}
