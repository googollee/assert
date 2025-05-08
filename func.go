package cc

import (
	"runtime"
	"testing"
)

func Func[Type any](passing func(t *testing.T, got Type)) Checker[Type] {
	return func(t checker, got Type) {
		t.Helper()
		passing(t.testing(), got)
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
