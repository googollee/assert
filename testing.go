package assert

import (
	"fmt"
	"path/filepath"
	"testing"
)

type T interface {
	testing.TB
	Check(file string, line int, msg ...any)
	Checkf(file string, line int, format string, v ...any)
}

type wantT struct {
	*testing.T
}

func (t wantT) Check(file string, line int, msg ...any) {
	t.Helper()
	t.Error(append([]any{definedInfo(file, line)}, msg...)...)
}

func (t wantT) Checkf(file string, line int, format string, v ...any) {
	t.Helper()
	t.Errorf(definedInfo(file, line)+format, v...)
}

type ensureT struct {
	*testing.T
}

func (t ensureT) Check(file string, line int, msg ...any) {
	t.Helper()
	t.Fatal(append([]any{definedInfo(file, line)}, msg...)...)
}

func (t ensureT) Checkf(file string, line int, format string, v ...any) {
	t.Helper()
	t.Fatalf(definedInfo(file, line)+format, v...)
}

func definedInfo(file string, line int) string {
	if file == "" {
		return ""
	}

	return fmt.Sprintf("Defined at %s:%d: ", filepath.Base(file), line)
}
