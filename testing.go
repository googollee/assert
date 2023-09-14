package assert

type Testing interface {
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
}
