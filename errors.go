package assert

import (
	"errors"
	"strings"
)

func ErrorIs(t Testing, want, got error, format string, args ...any) {
	t.Helper()

	if errors.Is(got, want) {
		return
	}

	args = append([]any{want, got}, args...)
	t.Fatalf("want = %v, got = %v, "+format, args...)
}

func ErrorAs[Error error](t Testing, got error, format string, args ...any) {
	t.Helper()
	var want Error

	if errors.As(got, &want) {
		return
	}

	args = append([]any{want, got}, args...)
	t.Fatalf("want = %T, got = %T, "+format, args...)
}

func ErrorWithSubstr(t Testing, substr string, got error, format string, args ...any) {
	t.Helper()

	errStr := got.Error()
	if strings.Contains(errStr, substr) {
		return
	}

	args = append([]any{substr, errStr}, args...)
	t.Fatalf("want substr = %s, got = %s, "+format, args...)

}

func ErrorNil(t Testing, got error, format string, args ...any) {
	t.Helper()

	if got == nil {
		return
	}

	args = append([]any{got}, args...)
	t.Fatalf("want = nil, got = %v, "+format, args...)
}
