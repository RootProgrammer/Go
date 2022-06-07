package main

import(
    "fmt"
    "strings"
)

func rotateRight(numbers []int, size int, k int) []int {
    new_numbers := make([]int, size)
    for index, value := range numbers {
        new_numbers[(index + k) % size] = value
    }

    return new_numbers
}

func main() {
    var test_case, size, k int

    fmt.Scanf("%v", &test_case)
	fmt.Scanln()
    for i := 0; i < test_case; i++ {
        fmt.Scanf("%v %v", &size, &k)
		fmt.Scanln()
        
        numbers := make([]int, size)
        for i := 0; i<size; i++ {
            fmt.Scanf("%v", &numbers[i])
        }

        result := rotateRight(numbers, size, k)

        fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
    }
}