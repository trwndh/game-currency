package http

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/trwndh/game-currency/internal/handler/http/gen"

	"github.com/trwndh/game-currency/internal/server/http/http_response"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"go.uber.org/zap"
)

func (h HttpServer) GetCurrency(w http.ResponseWriter, r *http.Request, params gen.GetCurrencyParams) {
	ctx := r.Context()

	result, err := h.currencyService.Find(ctx)
	if err != nil {
		loggers.Bg().Error("Error service handler.GetCurrency", zap.Error(err))
		http_response.HTTPErrorResponse(err, 500, w, r)
		return
	}

	var response gen.CurrencyList
	if len(result.Currencies) == 0 {
		response.Data = make([]gen.Currency, 0)
	}
	for _, val := range result.Currencies {
		response.Data = append(response.Data, gen.Currency{
			Id:   val.ID,
			Name: val.Name,
		})
	}

	render.JSON(w, r, response)

}
