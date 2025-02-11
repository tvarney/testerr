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
func (e ExpectedError) Check(actual error) error {
	if e.Checker == nil {
		return NilChecker{}.Check(actual)
	}
	//nolint:wrapcheck // We intentionally don't wrap this (it makes no sense to)
	return e.Checker.Check(actual)
}

// Assert fails the test if the given error isn't what was expected.
//
// This function will fail the current test if the given error doesn't match
// the expected error. The test will be allowed to continue past the failure.
//
// If extra parameters are provided they are logged via `t.Log()`. If format
// string evaluation is expected, use ExpectedError.Assertf instead of this.
func (e ExpectedError) Assert(tObj TestObject, actual error, logExtra ...interface{}) {
	tObj.Helper()
	if err := e.Checker.Check(actual); err != nil {
		tObj.Log(err.Error())
		if len(logExtra) > 0 {
			tObj.Log(logExtra...)
		}
		tObj.Fail()
	}
}

// Assertf fails the test if the given error isn't what was expected.
//
// This function will fail the current test if the given error doesn't match
// the expected error. The test will be allowed to continue past the failure.
func (e ExpectedError) Assertf(tObj TestObject, actual error, fmtstr string, args ...interface{}) {
	tObj.Helper()
	if err := e.Checker.Check(actual); err != nil {
		tObj.Log(err.Error())
		tObj.Logf(fmtstr, args...)
		tObj.Fail()
	}
}

// Require fails and stops the test if the given error isn't what was expected.
//
// This function will fail the current test if the given error doesn't match
// the expected error and stops test execution at that point. This will
// prevent any further assertions in the test from being evaluated.
//
// If extra parameters are provided they are logged via `t.Log()`. If format
// string evaluation is expected, use ExpectedError.Requiref instead of this.
func (e ExpectedError) Require(tObj TestObject, actual error, logExtra ...interface{}) {
	tObj.Helper()
	if err := e.Checker.Check(actual); err != nil {
		tObj.Log(err.Error())
		if len(logExtra) > 0 {
			tObj.Log(logExtra...)
		}
		tObj.FailNow()
	}
}

// Requiref fails and stops the test if the given error isn't what was
// expected.
//
// This function will fail the current test if the given error doesn't match
// the expected error and stops test execution at that point. This will
// prevent any further assertions in the test from being evaluated.
func (e ExpectedError) Requiref(tObj TestObject, actual error, fmtstr string, args ...interface{}) {
	tObj.Helper()
	if err := e.Checker.Check(actual); err != nil {
		tObj.Log(err.Error())
		tObj.Logf(fmtstr, args...)
		tObj.FailNow()
	}
}
