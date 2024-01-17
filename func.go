package assert

import (
	"errors"
	"fmt"
	"strings"
)

type Assert[Type any] func(got Type) string

func Func[Type any](passing func(got Type) bool, shouldMsg string) Assert[Type] {
	return func(got Type) string {
		if passing(got) {
			return ""
		}

		return shouldMsg
	}
}

func AllOf[Type any](asserts ...Assert[Type]) Assert[Type] {
	return func(got Type) string {
		var failed []string

		for _, assert := range asserts {
			if f := assert(got); f != "" {
				failed = append(failed, f)
			}
		}

		if len(failed) == 0 {
			return ""
		}

		return fmt.Sprintf("[%s]", strings.Join(failed, ", "))
	}
}

func OneOf[Type any](asserts ...Assert[Type]) Assert[Type] {
	return func(got Type) string {
		var failed []string

		for _, assert := range asserts {
			f := assert(got)
			if f == "" {
				return ""
			}

			failed = append(failed, f)
		}

		if len(failed) == 0 {
			return ""
		}

		return fmt.Sprintf("[%s]", strings.Join(failed, ", "))
	}
}

func Any[Type any]() Assert[Type] {
	return Func(func(got Type) bool {
		return true
	}, "any value")
}

func Equal[Type comparable](want Type) Assert[Type] {
	return Func(func(got Type) bool {
		return got == want
	}, fmt.Sprintf("%v", want))
}

func IsNil[Type any]() Assert[Type] {
	return Func(func(got Type) bool {
		return any(got) == nil
	}, "nil")
}

func IsError(err error) Assert[error] {
	return Func(func(got error) bool {
		return errors.Is(got, err)
	}, fmt.Sprintf("error(%v)", err))
}

func Len[Type any](want int) Assert[[]Type] {
	return Func(func(got []Type) bool {
		return len(got) == want
	}, fmt.Sprintf("len()==%d", want))
}

func Contain[Type comparable](want Type) Assert[[]Type] {
	return Func(func(got []Type) bool {
		for _, v := range got {
			if v == want {
				return true
			}
		}

		return false
	}, fmt.Sprintf("contain %v", want))
}
