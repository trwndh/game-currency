package dto

type ConvertCurrencyRequest struct {
	CurrencyIDFrom int64 `json:"currency_id_from"`
	CurrencyIDTo   int64 `json:"currency_id_to"`
	Amount         int64 `json:"amount"`
}

func (c ConvertCurrencyRequest) IsCurrencyIDFromEmpty() bool {
	return c.CurrencyIDFrom == 0
}
func (c ConvertCurrencyRequest) IsCurrencyIDToEmpty() bool {
	return c.CurrencyIDTo == 0
}
func (c ConvertCurrencyRequest) IsBothCurrencyIDIdentical() bool {
	return c.CurrencyIDFrom == c.CurrencyIDTo
}

type ConvertCurrencyResponse struct {
	Result int64  `json:"result"`
	Error  string `json:"error,omitempty"`
}
