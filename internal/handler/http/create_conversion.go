package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"

	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"github.com/trwndh/game-currency/internal/server/http/httperr"
	"go.uber.org/zap"

	"github.com/trwndh/game-currency/internal/handler/http/gen"
)

func (h HttpServer) CreateConversion(w http.ResponseWriter, r *http.Request, params gen.CreateConversionParams) {
	ctx := r.Context()

	bodyFromRequest := gen.ConversionRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		loggers.Bg().Error("Error Read Body from request at handler.CreateConversion", zap.Error(err))
		httperr.HTTPErrorResponse(err, 500, w, r)
		return
	}

	err = json.Unmarshal(body, &bodyFromRequest)
	if err != nil {
		loggers.Bg().Error("Error Read Body from request at handler.CreateConversion", zap.Error(err))
		httperr.HTTPErrorResponse(err, 500, w, r)
		return
	}

	createResponse, err := h.conversionService.Create(ctx, dto.CreateConversionRequest{
		CurrencyIDFrom: bodyFromRequest.CurrencyIdFrom,
		CurrencyIDTo:   bodyFromRequest.CurrencyIdTo,
		Rate:           bodyFromRequest.Rate,
	})
	if err != nil {
		loggers.Bg().Error("Error Read Body from request at handler.CreateConversion", zap.Error(err))
		httperr.HTTPErrorResponse(err, 422, w, r)
		return
	}

	var response gen.CreateSuccessReturn

	response.Status = &createResponse.Status
	render.JSON(w, r, response)
}
