package curry

import (
	"fmt"

	"jsouthworth.net/go/dyn"
)

func ExampleCurry_twoArgs() {
	curriedAdd := Curry(func(x, y int) int { return x + y })
	add1 := dyn.Apply(curriedAdd, 1)
	out := dyn.Apply(add1, 2)
	fmt.Println(out)
	// Output: 3
}

func ExampleCurry_threeArgs() {
	curriedAdd := Curry(func(x, y, z int) int { return x + y + z })
	add1 := dyn.Apply(curriedAdd, 1)
	add3 := dyn.Apply(add1, 2)
	out := dyn.Apply(add3, 4)
	fmt.Println(out)
	// Output: 7
}

func ExampleCurry_allArgsAtOnce() {
	curriedAdd := Curry(func(x, y, z int) int { return x + y + z })
	out := dyn.Apply(curriedAdd, 1, 2, 4)
	fmt.Println(out)
	// Output: 7
}

func ExampleCurry_threeArgsTwoAtOnce() {
	curriedAdd := Curry(func(x, y, z int) int { return x + y + z })
	add3 := dyn.Apply(curriedAdd, 1, 2)
	out := dyn.Apply(add3, 4)
	fmt.Println(out)
	// Output: 7
}

func ExampleCurry_variadic() {
	curriedAdd := Curry(func(x, y, z int, rest ...int) int {
		out := x + y + z
		for _, e := range rest {
			out += e
		}
		return out
	})
	out := dyn.Apply(curriedAdd, 1, 2, 3, 4, 5)
	fmt.Println(out)
	// Output: 15
}
