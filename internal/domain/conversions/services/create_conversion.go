package services

import (
	"context"

	"go.uber.org/zap"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"

	"github.com/trwndh/game-currency/internal/domain/conversions/errors"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

func (s service) Create(ctx context.Context, param dto.CreateConversionRequest) (dto.CreateConversionResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Service][Create]")
	defer span.Finish()

	if param.IsCurrencyIDFromEmpty() || param.IsCurrencyIDToEmpty() {
		return dto.CreateConversionResponse{Error: errors.GetErrorInvalidPayload().Error()}, errors.GetErrorInvalidPayload()
	}

	count, err := s.ConversionRepo.IsAlreadyExist(ctx, param)
	if err != nil {
		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
	}

	if count > 0 {
		return dto.CreateConversionResponse{Error: errors.GetErrorConversionAlreadyExist().Error()}, errors.GetErrorConversionAlreadyExist()
	}

	err = s.ConversionRepo.Create(ctx, param)
	if err != nil {
		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.CreateConversionResponse{Error: errors.GetErrorDatabase().Error()}, errors.GetErrorDatabase()
	}

	return dto.CreateConversionResponse{Status: "success"}, nil
}
