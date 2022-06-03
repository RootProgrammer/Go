/*
	A type switch is a construct that permits several type assertions in series.

	A type switch is like a regular switch statement,
	but the cases in a type switch specify types (not values),
	and those values are compared against the type of the value held by the given interface value.

	switch v := i.(type) {
	case T:
		// here v has type T
	case S:
		// here v has type S
	default:
		// no match; here v has the same type as i
	}
	The declaration in a type switch has the same syntax as a type assertion i.(T),
	but the specific type T is replaced with the keyword type.

	This switch statement tests whether the interface value i holds a value of type T or S.
	In each of the T and S cases, the variable v will be of type T or S respectively
	and hold the value held by i. In the default case (where there is no match),
	the variable v is of the same interface type and value as i.
*/

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
