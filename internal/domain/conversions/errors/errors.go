package errors

import "errors"

func GetErrorInvalidPayload() error {
	return errors.New("error invalid payload")
}

func GetErrorConversionAlreadyExist() error {
	return errors.New("error this conversion already exist")
}

func GetErrorConversionNotFound() error {
	return errors.New("error no conversion for these currency")
}

func GetErrorDatabase() error {
	return errors.New("database error")
}

func GetErrorRateIsZero() error {
	return errors.New("rate cannot be 0")
}

func GetErrorConvertingSameID() error {
	return errors.New("error cannot create same currency conversion")
}

func GetErrorInvalidCurrencyID() error {
	return errors.New("error invalid currency id")
}

func Get1452Error() error {
	return errors.New("error code: 1452. Cannot add or update a child row: a foreign key constraint fails ")
}
