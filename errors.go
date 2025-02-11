package testerr

import (
	"fmt"

	"github.com/tvarney/consterr"
)

const (
	// ErrUnexpectedValue is an error which indicates that the error being
	// checked did not match the expected value.
	ErrUnexpectedValue consterr.Error = "received error does not match expected state"
)

// NilValueError is an error indicating that the received error was nil when a
// non-nil error was expected.
type NilValueError struct{}

func (e NilValueError) Error() string {
	return "expected non-nil error but received nil"
}

func (e NilValueError) Unwrap() error {
	return ErrUnexpectedValue
}

// NonNilValueError is an error indicating that the received error was non-nil
// when nil was expected.
type NonNilValueError struct {
	ErrStr string
}

// NewNonNilValueError returns a canonical version of the NonNilValueError for
// the given received error.
func NewNonNilValueError(err error) NonNilValueError {
	return NonNilValueError{
		ErrStr: fmt.Sprintf("%#v", err),
	}
}

func (e NonNilValueError) Error() string {
	msg := "expected nil error but received non-nil value"
	if e.ErrStr != "" {
		msg += " " + e.ErrStr
	}
	return msg
}

func (e NonNilValueError) Unwrap() error {
	return ErrUnexpectedValue
}

// ConversionError is an error indicating that the received error was not able
// to be converted to the expected type.
type ConversionError struct {
	TypeName string
}

func (e ConversionError) Error() string {
	return "received error is not convertable to "
}

func (e ConversionError) Unwrap() error {
	return ErrUnexpectedValue
}
