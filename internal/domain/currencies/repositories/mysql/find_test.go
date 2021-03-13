package mysql

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCurrency_Find(t *testing.T) {
	db, mock := NewMock()

	defer func() {
		_ = db.Close()
	}()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	query := `SELECT id, name FROM currency ORDER BY id ASC`

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "knut")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	repo := currency{db: sqlxDB}
	defer func() {
		_ = repo.db.Close()
	}()

	listCurrencies, err := repo.Find(context.Background())

	assert.Equal(t, listCurrencies[0].Name, "knut")
	assert.NoError(t, err)
}
