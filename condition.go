package assert

import (
	"errors"
	"fmt"
)

type Condition[Type any] interface {
	Constrain(v Type) bool
	String() string
}

type condition[Type any] struct {
	fn  func(v Type) bool
	str string
}

func Func[Type any](str string, fn func(v Type) bool) Condition[Type] {
	return condition[Type]{
		fn:  fn,
		str: str,
	}
}

func (c condition[Type]) Constrain(v Type) bool {
	return c.fn(v)
}

func (c condition[Type]) String() string {
	return c.str
}

func OneOf[Type any](conds ...Condition[Type]) Condition[Type] {
	fn := func(got Type) bool {
		if len(conds) == 0 {
			return true
		}

		for _, cond := range conds {
			if cond.Constrain(got) {
				return true
			}
		}

		return false
	}

	strs := make([]string, 0, len(conds))
	for _, cond := range conds {
		strs = append(strs, cond.String())
	}

	return Func(fmt.Sprintf("one of %v", strs), fn)
}

func AllOf[Type any](conds ...Condition[Type]) Condition[Type] {
	fn := func(got Type) bool {
		for _, cond := range conds {
			if !cond.Constrain(got) {
				return false
			}
		}
		return true
	}

	strs := make([]string, 0, len(conds))
	for _, cond := range conds {
		strs = append(strs, cond.String())
	}

	return Func(fmt.Sprintf("all of %v", strs), fn)
}

func Contain[Type comparable](want Type) Condition[[]Type] {
	fn := func(got []Type) bool {
		for _, v := range got {
			if v == want {
				return true
			}
		}

		return false
	}
	return Func(fmt.Sprintf("contain %v", want), fn)

}

func Len[Type any](want int) Condition[[]Type] {
	return Func(fmt.Sprintf("len() == %d", want), func(got []Type) bool { return len(got) == want })
}

func Any[Type any]() Condition[Type] {
	return Func("any", func(got Type) bool { return true })
}

func Equal[Type comparable](want Type) Condition[Type] {
	return Func(fmt.Sprintf("== %v", want), func(got Type) bool { return got == want })
}

func IsNil[Type any]() Condition[Type] {
	return Func("is nil", func(got Type) bool { return any(got) == nil })
}

func IsError(want error) Condition[error] {
	return Func(fmt.Sprintf("is %v", want), func(got error) bool { return errors.Is(got, want) })
}
