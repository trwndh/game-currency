package errors

import "errors"

func GetErrorInvalidPayload() error {
	return errors.New("error invalid payload")
}

//
//func GetErrorConversionAlreadyExist() error {
//	return errors.New("error this conversion already exist")
//}
//
//func GetErrorConversionNotFound() error {
//	return errors.New("error no conversion for these currency")
//}

func GetErrorCurrencyAlreadyExist() error {
	return errors.New("error currency already exist")
}
func GetErrorDatabase() error {
	return errors.New("database error")
}
