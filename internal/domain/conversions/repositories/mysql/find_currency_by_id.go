package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/trwndh/game-currency/internal/domain/conversions/entity"
)

func (c conversion) FindCurrencyByID(ctx context.Context, id int64) (entity.CurrencyDAO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Conversion][Repo][FindCurrencyByID]")
	defer span.Finish()

	query := `SELECT id, name FROM currency WHERE id = ?`

	var currency entity.CurrencyDAO
	err := c.db.QueryRowContext(ctx, query, id).Scan(&currency.ID, &currency.Name)
	if err != nil {
		return entity.CurrencyDAO{}, err
	}
	return currency, nil
}
