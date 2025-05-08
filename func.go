package assert

import (
	"runtime"
	"testing"
)

type Assert[Type any] func(t T, got Type)

func (a Assert[Type]) Want(t *testing.T, got Type) {
	t.Helper()
	tc := wantT{t}
	a(tc, got)
}

func (a Assert[Type]) Ensure(t *testing.T, got Type) {
	t.Helper()
	tc := ensureT{t}
	a(tc, got)
}

func Func[Type any](passing func(t T, got Type)) Assert[Type] {
	return func(t T, got Type) {
		t.Helper()
		passing(t, got)
	}
}

func AllOf[Type any](asserts ...Assert[Type]) Assert[Type] {
	return func(t T, got Type) {
		t.Helper()
		for _, assert := range asserts {
			assert(t, got)
		}
	}
}

func Len[Type any](want int) Assert[[]Type] {
	_, file, line, _ := runtime.Caller(1)
	return Func(func(t T, got []Type) {
		t.Helper()
		if len(got) != want {
			t.Checkf(file, line, "got len(%v)=%d, want: %d", got, len(got), want)
		}
	})
}

func Contain[Type comparable](want Type) Assert[[]Type] {
	_, file, line, _ := runtime.Caller(1)
	return Func(func(t T, got []Type) {
		t.Helper()
		for _, v := range got {
			if v == want {
				return
			}
		}

		t.Checkf(file, line, "got %v, want contain %v", got, want)
	})
}
