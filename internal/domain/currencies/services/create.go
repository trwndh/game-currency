package services

import (
	"context"

	"github.com/trwndh/game-currency/internal/domain/currencies/entity"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"go.uber.org/zap"

	"github.com/trwndh/game-currency/internal/domain/currencies/errors"

	"github.com/opentracing/opentracing-go"
	"github.com/trwndh/game-currency/internal/domain/currencies/dto"
)

func (s service) Create(ctx context.Context, param dto.CreateCurrencyRequest) (dto.CreateCurrencyResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Currency][Service][Create]")
	defer span.Finish()

	if param.IsNameEmpty() {
		return dto.CreateCurrencyResponse{}, errors.GetErrorInvalidPayload()
	}

	count, err := s.CurrencyRepo.CountByName(ctx, param.Name)
	if err != nil {
		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.CreateCurrencyResponse{
			Error: errors.GetErrorDatabase().Error(),
		}, errors.GetErrorDatabase()
	}
	if count > 0 {
		return dto.CreateCurrencyResponse{}, errors.GetErrorCurrencyAlreadyExist()
	}

	err = s.CurrencyRepo.Create(ctx, entity.CurrencyDAO{
		Name: param.Name,
	})
	if err != nil {
		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.CreateCurrencyResponse{}, errors.GetErrorDatabase()
	}

	return dto.CreateCurrencyResponse{Status: "success"}, nil
}
