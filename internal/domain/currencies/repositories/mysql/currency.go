package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/trwndh/game-currency/internal/domain/currencies/repositories"
)

type currency struct {
	db *sqlx.DB
}

func NewCurrency(db *sqlx.DB) repositories.Currency {
	return currency{db: db}
}
