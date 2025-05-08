package cc

import (
	"errors"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Func[Type any](passing func(t testing.TB, got Type)) Checker[Type] {
	return func(t checker, got Type) {
		t.Helper()
		passing(t, got)
	}
}

func AllOf[Type any](asserts ...Checker[Type]) Checker[Type] {
	return func(t checker, got Type) {
		t.Helper()
		for _, assert := range asserts {
			assert(t, got)
		}
	}
}

func Equal[Type any](want Type) Checker[Type] {
	_, file, line, _ := runtime.Caller(1)
	return func(t checker, got Type) {
		t.Helper()
		diff := cmp.Diff(got, want)
		if diff == "" {
			return
		}

		t.Checkf(file, line, "diff (-got, +want):\n%s", cmp.Diff(got, want))
	}
}

func Len[Type any](want int) Checker[[]Type] {
	_, file, line, _ := runtime.Caller(1)
	return func(t checker, got []Type) {
		t.Helper()
		if len(got) != want {
			t.Checkf(file, line, "got len(%v)=%d, want: %d", got, len(got), want)
		}
	}
}

func Contain[Type comparable](want Type) Checker[[]Type] {
	_, file, line, _ := runtime.Caller(1)
	return func(t checker, got []Type) {
		t.Helper()
		for _, v := range got {
			if v == want {
				return
			}
		}

		t.Checkf(file, line, "got %v, want contain %v", got, want)
	}
}

func IsNil[Type any]() Checker[Type] {
	_, file, line, _ := runtime.Caller(1)
	return func(t checker, got Type) {
		t.Helper()
		var v any = got
		if v == nil {
			return
		}
		t.Checkf(file, line, "got: %v, want: nil", got)
	}
}

func IsError(want error) Checker[error] {
	_, file, line, _ := runtime.Caller(1)
	return func(t checker, got error) {
		t.Helper()
		if errors.Is(got, want) {
			return
		}
		t.Checkf(file, line, "got error: (%T)%v, want error: (%T)%v", got, got, want, want)
	}
}

func AsError[Want error]() Checker[error] {
	_, file, line, _ := runtime.Caller(1)
	return func(t checker, got error) {
		t.Helper()
		var want Want
		if errors.As(got, &want) {
			return
		}
		t.Checkf(file, line, "got error: (%T)%v, want as error: %T", got, got, want)
	}
}
