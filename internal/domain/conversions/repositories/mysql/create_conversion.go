package mysql

import (
	"context"

	"github.com/go-sql-driver/mysql"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
	"github.com/trwndh/game-currency/internal/domain/conversions/errors"
)

func (c conversion) Create(ctx context.Context, params dto.CreateConversionRequest) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Repo][Create]")
	defer span.Finish()

	query := `
		INSERT INTO conversion_rate(currency_id_from,currency_id_to,rate)
		VALUES (?,?,?)
	`

	_, err := c.db.ExecContext(ctx, query, params.CurrencyIDFrom, params.CurrencyIDTo, params.Rate)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1452 {
				return errors.Get1452Error()
			}
		}
		return err
	}

	return nil
}
