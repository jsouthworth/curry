// Package curry provides function currying for go.
package curry // import "jsouthworth.net/go/curry"

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
		return func(xs ...interface{}) interface{} {
			return dyn.Apply(f, append(args, xs...)...)
		}
	}
	return dyn.Apply(f.fn, args...)

}

// Curry takes a function and allows it to be called lazily until the
// arguments are all populated. The curried functions take arbitrary
// arguments instead of just one so that lazy population can occur in
// steps larger than one. This change from normal currying
// is simply a conveinence.
func Curry(f interface{}) interface{} {
	fnt := reflect.TypeOf(f)
	if fnt.Kind() != reflect.Func {
		panic(errors.New("can only curry functions"))
	}
	return curriedFunc{fn: f, numIn: fnt.NumIn()}.Apply()
}
