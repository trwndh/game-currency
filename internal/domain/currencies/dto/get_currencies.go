package dto

type Currency struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetCurrenciesRequest struct{}

type GetCurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
	Error      string     `json:"error,omitempty"`
}
