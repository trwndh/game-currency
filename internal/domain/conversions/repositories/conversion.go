package repositories

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/conversions/entity"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

type Conversion interface {
	CountExistingConversion(ctx context.Context, params dto.CreateConversionRequest) (int64, error)
	Create(ctx context.Context, params dto.CreateConversionRequest) error
	FindRate(ctx context.Context, params dto.CreateConversionRequest) (entity.ConversionRate, error)
}
