package phone_validator

import "strings"

type CountryCodeValidator struct {
	countryCode string
}

func NewCountryCodeValidator(countryCode string) *CountryCodeValidator {
	return &CountryCodeValidator{countryCode: countryCode}
}

// Check country code after prefix (+ or 00)
func (validator CountryCodeValidator) IsValid(number string) bool {
	nationalNumber := trimInternationalPrefix(number)
	if strings.HasPrefix(nationalNumber, validator.countryCode) {
		return true
	}

	return false
}

func trimInternationalPrefix(number string) string {
	validPrefixes := []string{"+", "00"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(number, prefix) {
			return strings.TrimPrefix(number, prefix)
		}
	}

	return number
}
