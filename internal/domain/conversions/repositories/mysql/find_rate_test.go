package mysql

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"go.uber.org/zap"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		loggers.Bg().Error("an error '%s' was not expected when opening a stub database connection", zap.Error(err))
	}

	return db, mock
}

func TestConversion_FindRate(t *testing.T) {
	db, mock := NewMock()

	defer func() {
		_ = db.Close()
	}()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `
	SELECT 
		currency_id_from, currency_id_to, rate
	FROM conversion_rate
	WHERE
		(currency_id_from = ? AND currency_id_to = ?)
		OR 
		(currency_id_from = ? AND currency_id_to = ?)
	`

	rows := sqlmock.NewRows([]string{"currency_id_from", "currency_id_to", "rate"}).AddRow(1, 2, 20)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1, 2, 2, 1).WillReturnRows(rows)

	repo := conversion{db: sqlxDB}
	defer func() {
		_ = repo.db.Close()
	}()

	rate, err := repo.FindRate(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
	})

	assert.Equal(t, rate.Rate, int64(20))
	assert.NoError(t, err)
}
