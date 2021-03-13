package repositories

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/currencies/entity"
)

type Currency interface {
	CountByName(ctx context.Context, name string) (int32, error)
	Create(ctx context.Context, params entity.CurrencyDAO) error
	Find(ctx context.Context) ([]entity.CurrencyDAO, error)
}
