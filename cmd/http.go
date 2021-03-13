package cmd

import (
	"net/http"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"github.com/trwndh/game-currency/config"
	conversionRepo "github.com/trwndh/game-currency/internal/domain/conversions/repositories/mysql"
	conversion "github.com/trwndh/game-currency/internal/domain/conversions/services"
	currencyRepo "github.com/trwndh/game-currency/internal/domain/currencies/repositories/mysql"
	currency "github.com/trwndh/game-currency/internal/domain/currencies/services"
	"github.com/trwndh/game-currency/internal/handler"
	"github.com/trwndh/game-currency/internal/handler/gen"
	httpServer "github.com/trwndh/game-currency/internal/server/http"
)

var HttpCmd = &cobra.Command{
	Use:   "http-start",
	Short: "Start http server for REST API",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadMainConfig()

		dbMysql, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.DSN)
		if err != nil {
			loggers.Bg().Error("error connecting to database", zap.Error(err))
		}
		defer func() {
			_ = dbMysql.Close()
		}()
		currencyService := currency.NewService(
			currencyRepo.NewCurrency(dbMysql),
		)

		conversionService := conversion.NewService(
			conversionRepo.NewConversion(dbMysql),
		)

		httpServer.RunHTTPServer(cfg, func(router chi.Router) http.Handler {
			return gen.HandlerFromMux(
				handler.NewHttpServer(cfg, currencyService, conversionService), router,
			)
		})
	},
}
