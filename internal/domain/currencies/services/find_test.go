package services

import (
	"context"
	"testing"

	"github.com/trwndh/game-currency/internal/domain/currencies/errors"

	"github.com/trwndh/game-currency/internal/domain/currencies/dto"

	"github.com/stretchr/testify/assert"

	"github.com/trwndh/game-currency/internal/domain/currencies/entity"

	"github.com/golang/mock/gomock"
	"github.com/trwndh/game-currency/internal/domain/currencies/repositories/mocks"
)

func TestService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currency := mocks.NewMockCurrency(ctrl)

	currencies := []entity.CurrencyDAO{
		{
			ID:   1,
			Name: "Knut",
		},
		{
			ID:   2,
			Name: "Sickle",
		},
		{
			ID:   3,
			Name: "Galleon",
		},
	}
	currency.EXPECT().Find(gomock.Any()).Return(currencies, nil)

	c := NewService(currency)

	got, err := c.Find(context.Background())
	dtoCurrency := []dto.Currency{
		{
			ID:   1,
			Name: "Knut",
		},
		{
			ID:   2,
			Name: "Sickle",
		},
		{
			ID:   3,
			Name: "Galleon",
		},
	}
	assert.Equal(t, got.Currencies, dtoCurrency)
	assert.NoError(t, err)

}

func TestService_Find_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currency := mocks.NewMockCurrency(ctrl)

	currency.EXPECT().Find(gomock.Any()).Return([]entity.CurrencyDAO{}, errors.GetErrorDatabase())

	c := NewService(currency)

	got, err := c.Find(context.Background())
	assert.Equal(t, got.Error, errors.GetErrorDatabase().Error())
	assert.Error(t, err)

}
