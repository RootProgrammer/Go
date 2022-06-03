package main

import "fmt"

func main() {
	fmt.Print("Enter a year = ")
	var year int
	fmt.Scanln(&year)

	if (year%400 == 0) || (year%100 != 0 && year%4 == 0) {
		fmt.Printf("%v is Leap Year\n", year)
	} else {
		fmt.Printf("%v is not Leap Year\n", year)
	}
}
