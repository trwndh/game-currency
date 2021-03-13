package entity

import "time"

type ConversionRate struct {
	ID             int64     `json:"id" db:"id"`
	CurrencyIDFrom int64     `json:"currency_id_from" db:"currency_id_from"`
	CurrencyIDTo   int64     `json:"currency_id_to" db:"currency_id_to"`
	Rate           int64     `json:"rate" db:"rate"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
