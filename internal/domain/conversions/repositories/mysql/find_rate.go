package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
	"github.com/trwndh/game-currency/internal/domain/conversions/entity"
)

func (c conversion) FindRate(ctx context.Context, params dto.CreateConversionRequest) (entity.ConversionRate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Repo][FindRate]")
	defer span.Finish()

	query := `
	SELECT 
		currency_id_from, currency_id_to, rate
	FROM conversion_rate
	WHERE
		(currency_id_from = ? AND currency_id_to = ?)
		OR 
		(currency_id_from = ? AND currency_id_to = ?)
	`
	var conversionRate entity.ConversionRate
	err := c.db.QueryRowContext(ctx, query,
		params.CurrencyIDFrom, params.CurrencyIDTo,
		params.CurrencyIDTo, params.CurrencyIDFrom,
	).Scan(&conversionRate.CurrencyIDFrom, &conversionRate.CurrencyIDTo, &conversionRate.Rate)
	if err != nil {
		return entity.ConversionRate{}, err
	}
	return conversionRate, nil
}
