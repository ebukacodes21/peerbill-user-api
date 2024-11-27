package validation

import (
	"errors"
	"fmt"
)

var (
	supportedNetwork = []string{"MTN", "AIRTEL", "GLO", "9MOBILE"}
)

func ValidateString(value string, min int, max int) error {
	n := len(value)

	if n < min || n > max {
		return fmt.Errorf("must contain from %d-%d characters", min, max)
	}
	return nil
}

func ValidateFiat(value string) error {
	return ValidateString(value, 3, 4)
}

func ValidateCrypto(value string) error {
	return ValidateString(value, 3, 4)
}

func IsSupported(value string) error {
	for _, network := range supportedNetwork {
		if network == value {
			return nil
		}
	}

	return errors.New("network not supported")
}
