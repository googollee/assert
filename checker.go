package cc

import (
	"fmt"
	"path/filepath"
	"testing"
)

type Checker[Type any] func(t checker, got Type)

func (c Checker[Type]) Want(t *testing.T, got Type) {
	t.Helper()
	wt := wantT{t}
	c(wt, got)
}

func (c Checker[Type]) Ensure(t *testing.T, got Type) {
	t.Helper()
	et := ensureT{t}
	c(et, got)
}

type checker interface {
	testing.TB
	Check(file string, line int, msg ...any)
	Checkf(file string, line int, format string, v ...any)

	testing() *testing.T
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

func (t wantT) testing() *testing.T {
	return t.T
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

func (t ensureT) testing() *testing.T {
	return t.T
}

func definedInfo(file string, line int) string {
	if file == "" {
		return ""
	}

	return fmt.Sprintf("Defined at %s:%d: ", filepath.Base(file), line)
}
