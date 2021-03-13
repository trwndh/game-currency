package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trwndh/game-currency/internal/domain/currencies/dto"
	"github.com/trwndh/game-currency/internal/domain/currencies/entity"

	"github.com/golang/mock/gomock"
	"github.com/trwndh/game-currency/internal/domain/currencies/repositories/mocks"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currency := mocks.NewMockCurrency(ctrl)
	currency.EXPECT().CountByName(gomock.Any(), "Knut").Return(int32(0), nil)
	currency.EXPECT().Create(gomock.Any(), entity.CurrencyDAO{Name: "Knut"}).Return(nil)
	c := NewService(currency)
	got, err := c.Create(context.Background(), dto.CreateCurrencyRequest{Name: "Knut"})

	assert.Equal(t, got.Status, "success")
	assert.NoError(t, err)
}

func TestService_Create_EmptyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currency := mocks.NewMockCurrency(ctrl)
	c := NewService(currency)
	got, err := c.Create(context.Background(), dto.CreateCurrencyRequest{Name: ""})

	assert.Equal(t, got.Error, "success")
	assert.NoError(t, err)
}
