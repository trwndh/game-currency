package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/trwndh/game-currency/internal/domain/currencies/entity"
)

func (c currency) Find(ctx context.Context) ([]entity.CurrencyDAO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Currency][Repo][Find]")
	defer span.Finish()

	query := `SELECT id, name FROM currency ORDER BY id ASC`

	type response struct {
		ID   int64  `db:"id"`
		Name string `db:"name"`
	}

	var res []response

	var responseDAO []entity.CurrencyDAO
	err := c.db.SelectContext(ctx, &res, query)
	if err != nil {
		return responseDAO, err
	}
	for _, v := range res {
		responseDAO = append(responseDAO, entity.CurrencyDAO{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return responseDAO, nil
}
