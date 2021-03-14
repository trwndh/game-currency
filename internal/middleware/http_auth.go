package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/trwndh/game-currency/internal/server/http/http_response"

	"github.com/trwndh/game-currency/internal/handler/http/gen"

	"github.com/trwndh/game-currency/config"
)

func CheckSecretKey(cfg *config.MainConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			authString := r.Header.Get("Authorization")

			var errResponse gen.GenericError
			unAuthorizedString := "unauthorized"
			errResponse.Error = &unAuthorizedString

			err := errors.New(unAuthorizedString)

			if authString == "" {
				http_response.HTTPErrorResponse(err, 401, w, r)
				return
			}

			authArray := strings.Split(authString, " ")
			if len(authArray) < 2 {
				http_response.HTTPErrorResponse(err, 401, w, r)
				return
			}

			if authArray[0] != "Basic" {
				http_response.HTTPErrorResponse(err, 401, w, r)
				return
			}

			if authArray[1] == "" {
				http_response.HTTPErrorResponse(err, 401, w, r)
				return
			}

			if authArray[1] != cfg.Secret.SecretKey {
				http_response.HTTPErrorResponse(err, 401, w, r)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
