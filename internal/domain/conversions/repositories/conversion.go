package repositories

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

type Conversion interface {
	IsAlreadyExist(ctx context.Context, params dto.CreateConversionRequest) (int64, error)
	Create(ctx context.Context, params dto.CreateConversionRequest) error
}
