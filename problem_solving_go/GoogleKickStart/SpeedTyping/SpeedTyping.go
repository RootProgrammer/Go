package main

import "fmt"

func matchString(i, p string) bool {
	frequency1 := [26]int{}
	frequency2 := [26]int{}

	count := 0

	fmt.Print(i)

	for _, char := range i {
		// fmt.Print(frequency1[char-'a'], "\n")
		frequency1[char-'a']++
	}
	for _, char := range p {
		frequency2[char-'a']++
	}

	for j := 0; j < 26; j++ {
		count += min(frequency1[j], frequency2[j])
	}

	return count
}


func main() {
	var given, typed string

	fmt.Scanln(&given)
	fmt.Scanln(&typed)

	fmt.Print(matchString(given, typed))
}
