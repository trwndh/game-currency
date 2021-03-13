package handler

import (
	"github.com/trwndh/game-currency/config"
	"github.com/trwndh/game-currency/internal/domain/currencies"
	"github.com/trwndh/game-currency/internal/handler/gen"
)

type HttpServer struct {
	cfg             *config.MainConfig
	currencyService currencies.Service
}

func NewHttpServer(cfg *config.MainConfig, currencyService currencies.Service) gen.ServerInterface {
	return &HttpServer{
		cfg:             cfg,
		currencyService: currencyService,
	}
}
