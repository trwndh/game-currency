package http

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"github.com/trwndh/game-currency/internal/server/http/httperr"
	"go.uber.org/zap"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"

	"github.com/trwndh/game-currency/internal/handler/http/gen"
)

func (h HttpServer) GetConversionAmount(w http.ResponseWriter, r *http.Request, params gen.GetConversionAmountParams) {
	ctx := r.Context()

	result, err := h.conversionService.ConvertCurrency(ctx, dto.ConvertCurrencyRequest{
		CurrencyIDFrom: params.CurrencyIdFrom,
		CurrencyIDTo:   params.CurrencyIdTo,
		Amount:         params.Amount,
	})
	if err != nil {
		loggers.Bg().Error("Error service handler.GetConversionAmount", zap.Error(err))
		httperr.HTTPErrorResponse(err, 500, w, r)
		return
	}

	var response gen.ConversionReturn
	response.Result = result.Result

	render.JSON(w, r, response)

}
