package main

import (
	"fmt"
	"math"
)

func convertDecimalToBinary(decimal int) int {
	binary, counter := 0, 1

	for decimal != 0 {
		reminder := decimal % 2
		decimal /= 2
		binary += reminder * counter
		counter *= 10
	}

	return binary
}

func reverseBinary(binary int) int {
	reversed_binary := 0

	for binary != 0 {
		reminder := binary % 10
		reversed_binary *= 10
		reversed_binary += reminder
		reversed_binary /= 10
	}

	return reversed_binary
}

func convertBinaryToDecimal(binary int) int {
	decimal, counter := 0, 0

	for binary != 0 {
		reminder := binary % 10
		decimal += reminder * int(math.Pow(2.0, float64(counter)))
		binary /= 10
		counter++
	}

	return decimal
}

func main() {
	var x, y int
	fmt.Scanf("%v %v", &x, &y)

	binary := convertDecimalToBinary(x)
	binary0 := binary * 10
	binary1 := binary0 + 1
	reversed_binary0 := reverseBinary(binary0)
	reversed_binary1 := reverseBinary(binary1)
	decimal0 := convertBinaryToDecimal(reversed_binary0)
	decimal1 := convertBinaryToDecimal(reversed_binary1)

	if decimal0 == y {
		fmt.Println("YES")
	} else if decimal1 == y {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
