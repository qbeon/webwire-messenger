package validator

import "fmt"

// Error represents a validator error
type Error struct {
	err error
}

// Error implements the standard Go error interface
func (err Error) Error() string {
	return err.err.Error()
}

// Errorf creates a new formatted validator error
func Errorf(format string, v ...interface{}) Error {
	return Error{
		err: fmt.Errorf(format, v...),
	}
}

// IsValidatorError returns true if the given error represents a validation
// error and originates from a validator module
func IsValidatorError(err error) bool {
	_, isValidatorErr := err.(Error)
	return isValidatorErr
}
