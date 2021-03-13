package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/currencies/entity"
)

func (c currency) Create(ctx context.Context, params entity.CurrencyDAO) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Currency][Repo][Create]")
	defer span.Finish()

	query := `
		INSERT INTO currency(name) VALUES (?)
	`
	_, err := c.db.Slave.ExecContext(ctx, query, params.Name)
	if err != nil {
		return err
	}
	return nil
}
