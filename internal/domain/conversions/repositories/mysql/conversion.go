package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/trwndh/game-currency/internal/domain/conversions/repositories"
)

type conversion struct {
	db *sqlx.DB
}

func NewConversion(db *sqlx.DB) repositories.Conversion {
	return conversion{db: db}
}
