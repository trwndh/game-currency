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
		loggers.For(ctx).Error(errors.GetErrorInvalidPayload().Error(), zap.Error(errors.GetErrorInvalidPayload()))
		return dto.CreateConversionResponse{
			Error: errors.GetErrorInvalidPayload().Error(),
		}, errors.GetErrorInvalidPayload()
	}

	if param.IsRateEmpty() {
		loggers.For(ctx).Error(errors.GetErrorRateIsZero().Error(), zap.Error(errors.GetErrorRateIsZero()))
		return dto.CreateConversionResponse{
			Error: errors.GetErrorRateIsZero().Error(),
		}, errors.GetErrorRateIsZero()
	}

	if param.IsBothCurrencyIDIdentical() {
		loggers.For(ctx).Error(errors.GetErrorConvertingSameID().Error(), zap.Error(errors.GetErrorConvertingSameID()))
		return dto.CreateConversionResponse{
			Error: errors.GetErrorConvertingSameID().Error(),
		}, errors.GetErrorConvertingSameID()
	}

	count, err := s.ConversionRepo.CountExistingConversion(ctx, param)
	if err != nil {
		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.CreateConversionResponse{
			Error: errors.GetErrorDatabase().Error(),
		}, errors.GetErrorDatabase()
	}

	if count > 0 {
		return dto.CreateConversionResponse{
			Error: errors.GetErrorConversionAlreadyExist().Error(),
		}, errors.GetErrorConversionAlreadyExist()
	}

	err = s.ConversionRepo.Create(ctx, param)
	if err != nil {

		// currency ID not available
		if err.Error() == errors.Get1452Error().Error() {
			loggers.For(ctx).Error("elele"+errors.Get1452Error().Error(), zap.Error(err))
			return dto.CreateConversionResponse{
				Error: errors.GetErrorCurrenciesNotFound().Error(),
			}, errors.GetErrorCurrenciesNotFound()
		}

		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.CreateConversionResponse{
			Error: errors.GetErrorDatabase().Error(),
		}, errors.GetErrorDatabase()
	}

	return dto.CreateConversionResponse{Status: "success"}, nil
}
