package curry

import (
	"errors"
	"reflect"

	"jsouthworth.net/go/dyn"
)

type curriedFunc struct {
	fn    interface{}
	numIn int
}

func (f curriedFunc) Apply(args ...interface{}) interface{} {
	if len(args) < f.numIn {
		return func(x interface{}) interface{} {
			return dyn.Apply(f, append(args, x)...)
		}
	}
	return dyn.Apply(f.fn, args...)

}
func Curry(f interface{}) interface{} {
	fnt := reflect.TypeOf(f)
	if fnt.Kind() != reflect.Func {
		panic(errors.New("can only curry functions"))
	}
	return curriedFunc{fn: f, numIn: fnt.NumIn()}
}
