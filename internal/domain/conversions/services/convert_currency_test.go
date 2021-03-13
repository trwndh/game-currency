package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/trwndh/game-currency/internal/domain/conversions/errors"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/trwndh/game-currency/internal/domain/conversions/dto"
	"github.com/trwndh/game-currency/internal/domain/conversions/entity"
	"github.com/trwndh/game-currency/internal/domain/conversions/repositories/mocks"
)

func Test_service_ConvertCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	conversion.EXPECT().FindRate(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
	}).Return(entity.ConversionRate{
		CurrencyIDFrom: 2,
		CurrencyIDTo:   1,
		Rate:           29,
	}, nil)

	c := NewService(conversion)

	got, err := c.ConvertCurrency(context.Background(), dto.ConvertCurrencyRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Amount:         580,
	})
	expectedReturn := dto.ConvertCurrencyResponse{
		Result: 20,
	}

	assert.Equal(t, got.Result, expectedReturn.Result)
	assert.NoError(t, err)
}

func Test_service_ConvertCurrency_Error_Payload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	c := NewService(conversion)

	_, err := c.ConvertCurrency(context.Background(), dto.ConvertCurrencyRequest{
		CurrencyIDFrom: 0,
		CurrencyIDTo:   0,
		Amount:         580,
	})

	assert.Error(t, err, "error Test_service_ConvertCurrency_Error_Payload")
}

func Test_service_ConvertCurrency_Error_SqlNoRows_FindRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	conversion.EXPECT().FindRate(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
	}).Return(entity.ConversionRate{}, sql.ErrNoRows)

	c := NewService(conversion)

	_, err := c.ConvertCurrency(context.Background(), dto.ConvertCurrencyRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Amount:         580,
	})

	assert.Error(t, err)
}

func Test_service_ConvertCurrency_Error_Database_FindRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	conversion.EXPECT().FindRate(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
	}).Return(entity.ConversionRate{}, errors.GetErrorInvalidCurrencyID())

	c := NewService(conversion)

	_, err := c.ConvertCurrency(context.Background(), dto.ConvertCurrencyRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Amount:         580,
	})

	assert.Error(t, err)
}
