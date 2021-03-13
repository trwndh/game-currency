package mysql

import (
	"github.com/trwndh/game-currency/internal/domain/currencies/repositories"
	"github.com/trwndh/game-currency/pkg/database"
)

type currency struct {
	db *database.Store
}

func NewCurrency(db *database.Store) repositories.Currency {
	return currency{db: db}
}
