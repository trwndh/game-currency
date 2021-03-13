package handler

import (
	"github.com/trwndh/game-currency/config"
	"github.com/trwndh/game-currency/internal/domain/conversions"
	"github.com/trwndh/game-currency/internal/domain/currencies"
	"github.com/trwndh/game-currency/internal/handler/gen"
)

type HttpServer struct {
	cfg               *config.MainConfig
	currencyService   currencies.Service
	conversionService conversions.Service
}

func NewHttpServer(
	cfg *config.MainConfig,
	currencyService currencies.Service,
	conversionService conversions.Service,
) gen.ServerInterface {
	return &HttpServer{
		cfg:               cfg,
		currencyService:   currencyService,
		conversionService: conversionService,
	}
}
