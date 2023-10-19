package assert

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type T interface {
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

type Assert[TWant any] struct {
	check     func(v TWant) (diff string)
	definedIn string
	definedAt int
}

func newAssert[TWant any](skip int, check func(v TWant) (diff string)) Assert[TWant] {
	_, file, line, _ := runtime.Caller(skip + 1)
	if file != "" {
		file = filepath.Base(file)
	}
	return Assert[TWant]{
		check:     check,
		definedIn: file,
		definedAt: line,
	}
}

func (a Assert[TWant]) atLine() string {
	if a.definedIn == "" {
		return ""
	}
	return fmt.Sprintf("at %s:%d: ", a.definedIn, a.definedAt)
}

func (a Assert[TWant]) Check(t T, v TWant, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Log(msg...)
	t.Error(a.atLine() + out)
}

func (a Assert[TWant]) Checkf(t T, v TWant, format string, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Logf(format, msg...)
	t.Error(a.atLine() + out)
}

func (a Assert[TWant]) Ensure(t T, v TWant, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Log(msg...)
	t.Fatal(a.atLine() + out)
}

func (a Assert[TWant]) Ensuref(t T, v TWant, format string, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Logf(format, msg...)
	t.Fatal(a.atLine() + out)
}
