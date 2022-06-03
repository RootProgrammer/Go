package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkSubStrings(jumbled_mess string, word_list []string, entry int) bool {
	is_complete_match := false
	match := 0

	for _, word := range word_list {
		if strings.Contains(jumbled_mess, word) {
			match += 1
		}
	}

	// fmt.Println(match)

	if match == entry {
		is_complete_match = true
	}

	return is_complete_match
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var entry int
	fmt.Scanln(&entry)
	reader.ReadString('\n')

	word_list := make([]string, entry)
	for i := 0; i < entry; i++ {
		fmt.Scanln(&word_list[i])
	}

	reader.ReadString('\n')

	var jumbled_mess string
	fmt.Scanln(&jumbled_mess)

	if checkSubStrings(jumbled_mess, word_list, entry) == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
