package testerr

import (
	"errors"
	"fmt"
	"reflect"
)

// Checker is a type which allows checking the state of an error.
//
// If the error given does not match the expected state of the Checker, then
// the checker must return a new error indicating how it deviated from the
// expected state.
type Checker interface {
	Check(actual error) error
}

// NilChecker is a Checker which validates that the received error is nil.
type NilChecker struct{}

// Check validates that the given error is nil.
func (c NilChecker) Check(actual error) error {
	if actual != nil {
		return fmt.Errorf("%w: expected nil but got %#v", ErrUnexpectedValue, actual.Error())
	}
	return nil
}

// IsErrChecker is a Checker which validates the received error is an expected
// error.
type IsErrChecker struct {
	Expected error
}

// Check validates that the given error is the expected error.
//
// This uses the errors.Is function to check the given error; this means
// that it will properly handle wrapped errors.
func (c IsErrChecker) Check(actual error) error {
	if actual == nil {
		return NilValueError{}
	}
	if errors.Is(actual, c.Expected) {
		return nil
	}
	//nolint:errorlint // We have 3 error values and explicitly only want to wrap *one*
	return fmt.Errorf("%w: %#v != %#v", ErrUnexpectedValue, c.Expected, actual)
}

// AsErrChecker is a Checker which validates the received error is convertable
// to the given error type.
//
// The Expected value must be a pointer to a type which implements the error
// interface.
type AsErrChecker[T error] struct{}

// Check validates that the given error is convertable to the expected error
// type.
//
// This uses the errors.As function to check that the error is able to be
// converted.
func (c AsErrChecker[T]) Check(actual error) error {
	if actual == nil {
		return NilValueError{}
	}

	var dst T
	if errors.As(actual, &dst) {
		return nil
	}
	// Return a conversion error with the given type name
	return ConversionError{
		TypeName: reflect.ValueOf(dst).Elem().Type().Name(),
	}
}
