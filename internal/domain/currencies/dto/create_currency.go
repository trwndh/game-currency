package dto

type CreateCurrencyRequest struct {
	Name string `json:"name"`
}

func (c CreateCurrencyRequest) IsNameEmpty() bool {
	return c.Name == ""
}

type CreateCurrencyResponse struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}
