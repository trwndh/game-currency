package services

import (
	"context"
	"database/sql"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"go.uber.org/zap"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
	"github.com/trwndh/game-currency/internal/domain/conversions/errors"
)

func (s service) ConvertCurrency(ctx context.Context, param dto.ConvertCurrencyRequest) (dto.ConvertCurrencyResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Service][ConvertCurrency]")
	defer span.Finish()

	if param.IsCurrencyIDFromEmpty() || param.IsCurrencyIDToEmpty() {
		return dto.ConvertCurrencyResponse{Error: errors.GetErrorInvalidPayload().Error()}, errors.GetErrorInvalidPayload()
	}

	conversionRate, err := s.ConversionRepo.FindRate(ctx, dto.CreateConversionRequest{
		CurrencyIDFrom: param.CurrencyIDFrom,
		CurrencyIDTo:   param.CurrencyIDTo,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			loggers.For(ctx).Error(errors.GetErrorConversionNotFound().Error(), zap.Error(err))
			return dto.ConvertCurrencyResponse{
				Error: errors.GetErrorConversionNotFound().Error(),
			}, errors.GetErrorConversionNotFound()
		}

		loggers.For(ctx).Error(errors.GetErrorDatabase().Error(), zap.Error(err))
		return dto.ConvertCurrencyResponse{
			Error: errors.GetErrorDatabase().Error(),
		}, errors.GetErrorDatabase()
	}

	// for reversed result
	//  example, in db stored mapping from 1 to 2
	// but found mapping 2 to 1
	// result = amount / rate
	var result int64
	if conversionRate.CurrencyIDTo == param.CurrencyIDFrom && conversionRate.CurrencyIDFrom == param.CurrencyIDTo {
		result = param.Amount / conversionRate.Rate
	} else {
		result = param.Amount * conversionRate.Rate
	}

	return dto.ConvertCurrencyResponse{
		Result: result,
	}, nil

}
