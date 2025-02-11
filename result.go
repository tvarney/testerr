package testerr

// TestObject is the interface of the type that must be provided when calling
// the Assert or Require functions of the ExpectedError type.
type TestObject interface {
	Helper()
	Log(msgAndArgs ...interface{})
	Logf(fmtstr string, args ...interface{})
	Fail()
	FailNow()
}

// ExpectedError is a wrapper around a Checker which provides assertions which
// may be used in tests.
type ExpectedError struct {
	Checker Checker
}

// Nil returns a ExpectedError which checks that the received error is nil.
func Nil() ExpectedError {
	return ExpectedError{
		Checker: NilChecker{},
	}
}

// As returns a ExpectedError which checks that the received error is convertable to
// an error of the same type as given.
//
// The value given is expected to be an error type which can be copied into
// the underlying Checker such that it works with errors.As.
func As[T error]() ExpectedError {
	return ExpectedError{
		Checker: AsErrChecker[T]{},
	}
}

// Is returns a ExpectedError which checks that the received error is the same as the
// given error.
func Is(expected error) ExpectedError {
	return ExpectedError{
		Checker: IsErrChecker{
			Expected: expected,
		},
	}
}

// Check calls the Check function of the underlying Checker.
//
// If the underlying Checker is nil, this function will use a NilChecker to
// provide the Check function.
func (r ExpectedError) Check(actual error) error {
	if r.Checker == nil {
		return NilChecker{}.Check(actual)
	}
	//nolint:wrapcheck // We intentionally don't wrap this (it makes no sense to)
	return r.Checker.Check(actual)
}

// Assert fails the test if the given error isn't as expected.
//
// This function will fail the current test if the given error doesn't match
// the expected error. The test will be allowed to continue past the failure.
//
// If extra arguments are provided they are treated as a format string and
// arguments to pass to `t.Logf`.
func (r ExpectedError) Assert(t TestObject, actual error, msgAndArgs ...interface{}) {
	t.Helper()
	if err := r.Checker.Check(actual); err != nil {
		t.Log(err.Error())
		LogMsgAndArgs(t, msgAndArgs)
		t.Fail()
	}
}

// Require fails and stops the test if the given error isn't as expected.
//
// This function will fail the current test if the given error doesn't match
// the expected error and stops test execution at that point. This will
// prevent any further assertions in the test from being evaluated.
//
// If extra arguments are provided they are treated as a format string and
// arguments to pass to `t.Logf`.
func (r ExpectedError) Require(t TestObject, actual error, msgAndArgs ...interface{}) {
	t.Helper()
	if err := r.Checker.Check(actual); err != nil {
		t.Log(err.Error())
		LogMsgAndArgs(t, msgAndArgs)
		t.FailNow()
	}
}

// LogMsgAndArgs takes an array of values and logs it to the given testing
// instance.
func LogMsgAndArgs(tObj TestObject, msgAndArgs []interface{}) {
	tObj.Helper()
	switch len(msgAndArgs) {
	case 0:
		// Do nothing, no message/args
	case 1:
		tObj.Log(msgAndArgs...)
	default:
		fmtstr, ok := msgAndArgs[0].(string)
		if ok {
			// If first element is a string, assume it's a Printf like format
			// string
			tObj.Logf(fmtstr, msgAndArgs[1:]...)
		} else {
			tObj.Log(msgAndArgs...)
		}
	}
}
