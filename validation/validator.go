package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

var (
	supportedNetwork = []string{"MTN", "AIRTEL", "GLO", "9MOBILE"}
)

var (
	isUsernameValid = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
)

func ValidateString(value string, min int, max int) error {
	n := len(value)

	if n < min || n > max {
		return fmt.Errorf("must contain from %d-%d characters", min, max)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isUsernameValid(value) {
		return fmt.Errorf("must contain only letters, digits or underscore")
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

func ValidateWalletAddress(value string) error {
	if ok := common.IsHexAddress(value); !ok {
		return fmt.Errorf("not a valid address")
	}
	return nil
}

func ValidateId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("value must be a positive integer")
	}
	return nil
}

func ValidateType(value string) error {
	tradeTypes := []string{"buy", "sell"}
	for _, tradeType := range tradeTypes {
		if value == tradeType {
			return nil
		}
	}
	return fmt.Errorf("value must be either 'buy' or 'sell', but got '%s'", value)
}
