package mysql

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func (c currency) CountByName(ctx context.Context, name string) (int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "[Currency][Repo][CountByName]")
	defer span.Finish()

	var count int32
	query := `SELECT count(id) 'count' FROM currency WHERE name = ?`
	err := c.db.QueryRowContext(ctx, query, name).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
