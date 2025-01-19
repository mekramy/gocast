package gocast

import (
	"fmt"
	"strings"
)

const errorNil = "value is nil"
const errorType = "cannot convert value to type"
const errorOverflow = "value is out of range for"

func nilErr() error {
	return fmt.Errorf(errorNil)
}

func typeError(t string) error {
	return fmt.Errorf("%s %s", errorType, t)
}

func overflowError(t string) error {
	return fmt.Errorf("%s %s", errorOverflow, t)
}

// IsNilError checks if the provided error is a bil error.
// It returns true if the error is not nil and its nil error.
func IsNilError(err error) bool {
	return err != nil && err.Error() == errorNil
}

// IsCastError checks if the provided error is a casting error.
// It returns true if the error is not nil and its a casting error.
func IsCastError(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), errorType)
}

// IsOverflowError checks if the provided error is a overflow error.
// It returns true if the error is not nil and its a overflow error.
func IsOverflowError(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), errorType)
}
