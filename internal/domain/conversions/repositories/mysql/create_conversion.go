package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

func (c conversion) Create(ctx context.Context, params dto.CreateConversionRequest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Repo][Create]")
	defer span.Finish()

	query := `
		INSERT INTO conversion_rate(currency_id_from,currency_id_to,rate)
		VALUES (?,?,?)
	`

	_, err := c.db.Master.ExecContext(ctx, query, params.CurrencyIDFrom, params.CurrencyIDTo, params.Rate)
	if err != nil {
		return err
	}

	return nil
}
