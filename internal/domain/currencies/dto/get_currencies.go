package dto

type Currency struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetCurrenciesRequest struct{}

type GetCurrenciesResponse struct {
	Currencies []Currency `json:"currencies"`
	Error      string     `json:"error,omitempty"`
}
