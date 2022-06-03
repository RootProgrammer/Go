/*
	User Input in Go
	Used to accept text input from the user.


	Syntax
	stores space separated values into successive arguments
	var storageVariable variableType
	fmt.Scan(&storageVariable) //assuming fmt is imported

	reads line all in one go
	reader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
	var storageString string
	storageString, _ := reader.ReadString('\n')

	Notes
	The Scan function is part of "fmt",
	so make sure that package is imported.
	Fmt also comes with Scanf (for specifying string formatting) and Scanln (for scanning until the end of the line).
	The Scan functions store to a pointer variable.

	The scan function itself returns the number of successfully scanned items and if necessary, an error (in that order).
	It is good practice to error check when using the Scan functions[Scan, Scanf, Scanln].

	The Scan functions are used for splitting space-delimited tokens, whereas the reader is used to read full lines.


*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//reading an integer
	var age int
	fmt.Println("What is your age?")
	_, err := fmt.Scan(&age)

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	// var name string
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')

	fmt.Println("Your name is ", name, "and your age is ", age)
}
