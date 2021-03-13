package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"

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

			if authString == "" {

				render.JSON(w, r, errResponse)
				return
			}

			authArray := strings.Split(authString, " ")
			if len(authArray) < 2 {

				render.JSON(w, r, errResponse)
				return
			}

			if authArray[0] != "Basic" {
				render.JSON(w, r, errResponse)
				return
			}

			if authArray[1] == "" {
				render.JSON(w, r, errResponse)
				return
			}

			if authArray[1] != cfg.Secret.SecretKey {
				render.JSON(w, r, errResponse)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
