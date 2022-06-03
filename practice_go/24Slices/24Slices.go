/*
	Slices
	An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
	flexible view into the elements of an array. In practice, slices are much more common than arrays.

	The type []T is a slice with elements of type T.

	A slice is formed by specifying two indices, a low and high bound, separated by a colon:

	a[low : high]
	This selects a half-open range which includes the first element, but excludes the last one.

	The following expression creates a slice which includes elements 1 through 3 of a:

	a[1:4]
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

/*
	nil slice
	They’re useful when you want to represent a slice that doesn’t exist, such as when an exception occurs in a function that returns a slice.

	Create a nil slice of integers.
	var slice []int

	empty slice
	Empty slices are useful when you want to represent an empty collection, such as when a database query returns zero results.

	Use make to create an empty slice of integers.
	slice := make([]int, 0)

	Use a slice literal to create an empty slice of integers.
	slice := []int{}
	Regardless of whether you’re using a nil slice or an empty slice, the built-in functions append, len, and cap work the same.
*/

/*
	Slices are like references to arrays
	A slice does not store any data, it just describes a section of an underlying array.

	Changing the elements of a slice modifies the corresponding elements of its underlying array.

	Other slices that share the same underlying array will see those changes.

	package main

	import "fmt"

	func main() {
		names := [4]string{
			"John",
			"Paul",
			"George",
			"Ringo",
		}
		fmt.Println(names)

		a := names[0:2]
		b := names[1:3]
		fmt.Println(a, b)

		b[0] = "XXX"
		fmt.Println(a, b)
		fmt.Println(names)
	}

*/

/*
	Slice literals
	A slice literal is like an array literal without the length.

	This is an array literal:

	[3]bool{true, true, false}
	And this creates the same array as above, then builds a slice that references it:

	[]bool{true, true, false}

	package main

	import "fmt"

	func main() {
		q := []int{2, 3, 5, 7, 11, 13}
		fmt.Println(q)

		r := []bool{true, false, true, true, false, true}
		fmt.Println(r)

		s := []struct {
			i int
			b bool
		}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println(s)
	}

*/

/*
	Slice defaults
	When slicing, you may omit the high or low bounds to use their defaults instead.

	The default is zero for the low bound and the length of the slice for the high bound.

	For the array

	var a [10]int
	these slice expressions are equivalent:

	a[0:10]
	a[:10]
	a[0:]
	a[:]

	package main

	import "fmt"

	func main() {
		s := []int{2, 3, 5, 7, 11, 13}

		s = s[1:4]
		fmt.Println(s)

		s = s[:2]
		fmt.Println(s)

		s = s[1:]
		fmt.Println(s)
	}

*/

/*
	Slice length and capacity
	A slice has both a length and a capacity.

	The length of a slice is the number of elements it contains.

	The capacity of a slice is the number of elements in the underlying array,
	counting from the first element in the slice.

	The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

	You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	Try changing one of the slice operations in the example program to extend it beyond its capacity
	and see what happens.

	package main

	import "fmt"

	func main() {
		s := []int{2, 3, 5, 7, 11, 13}
		printSlice(s)

		Slice the slice to give it zero length.
		s = s[:0]
		printSlice(s)

		Extend its length.
		s = s[:4]
		printSlice(s)

		Drop its first two values.
		s = s[2:]
		printSlice(s)
	}

	func printSlice(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

*/

/*
	Nil slices
	The zero value of a slice is nil.

	A nil slice has a length and capacity of 0 and has no underlying array.

	package main

	import "fmt"

	func main() {
		var s []int
		fmt.Println(s, len(s), cap(s))
		if s == nil {
			fmt.Println("nil!")
		}
	}

*/

/*
	Creating a slice with make
	Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

	The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5)  // len(a)=5
	To specify a capacity, pass a third argument to make:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4

	package main

	import "fmt"

	func main() {
		a := make([]int, 5)
		printSlice("a", a)

		b := make([]int, 0, 5)
		printSlice("b", b)

		c := b[:2]
		printSlice("c", c)

		d := c[2:5]
		printSlice("d", d)
	}

	func printSlice(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n",
			s, len(x), cap(x), x)
	}

*/

/*
	Slices of slices
	Slices can contain any type, including other slices.

	package main

	import (
		"fmt"
		"strings"
	)

	func main() {
		Create a tic-tac-toe board.
		board := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}

		The players take turns.
		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"

		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
	}

*/

/*
	Appending to a slice
	It is common to append new elements to a slice, and so Go provides a built-in append function.
	The documentation of the built-in package describes append.

	func append(s []T, vs ...T) []T
	The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

	The resulting value of append is
	a slice containing all the elements of the original slice plus the provided values.

	If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	The returned slice will point to the newly allocated array.

	package main

	import "fmt"

	func main() {
		var s []int
		printSlice(s)

		append works on nil slices.
		s = append(s, 0)
		printSlice(s)

		The slice grows as needed.
		s = append(s, 1)
		printSlice(s)

		We can add more than one element at a time.
		s = append(s, 2, 3, 4)
		printSlice(s)
	}

	func printSlice(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

*/

// for exercise, go to https://tour.golang.org/moretypes/18
