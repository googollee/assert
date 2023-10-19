package assert

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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

	tests := []struct {
		name        string
		checking    func(t T)
		wantFailed  bool
		wantPaniced bool
		wantOutput  string
	}{
		{
			name:        "CheckOK",
			checking:    func(t T) { a.Check(t, 10, "should not be logged") },
			wantFailed:  false,
			wantPaniced: false,
			wantOutput:  "",
		},
		{
			name:        "CheckFail",
			checking:    func(t T) { a.Check(t, 1, "should be logged") },
			wantFailed:  true,
			wantPaniced: false,
			wantOutput:  "should be logged\nat asserter_test.go:50: diff (-got, +want):\n  int(\n- \t1,\n+ \t10,\n  )\n\n",
		},
		{
			name:        "EnsureOK",
			checking:    func(t T) { a.Ensure(t, 10, "should not be logged") },
			wantFailed:  false,
			wantPaniced: false,
			wantOutput:  "",
		},
		{
			name:        "EnsureFail",
			checking:    func(t T) { a.Ensure(t, 1, "should be logged") },
			wantFailed:  true,
			wantPaniced: true,
			wantOutput:  "should be logged\nat asserter_test.go:50: diff (-got, +want):\n  int(\n- \t1,\n+ \t10,\n  )\n\n",
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var ft fakeTesting
			defer func() {
				v := recover()
				if got, want := v != nil, tc.wantPaniced; got != want {
					t.Fatalf("recover() = %v, want panic: %v", v, tc.wantPaniced)
				}

				if got, want := ft.failed, tc.wantFailed; got != want {
					t.Fatalf("ft.failed = %v, want: %v", got, want)
				}
				output := uniqueSpace(ft.output.String())
				if diff := cmp.Diff(output, tc.wantOutput); diff != "" {
					t.Fatalf("output diff (-got, +want):\n%s", diff)
				}
			}()

			tc.checking(&ft)
		})
	}
}
