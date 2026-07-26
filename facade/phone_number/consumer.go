package phone_number

import (
	"facade/phone_validator"
)

type Consumer struct {
	validator ValidatorService
}

func NewConsumer() *Consumer {
	lengthChecker := phone_validator.NewLengthChecker(10)
	prefixChecker := phone_validator.NewPrefixValidator()
	countryCodeValidator := phone_validator.NewCountryCodeValidator("44")

	validator := NewService(*lengthChecker, *prefixChecker, *countryCodeValidator)

	return &Consumer{validator: validator}
}

func (consumer Consumer) Process(number string) {
	consumer.validator.IsValidNumber(number)
}
