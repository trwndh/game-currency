package services

import (
	"github.com/trwndh/game-currency/internal/domain/currencies"
	"github.com/trwndh/game-currency/internal/domain/currencies/repositories"
)

type service struct {
	CurrencyRepo repositories.Currency
}

func NewService(CurrencyRepo repositories.Currency) currencies.Service {
	return service{CurrencyRepo: CurrencyRepo}
}
