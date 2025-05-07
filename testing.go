package assert

import "testing"

type T interface {
	testing.TB
	Check(msg ...any)
	Checkf(format string, v ...any)
}

type wantT struct {
	*testing.T
}

func (t wantT) Check(msg ...any) {
	t.Helper()
	t.Error(msg...)
}

func (t wantT) Checkf(format string, v ...any) {
	t.Helper()
	t.Errorf(format, v...)
}

type ensureT struct {
	*testing.T
}

func (t ensureT) Check(msg ...any) {
	t.Helper()
	t.Fatal(msg...)
}

func (t ensureT) Checkf(format string, v ...any) {
	t.Helper()
	t.Fatalf(format, v...)
}
