package testerr

import (
	"github.com/tvarney/consterr"
)

const (
	ErrAny            consterr.Error = "anyerr"
	ErrExpectedNil    consterr.Error = "expected nil"
	ErrExpectedNonNil consterr.Error = "expected non-nil error but got nil"
	ErrBadMatch       consterr.Error = "expected error does not match actual error"
)
