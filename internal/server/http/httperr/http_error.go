package httperr

import (
	"net/http"

	"github.com/go-chi/render"
)

func BadRequest(code string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(code, err, w, r, http.StatusBadRequest)
}

func httpRespondWithError(slug string, err error, w http.ResponseWriter, r *http.Request, status int) {
	resp := ErrorResponse{slug, err.Error(), status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.httpStatus)
	return nil
}
