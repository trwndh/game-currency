package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/trwndh/game-currency/internal/domain/currencies/dto"

	"github.com/trwndh/game-currency/internal/handler/http/gen"
	"github.com/trwndh/game-currency/internal/instrumentation/loggers"
	"github.com/trwndh/game-currency/internal/server/http/http_response"
	"go.uber.org/zap"
)

func (h HttpServer) CreateCurrency(w http.ResponseWriter, r *http.Request, params gen.CreateCurrencyParams) {
	ctx := r.Context()

	var bodyFromRequest gen.CreateCurrencyJSONRequestBody
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		loggers.Bg().Error("Error Read Body from request at handler.CreateCurrency", zap.Error(err))
		http_response.HTTPErrorResponse(err, 500, w, r)
		return
	}

	err = json.Unmarshal(body, &bodyFromRequest)
	if err != nil {
		loggers.Bg().Error("Error Unmarshall body to bodyFromRequest at handler.GetCurrency", zap.Error(err))
		http_response.HTTPErrorResponse(err, 500, w, r)
		return
	}

	createResponse, err := h.currencyService.Create(ctx, dto.CreateCurrencyRequest{Name: bodyFromRequest.Name})
	if err != nil {
		loggers.Bg().Error("Error Unmarshall body to bodyFromRequest at handler.GetCurrency", zap.Error(err))
		http_response.HTTPErrorResponse(err, http.StatusUnprocessableEntity, w, r)
		return
	}

	http_response.HTTPSuccessResponse(createResponse.Status, 201, w, r)

}
