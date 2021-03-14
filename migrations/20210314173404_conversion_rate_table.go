package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upConversionRateTable, downConversionRateTable)
}

const createConversionRate = `CREATE TABLE conversion_rate (
id int(5) NOT NULL AUTO_INCREMENT,
currency_id_from int(5) NOT NULL,
currency_id_to int(5) NOT NULL,
rate int(5) NOT NULL,
created_at timestamp NOT NULL DEFAULT current_timestamp(),
updated_at timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
PRIMARY KEY (id),
KEY id_currency_from (currency_id_from),
KEY id_currency_to (currency_id_to),
CONSTRAINT conversion_rate_ibfk_1 FOREIGN KEY (currency_id_from) REFERENCES currency (id),
CONSTRAINT conversion_rate_ibfk_2 FOREIGN KEY (currency_id_to) REFERENCES currency (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

func upConversionRateTable(tx *sql.Tx) error {
	// This code is executed when the migrations is applied.
	_, err := tx.Exec(createConversionRate)
	if err != nil {
		return err
	}
	return nil
}

func downConversionRateTable(tx *sql.Tx) error {
	// This code is executed when the migrations is rolled back.
	_, err := tx.Exec(`DROP TABLE conversion_rate`)
	if err != nil {
		return err
	}
	return nil
}
