package currencies

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/currencies/dto"
)

type Service interface {
	Find(ctx context.Context) (dto.GetCurrenciesResponse, error)
	Create(ctx context.Context, param dto.CreateCurrencyRequest) (dto.CreateCurrencyResponse, error)
}
