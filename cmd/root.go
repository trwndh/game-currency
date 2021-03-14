package cmd

import (
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"github.com/trwndh/game-currency/pkg/tracing"
	"go.uber.org/zap"
)

var RootCmd = &cobra.Command{
	Use:   "game-currency",
	Short: "Game currency API",
}

func Execute() {
	tracer, closer := tracing.InitFromEnv("game-currency")
	defer func() {
		if err := closer.Close(); err != nil {
			loggers.Bg().Error("error on execute game-currency", zap.Error(err))
		}
	}()
	opentracing.SetGlobalTracer(tracer)

	RootCmd.AddCommand(HttpCmd)
	RootCmd.AddCommand(MigrateDatabase)

	if err := RootCmd.Execute(); err != nil {
		loggers.Bg().Error("error when execute root command", zap.Error(err))
	}
}
