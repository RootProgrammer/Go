/*
	An interface type is defined as a set of method signatures.

	A value of interface type can hold any value that implements those methods.

	Note: There is an error in the example code on line 22.
	Vertex (the value type) doesn't implement Abser
	because the Abs method is defined only on *Vertex (the pointer type).
*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser. uncomment to see the error
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
	Interfaces are implemented implicitly
	A type implements an interface by implementing its methods.
	There is no explicit declaration of intent, no "implements" keyword.

	Implicit interfaces decouple the definition of an interface from its implementation,
	which could then appear in any package without prearrangement.

	package main

	import "fmt"

	type I interface {
		M()
	}

	type T struct {
		S string
	}

	This method means type T implements the interface I,
	but we don't need to explicitly declare that it does so.
	func (t T) M() {
		fmt.Println(t.S)
	}

	func main() {
		var i I = T{"hello"}
		i.M()
	}

*/

/*
	Interface values
	Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

	(value, type)
	An interface value holds a value of a specific underlying concrete type.

	Calling a method on an interface value executes the method of the same name on its underlying type.

	package main

	import (
		"fmt"
		"math"
	)

	type I interface {
		M()
	}

	type T struct {
		S string
	}

	func (t *T) M() {
		fmt.Println(t.S)
	}

	type F float64

	func (f F) M() {
		fmt.Println(f)
	}

	func main() {
		var i I

		i = &T{"Hello"}
		describe(i)
		i.M()

		i = F(math.Pi)
		describe(i)
		i.M()
	}

	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

*/

/*
	Interface values with nil underlying values
	If the concrete value inside the interface itself is nil,
	the method will be called with a nil receiver.

	In some languages this would trigger a null pointer exception,
	but in Go it is common to write methods that gracefully handle being called with a nil receiver
	(as with the method M in this example.)

	Note that an interface value that holds a nil concrete value is itself non-nil.

	package main

	import "fmt"

	type I interface {
		M()
	}

	type T struct {
		S string
	}

	func (t *T) M() {
		if t == nil {
			fmt.Println("<nil>")
			return
		}
		fmt.Println(t.S)
	}

	func main() {
		var i I

		var t *T
		i = t
		describe(i)
		i.M()

		i = &T{"hello"}
		describe(i)
		i.M()
	}

	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

*/

/*
	Nil interface values
	A nil interface value holds neither value nor concrete type.

	Calling a method on a nil interface is a run-time error
	because there is no type inside the interface tuple to indicate which concrete method to call.

	package main

	import "fmt"

	type I interface {
		M()
	}

	func main() {
		var i I
		describe(i)
		i.M()
	}

	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

*/

/*
	The empty interface
	The interface type that specifies zero methods is known as the empty interface:

	interface{}
	An empty interface may hold values of any type. (Every type implements at least zero methods.)

	Empty interfaces are used by code that handles values of unknown type.
	For example, fmt.Print takes any number of arguments of type interface{}.

	package main

	import "fmt"

	func main() {
		var i interface{}
		describe(i)

		i = 42
		describe(i)

		i = "hello"
		describe(i)
	}

	func describe(i interface{}) {
		fmt.Printf("(%v, %T)\n", i, i)
	}

*/
