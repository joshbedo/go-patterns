package phone_number

import (
	"facade/phone_validator"
	"fmt"
)

type ValidatorService interface {
	IsValidNumber(number string)
}

type Service struct {
	lengthChecker        phone_validator.LengthChecker
	prefixValidator      phone_validator.PrefixValidator
	countryCodeValidator phone_validator.CountryCodeValidator
}

func NewService(
	lengthChecker phone_validator.LengthChecker,
	prefixValidator phone_validator.PrefixValidator,
	countryCodeValidator phone_validator.CountryCodeValidator,
) *Service {
	return &Service{
		lengthChecker:        lengthChecker,
		prefixValidator:      prefixValidator,
		countryCodeValidator: countryCodeValidator,
	}
}

func (service Service) IsValidNumber(number string) {
	fmt.Println("number: ", number)

	validCountry := service.countryCodeValidator.IsValid(number)
	fmt.Println("validCountry: ", validCountry)

	validPrefix := service.prefixValidator.IsValid(number)
	fmt.Println("validPrefix: ", validPrefix)

	validLength := service.lengthChecker.IsValid(number)
	fmt.Println("validLength: ", validLength)
}
