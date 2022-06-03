/*
	Go's basic types are

	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte  alias for uint8

	rune		alias for int32	represents a Unicode code point

	float32 float64

	complex64 complex128
	The example shows variables of several types,
	and also that variable declarations may be "factored" into blocks, as with import statements.

	The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems
	and 64 bits wide on 64-bit systems.
	When you need an integer value you should use int unless you have a specific reason to use a sized
	or unsigned integer type.
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1 //shifting a 1 bit left 64 places.
	// In other words, the binary number that is 1 followed by 64 zeroes.
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
	<< is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.

	So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".

	For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.
*/
