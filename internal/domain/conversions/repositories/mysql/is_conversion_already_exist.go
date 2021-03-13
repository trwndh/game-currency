package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

func (c conversion) IsAlreadyExist(ctx context.Context, params dto.CreateConversionRequest) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Repo][IsAlreadyExist]")
	defer span.Finish()

	query := `
		SELECT count(id) 'count' 
		FROM conversion_rate 
		WHERE
			(currency_id_from = ? AND currency_id_to = ?)
			OR
			(currency_id_to = ? AND currency_id_from = ?)
	`
	var count int64
	err := c.db.Slave.QueryRowContext(ctx, query,
		params.CurrencyIDFrom, params.CurrencyIDTo,
		params.CurrencyIDTo, params.CurrencyIDFrom,
	).Scan(&count)

	if err != nil {
		return 0, err
	}
	return count, nil

	panic("implement me")
}
