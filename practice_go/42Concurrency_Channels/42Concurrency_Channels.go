/*
	Channels are a typed conduit through which you can send and receive values
	with the channel operator, <-.

	ch <- v    // Send v to channel ch.
	v := <-ch  // Receive from ch, and
			// assign value to v.
	(The data flows in the direction of the arrow.)

	Like maps and slices, channels must be created before use:

	ch := make(chan int)
	By default, sends and receives block until the other side is ready.
	This allows goroutines to synchronize without explicit locks or condition variables.

	The example code sums the numbers in a slice, distributing the work between two goroutines.
	Once both goroutines have completed their computation, it calculates the final result.
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
	Buffered Channels
	Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

	ch := make(chan int, 100)
	Sends to a buffered channel block only when the buffer is full.
	Receives block when the buffer is empty.

	Modify the example to overfill the buffer and see what happens.

	package main

	import "fmt"

	func main() {
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}

*/
