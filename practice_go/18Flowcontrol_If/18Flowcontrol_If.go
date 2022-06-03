// Go's if statements are like its for loops;
// the expression need not be surrounded by parentheses ( ) but the braces { } are required.

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

/*
	If with a short statement
	Like for, the if statement can start with a short statement to execute before the condition.

	Variables declared by the statement are only in scope until the end of the if.

	(Try using v in the last return statement.)

	package main

	import (
		"fmt"
		"math"
	)

	func pow(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}

	func main() {
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	}
*/

/*
	If and else
	Variables declared inside an if short statement are also available inside any of the else blocks.

	(Both calls to pow return their results before the call to fmt.Println in main begins.)

	package main

	import (
		"fmt"
		"math"
	)

	func pow(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("%g >= %g\n", v, lim)
		}
	can't use v here, though
		return lim
	}

	func main() {
		fmt.Println(
			pow(3, 2, 10),
			pow(3, 3, 20),
		)
	}
*/
