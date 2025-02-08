package testerr

// TestObject is the interface of the type that must be provided when calling
// the Assert or Require functions of the Result type.
type TestObject interface {
	Helper()
	Log(...interface{})
	Logf(string, ...interface{})
	Fail()
	FailNow()
}

// Result
type Result struct {
	Checker Checker
}

// Nil returns a Result which checks that the received error is nil.
func Nil() Result {
	return Result{
		Checker: NilChecker{},
	}
}

// As returns a Result which checks that the received error is convertable to
// an error of the same type as given.
//
// The value given is expected to be an error type which can be copied into
// the underlying Checker such that it works with errors.As.
func As(expected any) Result {
	return Result{
		Checker: AsErrChecker{
			Expected: &expected,
		},
	}
}

// Is returns a Result which checks that the received error is the same as the
// given error.
func Is(expected error) Result {
	return Result{
		Checker: IsErrChecker{
			Expected: expected,
		},
	}
}

// Check calls the Check function of the underlying Checker.
//
// If the underlying Checker is nil, this function will use a NilChecker to
// provide the Check function.
func (r Result) Check(actual error) error {
	if r.Checker == nil {
		return NilChecker{}.Check(actual)
	}
	return r.Checker.Check(actual)
}

// Assert validates the given error is what was expected, and if not fails the
// test but allows it to continue.
//
// If extra arguments are provided they are treated as a format string and
// arguments to pass to `t.Logf`.
func (r Result) Assert(t TestObject, actual error, msgAndArgs ...interface{}) {
	t.Helper()
	if err := r.Checker.Check(actual); err != nil {
		t.Log(err.Error())
		LogMsgAndArgs(t, msgAndArgs)
		t.Fail()
	}
}

// Require validates the given error is what was expected, and if not fails and
// stops the current test.
//
// If extra arguments are provided they are treated as a format string and
// arguments to pass to `t.Logf`
func (r Result) Require(t TestObject, actual error, msgAndArgs ...interface{}) {
	t.Helper()
	if err := r.Checker.Check(actual); err != nil {
		t.Log(err.Error())
		LogMsgAndArgs(t, msgAndArgs)
		t.FailNow()
	}
}

// LogMsgAndArgs takes an array of values and logs it to the given testing
// instance.
func LogMsgAndArgs(t TestObject, msgAndArgs []interface{}) {
	t.Helper()
	switch len(msgAndArgs) {
	case 0:
		// Do nothing, no message/args
	case 1:
		t.Log(msgAndArgs...)
	default:
		fmtstr, ok := msgAndArgs[0].(string)
		if ok {
			// If first element is a string, assume it's a Printf like format
			// string
			t.Logf(fmtstr, msgAndArgs[1:]...)
		} else {
			t.Log(msgAndArgs...)
		}
	}
}
