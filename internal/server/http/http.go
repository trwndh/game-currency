package http

import (
	"net/http"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"

	"github.com/trwndh/game-currency/internal/middleware"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/trwndh/game-currency/config"
)

func RunHTTPServer(cfg *config.MainConfig, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter, cfg)

	rootRouter := chi.NewRouter()
	rootRouter.Mount(cfg.Server.BaseURL, createHandler(apiRouter))

	loggers.Bg().Info("starting http server")

	_ = http.ListenAndServe(cfg.Server.Port, rootRouter)
}

func setMiddlewares(router *chi.Mux, cfg *config.MainConfig) {
	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(middleware.CheckSecretKey(cfg))
	router.Use(chiMiddleware.NoCache)
}
