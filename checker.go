package testerr

import (
	"errors"
	"fmt"
)

// Checker
type Checker interface {
	Check(err error) error
}

// NilChecker is a Checker which validates that the received error is nil.
type NilChecker struct{}

// Check validates that the given error is nil.
func (c NilChecker) Check(actual error) error {
	if actual != nil {
		return fmt.Errorf("%w: %s", ErrExpectedNil, actual.Error())
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
// This uses the `errors.Is` function to check the given error; this means
// that it will properly handle wrapped errors.
func (c IsErrChecker) Check(actual error) error {
	if actual == nil {
		return ErrExpectedNonNil
	}
	if errors.Is(actual, c.Expected) {
		return nil
	}
	return fmt.Errorf("%w: %s != %s", ErrBadMatch, c.Expected.Error(), actual.Error())
}

// AsErrChecker is a Checker which validates the received error is convertable
// to the given error type.
type AsErrChecker struct {
	Expected any
}

func (c AsErrChecker) Check(actual error) error {
	if actual == nil {
		return ErrExpectedNonNil
	}
	if errors.As(actual, c.Expected) {
		return nil
	}
	return fmt.Errorf("%w: %s is not convertable to %T", ErrBadMatch, actual.Error(), c.Expected)
}
