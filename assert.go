package assert

import "testing"

type Assert[T any] struct {
	check func(v T) string
}

func (a Assert[T]) Check(t *testing.T, v T, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Log(msg...)
	t.Error(out)
}

func (a Assert[T]) Checkf(t *testing.T, v T, format string, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Logf(format, msg...)
	t.Error(out)
}

func (a Assert[T]) Ensure(t *testing.T, v T, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Log(msg...)
	t.Fatal(out)
}

func (a Assert[T]) Ensuref(t *testing.T, v T, format string, msg ...any) {
	t.Helper()

	out := a.check(v)
	if out == "" {
		return
	}

	t.Logf(format, msg...)
	t.Fatal(out)
}
