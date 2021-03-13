package conversions

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

type Service interface {
	Create(ctx context.Context, param dto.CreateConversionRequest) (dto.CreateConversionResponse, error)
	ConvertCurrency(ctx context.Context, param dto.ConvertCurrencyRequest) (dto.ConvertCurrencyResponse, error)
}
