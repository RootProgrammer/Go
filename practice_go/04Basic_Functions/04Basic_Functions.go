/*
	A function can take zero or more arguments.

	In this example, add takes two parameters of type int.

	Notice that the type comes after the variable name.
*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
