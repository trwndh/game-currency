package http_response

import (
	"net/http"

	"github.com/go-chi/render"
)

type SuccessResponse struct {
	Status     string `json:"status"`
	httpStatus int
}

func HTTPSuccessResponse(status string, httpStatusCode int, w http.ResponseWriter, r *http.Request) {

	resp := SuccessResponse{status, httpStatusCode}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

func (s SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(s.httpStatus)
	return nil
}
