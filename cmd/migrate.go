package cmd

import (
	"github.com/pressly/goose"
	"github.com/spf13/cobra"
	"github.com/trwndh/game-currency/config"
	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	_ "github.com/trwndh/game-currency/migrations"
)

var MigrateDatabase = &cobra.Command{
	Use:   "migrate",
	Short: "Create table needed for this service",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadMainConfig()

		db, err := goose.OpenDBWithDriver(cfg.Database.Driver, cfg.Database.DSN)
		if err != nil {
			loggers.Bg().Fatal("cannot conncect to database")
		}

		defer func() {
			if err := db.Close(); err != nil {
				loggers.Bg().Fatal("cannot conncect to database")
			}
		}()

		if err := goose.Run("up", db, "migrations"); err != nil {
			loggers.Bg().Fatal(err.Error())
		}
	},
}
