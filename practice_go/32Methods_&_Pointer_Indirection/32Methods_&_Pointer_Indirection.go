/*
	Comparing the previous two programs,
	you might notice that functions with a pointer argument must take a pointer:

	var v Vertex
	ScaleFunc(v, 5)  // Compile error!
	ScaleFunc(&v, 5) // OK
	while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

	var v Vertex
	v.Scale(5)  // OK
	p := &v
	p.Scale(10) // OK
	For the statement v.Scale(5), even though v is a value and not a pointer,
	the method with the pointer receiver is called automatically. That is, as a convenience,
	Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
*/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/*
	The equivalent thing happens in the reverse direction.

	Functions that take a value argument must take a value of that specific type:

	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // Compile error!
	while methods with value receivers take either a value
	or a pointer as the receiver when they are called:

	var v Vertex
	fmt.Println(v.Abs()) // OK
	p := &v
	fmt.Println(p.Abs()) // OK
	In this case, the method call p.Abs() is interpreted as (*p).Abs().

	package main

	import (
		"fmt"
		"math"
	)

	type Vertex struct {
		X, Y float64
	}

	func (v Vertex) Abs() float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	func AbsFunc(v Vertex) float64 {
		return math.Sqrt(v.X*v.X + v.Y*v.Y)
	}

	func main() {
		v := Vertex{3, 4}
		fmt.Println(v.Abs())
		fmt.Println(AbsFunc(v))

		p := &Vertex{4, 3}
		fmt.Println(p.Abs())
		fmt.Println(AbsFunc(*p))
	}

*/
