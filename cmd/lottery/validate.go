package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// validateNonPromptCount checks if the count is valid when the count-prompt flag is not set.
func validateNonPromptCount() error {
	if viper.GetBool("count-prompt") {
		return nil
	}

	return validateCount(viper.GetInt("count"))
}

// validateCount checks if the count is valid.
func validateCount(count int) error {
	if count < 1 {
		return fmt.Errorf("count must be greater than 0")
	}

	return nil
}

// parseCount parses the count from a string and validates it.
func parseCount(raw string) (int, error) {
	count, err := strconv.Atoi(strings.TrimSpace(raw))
	if err != nil {
		return 0, fmt.Errorf("count must be a whole number")
	}

	if err := validateCount(count); err != nil {
		return 0, err
	}

	return count, nil
}
