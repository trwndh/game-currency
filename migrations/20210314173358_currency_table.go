package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCurrencyTable, downCurrencyTable)
}

const createCurrency = `CREATE TABLE currency (
id int(5) NOT NULL AUTO_INCREMENT,
name varchar(25) NOT NULL,
created_at timestamp NULL DEFAULT current_timestamp(),
updated_at timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
UNIQUE KEY name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
`

func upCurrencyTable(tx *sql.Tx) error {
	// This code is executed when the migrations is applied.
	_, err := tx.Exec(createCurrency)
	if err != nil {
		return err
	}
	return nil
}

func downCurrencyTable(tx *sql.Tx) error {
	// This code is executed when the migrations is rolled back.
	_, err := tx.Exec(`DROP TABLE currency`)
	if err != nil {
		return err
	}
	return nil

}
