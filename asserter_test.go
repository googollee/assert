package assert

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

var panicFailed = errors.New("failed")

type fakeTesting struct {
	output bytes.Buffer
	failed bool
}

func (t *fakeTesting) Helper() {}
func (t *fakeTesting) Log(args ...any) {
	fmt.Fprintln(&t.output, args...)
}
func (t *fakeTesting) Logf(format string, args ...any) {
	fmt.Fprintf(&t.output, format, args...)
	fmt.Fprintln(&t.output)
}
func (t *fakeTesting) Error(args ...any) {
	t.Log(args...)
	t.failed = true
}
func (t *fakeTesting) Errorf(format string, args ...any) {
	t.Logf(format, args...)
	t.failed = true
}
func (t *fakeTesting) Fatal(args ...any) {
	t.Error(args...)
	panic(panicFailed)
}
func (t *fakeTesting) Fatalf(format string, args ...any) {
	t.Errorf(format, args...)
	panic(panicFailed)
}

func uniqueSpace(str string) string {
	return strings.ReplaceAll(str, "\u00a0", " ")
}

func TestEqual(t *testing.T) {
	a := Equal[int](10)

	{
		var ft fakeTesting
		a.Check(&ft, 10, "should not be logged")
		if ft.failed {
			t.Fatal("a.Check(10) should not fail")
		}
		if ft.output.Len() != 0 {
			t.Fatal("a.Check(10) should not fail")
		}
	}

	{
		var ft fakeTesting
		a.Check(&ft, 1, "should be logged")
		if !ft.failed {
			t.Fatal("a.Check(1) should fail")
		}
		if got, want := uniqueSpace(ft.output.String()), "should be logged\nat asserter_test.go:48: diff (-got, +want):\n  int(\n- \t1,\n+ \t10,\n  )\n\n"; got != want {
			t.Fatalf("a.Check(1) output: %q, want: %q", got, want)
		}
	}

	{
		var ft fakeTesting
		defer func() {
			if v := recover(); v != nil {
				t.Fatal("a.Ensure(10) should not panic")
			}
			if ft.failed {
				t.Fatal("a.Check(10) should not fail")
			}
			if ft.output.Len() != 0 {
				t.Fatal("a.Check(10) should not fail")
			}
		}()
		a.Ensure(&ft, 10, "should not be logged")
	}

	{
		var ft fakeTesting
		defer func() {
			if v := recover(); v != panicFailed {
				t.Fatal("a.Ensure(1) should panic, got:", v)
			}
			if !ft.failed {
				t.Fatal("a.Check(1) should fail")
			}
			if got, want := uniqueSpace(ft.output.String()), "should be logged\nat asserter_test.go:48: diff (-got, +want):\n  int(\n- \t1,\n+ \t10,\n  )\n\n"; got != want {
				t.Fatalf("a.Check(1) output: %q, want: %q", got, want)
			}
		}()
		a.Ensure(&ft, 1, "should be logged")
	}
}
