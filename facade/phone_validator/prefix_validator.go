package phone_validator

import "strings"

type PrefixValidator struct {
}

func NewPrefixValidator() *PrefixValidator {
	return &PrefixValidator{}
}

func (validator PrefixValidator) IsValid(number string) bool {
	validPrefix := []string{"+", "00"}

	for _, prefix := range validPrefix {
		if strings.HasPrefix(number, prefix) {
			return true
		}
	}

	return false
}
