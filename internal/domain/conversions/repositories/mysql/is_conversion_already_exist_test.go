package mysql

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
)

func TestConversion_CountExistingConversion(t *testing.T) {
	db, mock := NewMock()

	defer func() {
		_ = db.Close()
	}()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `
		SELECT count(id) 'count' 
		FROM conversion_rate 
		WHERE
			(currency_id_from = ? AND currency_id_to = ?)
			OR
			(currency_id_from = ? AND currency_id_to = ?)
	`

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1, 2, 2, 1).WillReturnRows(rows)

	repo := conversion{db: sqlxDB}
	defer func() {
		_ = repo.db.Close()
	}()

	count, err := repo.CountExistingConversion(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
	})

	assert.Equal(t, count, int64(1))
	assert.NoError(t, err)
}
