package mysql

import (
	"github.com/trwndh/game-currency/internal/domain/conversions/repositories"
	"github.com/trwndh/game-currency/pkg/database"
)

type conversion struct {
	db *database.Store
}

func NewConversion(db *database.Store) repositories.Conversion {
	return conversion{db: db}
}
