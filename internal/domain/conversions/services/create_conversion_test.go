package services

import (
	"context"
	errorsPackage "errors"
	"testing"

	"github.com/trwndh/game-currency/internal/domain/conversions/errors"

	"github.com/stretchr/testify/assert"

	"github.com/trwndh/game-currency/internal/domain/conversions/dto"

	"github.com/trwndh/game-currency/internal/domain/conversions/repositories/mocks"

	"github.com/golang/mock/gomock"
)

func Test_service_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)
	conversion.EXPECT().CountExistingConversion(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           29,
	}).Return(int64(0), nil)

	conversion.EXPECT().Create(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           29,
	}).Return(nil)

	c := NewService(conversion)

	got, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           29,
	})

	expectedReturn := dto.CreateConversionResponse{
		Status: "success",
	}

	assert.Equal(t, got.Status, expectedReturn.Status)
	assert.NoError(t, err)
}

func Test_service_Create_Error_Empty_CurrencyID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	c := NewService(conversion)

	_, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 0,
		CurrencyIDTo:   0,
		Rate:           20,
	})

	assert.Equal(t, err, errors.GetErrorInvalidPayload())
	assert.Error(t, err)
}

func Test_service_Create_Error_Empty_Rate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	c := NewService(conversion)

	_, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 10,
		CurrencyIDTo:   20,
		Rate:           0,
	})

	assert.Equal(t, err, errors.GetErrorRateIsZero())
	assert.Error(t, err)
}

func Test_service_Create_Error_Identical_CurrencyID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)

	c := NewService(conversion)

	_, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 10,
		CurrencyIDTo:   10,
		Rate:           20,
	})

	assert.Equal(t, err, errors.GetErrorConvertingSameID())
	assert.Error(t, err)
}

func Test_service_Create_Error_CountExistingConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)
	conversion.EXPECT().CountExistingConversion(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	}).Return(int64(0), errors.GetErrorDatabase())
	c := NewService(conversion)

	_, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	})

	assert.Equal(t, err, errors.GetErrorDatabase())
	assert.Error(t, err)
}

func Test_service_Create_Error_ConversionAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)
	conversion.EXPECT().CountExistingConversion(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	}).Return(int64(1), nil)
	c := NewService(conversion)

	_, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	})

	assert.Equal(t, err, errors.GetErrorConversionAlreadyExist())
	assert.Error(t, err)
}

func Test_service_Create_Error_CreateRepo_1452(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)
	conversion.EXPECT().CountExistingConversion(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	}).Return(int64(0), nil)

	conversion.EXPECT().Create(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	}).Return(errors.Get1452Error())

	c := NewService(conversion)

	_, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	})

	assert.Equal(t, err, errors.GetErrorDatabase())
	assert.Error(t, err)
}

func Test_service_Create_Error_CreateRepo_AnotherError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	conversion := mocks.NewMockConversion(ctrl)
	conversion.EXPECT().CountExistingConversion(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	}).Return(int64(0), nil)

	conversion.EXPECT().Create(gomock.Any(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	}).Return(errorsPackage.New("another error"))

	c := NewService(conversion)

	resp, err := c.Create(context.Background(), dto.CreateConversionRequest{
		CurrencyIDFrom: 1,
		CurrencyIDTo:   2,
		Rate:           30,
	})
	expected := dto.CreateConversionResponse{
		Error: errors.GetErrorDatabase().Error(),
	}
	assert.Equal(t, resp.Error, expected.Error)
	assert.Equal(t, err, errors.GetErrorDatabase())
	assert.Error(t, err)
}
