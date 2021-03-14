package http_response

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	httpStatus int
}

func HTTPErrorResponse(err error, httpStatusCode int, w http.ResponseWriter, r *http.Request) {

	resp := ErrorResponse{err.Error(), httpStatusCode}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.httpStatus)
	return nil
}
