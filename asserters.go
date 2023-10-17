package assert

import (
	"errors"
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Any[T any](conds ...Assert[T]) Assert[T] {
	return Assert[T]{
		check: func(got T) string {
			if len(conds) == 0 {
				return ""
			}

			outs := make([]string, 0, len(conds))
			for _, cond := range conds {
				if out := cond.check(got); out != "" {
					outs = append(outs, cond.check(got))
				} else {
					return ""
				}
			}

			return fmt.Sprintf("all fails: %v", outs)
		},
	}
}

func All[T any](conds ...Assert[T]) Assert[T] {
	return Assert[T]{
		check: func(got T) string {
			outs := make([]string, 0, len(conds))
			for _, cond := range conds {
				if out := cond.check(got); out != "" {
					outs = append(outs, cond.check(got))
				}
			}

			if len(outs) == 0 {
				return ""
			}

			return fmt.Sprintf("contain failures: %v", outs)
		},
	}
}

func Contain[T comparable](want T) Assert[[]T] {
	return Assert[[]T]{
		check: func(got []T) string {
			for _, g := range got {
					if g == want {
						return ""
				}
			}

			return fmt.Sprintf("%+v doesn't contain %+v", got, want)
		},
	}
}

func Len[T any](want int) Assert[[]T] {
	return Assert[[]T]{
		check: func(got []T) string {
			if len(got) == want {
				return ""
			}

			return fmt.Sprintf("len(%+v) = %d, want: %d", got, len(got), want)
		},
	}
}

func Func[T any](fn func(got T) bool) Assert[T] {
	return Assert[T]{
		check: func(got T) string {
			if fn(got) {
				return ""
			}

			return fmt.Sprintf("fails: %v", got)
		},
	}
}

func Equal[T comparable](want T) Assert[T] {
	return Assert[T]{
		check: func(got T) string {
			diff := cmp.Diff(got, want)
			if diff == "" {
				return ""
			}
			return "diff (-got, +want):\n" + diff
		},
	}
}

func IsNil[T comparable]() Assert[T] {
	return Assert[T]{
		check: func(got T) string {
			if any(got) == nil {
				return ""
			}

			return fmt.Sprintf("got: %+v, want: nil", got)
		},
	}
}

func IsError(want error) Assert[error] {
	return Assert[error]{
		check: func(got error) string {
			if errors.Is(got, want) {
				return ""
			}

			return fmt.Sprintf("got: %+v, want: %+v", got, want)
		},
	}
}
