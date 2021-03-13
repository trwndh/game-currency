package loggers

import (
	"runtime"
	"strings"

	stackdriver "github.com/tommy351/zap-stackdriver"
	"gitlab.warungpintar.co/back-end/libwp/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggers = func() *zap.Logger {
	_, file, line, _ := runtime.Caller(1)
	slash := strings.LastIndex(file, "/")
	file = file[slash+1:]

	var logger *zap.Logger
	if env.IsDevelopment() {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, _ = config.Build()
	} else {
		config := &zap.Config{
			Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Encoding:         "json",
			EncoderConfig:    stackdriver.EncoderConfig,
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
		logger, _ = config.Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return &stackdriver.Core{
				Core: core,
			}
		}), zap.Fields(
			stackdriver.LogServiceContext(&stackdriver.ServiceContext{
				Service: "garfield",
				Version: "1.0.0",
			}),
		),
			zap.Fields(
				stackdriver.LogReportLocation(&stackdriver.ReportLocation{
					FilePath:     file,
					LineNumber:   line,
					FunctionName: "",
				}),
			),
		)
	}
	return logger
}()
