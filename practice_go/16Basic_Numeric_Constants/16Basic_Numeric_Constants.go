/*
	Numeric constants are high-precision values.

	An untyped constant takes the type needed by its context.

	Try printing needInt(Big) too.

	(An int can store at maximum a 64-bit integer, and sometimes less.)
*/

package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*
	<< is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.

	So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".

	For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.
*/
