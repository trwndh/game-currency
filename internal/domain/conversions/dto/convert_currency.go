package dto

type ConvertCurrencyRequest struct {
	CurrencyIDFrom int64 `json:"currency_id_from"`
	CurrencyIDTo   int64 `json:"currency_id_to"`
	Amount         int64 `json:"amount"`
}

type ConvertCurrencyResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
