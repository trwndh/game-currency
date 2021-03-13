package services

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/currencies/errors"
	"go.uber.org/zap"

	"github.com/opentracing/opentracing-go"
	"github.com/trwndh/game-currency/internal/instrumentation/loggers"

	"github.com/trwndh/game-currency/internal/domain/currencies/dto"
)

func (s service) Find(ctx context.Context) (dto.GetCurrenciesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Currency][Service][Find]")
	defer span.Finish()

	listCurrencies, err := s.CurrencyRepo.Find(ctx)
	if err != nil {
		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.GetCurrenciesResponse{
			Error: errors.GetErrorDatabase().Error(),
		}, err
	}

	var response dto.GetCurrenciesResponse

	for _, currency := range listCurrencies {
		id := currency.ID
		name := currency.Name
		response.Currencies = append(response.Currencies, dto.Currency{
			ID:   id,
			Name: name,
		})
	}

	return response, nil
}
