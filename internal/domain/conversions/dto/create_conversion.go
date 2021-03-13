package dto

type CreateConversionRequest struct {
	CurrencyIDFrom int64 `json:"currency_id_from"`
	CurrencyIDTo   int64 `json:"currency_id_to"`
	Rate           int64 `json:"rate"`
}

func (c CreateConversionRequest) IsCurrencyIDFromEmpty() bool {
	return c.CurrencyIDFrom == 0
}

func (c CreateConversionRequest) IsCurrencyIDToEmpty() bool {
	return c.CurrencyIDTo == 0
}

type CreateConversionResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
