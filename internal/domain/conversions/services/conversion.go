package services

import (
	"github.com/trwndh/game-currency/internal/domain/conversions"
	"github.com/trwndh/game-currency/internal/domain/conversions/repositories"
)

type service struct {
	ConversionRepo repositories.Conversion
}

func NewConversion(ConversionRepo repositories.Conversion) conversions.Service {
	return service{ConversionRepo: ConversionRepo}
}
