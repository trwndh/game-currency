package mysql

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
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

func TestCurrency_CountByName(t *testing.T) {
	db, mock := NewMock()

	defer func() {
		_ = db.Close()
	}()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `SELECT count(id) 'count' FROM currency WHERE name = ?`

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("knut").WillReturnRows(rows)

	repo := currency{db: sqlxDB}
	defer func() {
		_ = repo.db.Close()
	}()

	count, err := repo.CountByName(context.Background(), "knut")

	assert.Equal(t, count, int32(0))
	assert.NoError(t, err)

}

func TestCurrency_CountByName_Error(t *testing.T) {
	db, mock := NewMock()

	defer func() {
		_ = db.Close()
	}()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `SELECT count(id) 'count' FROM currency WHERE name = ?`

	mock.ExpectQuery(query).WithArgs("knut").WillReturnError(sql.ErrNoRows)

	repo := currency{db: sqlxDB}
	defer func() {
		_ = repo.db.Close()
	}()

	count, err := repo.CountByName(context.Background(), "knut")

	assert.Equal(t, count, int32(0))
	assert.Error(t, err)

}
