package main

import (
	"fmt"
	"strings"
)

func main() {
	// ARRAYS

	// var nums [3]int = [3]int{21, 42, 84}
	// var nums = [3]int{21, 42, 84}
	nums := [3]int{21, 42, 84}

	fmt.Println(nums)

	// SLICES

	names := []string{"Tom", "Sam", "John"}
	fmt.Println(names)

	names[1] = "Harry"
	names = append(names, "Tony")

	fmt.Println(names)

	// slice ranges

	rangeOne := names[1:3]
	fmt.Println(rangeOne)

	// USING STRINGS PACKAGE
	sentence := "Hello there! Game on.."

	fmt.Println(strings.Contains(sentence, "Game"))
	fmt.Println(strings.Contains(sentence, "game"))

	fmt.Println(strings.ReplaceAll(sentence, "Hello", "Hi"))
	fmt.Println(sentence)

	x := 0

	for x < 5 {
		fmt.Println("X is", x)
		x++
	}

	for x := 0; x < 5; x++ {
		fmt.Println("x is:", x)
	}

	values := []int{5, 4, 3, 2, 1}

	for iterator := 0; iterator < len(values); iterator++ {
		fmt.Println(values[iterator])
	}

	for index, value := range values {
		fmt.Printf("Value at index %v is %v\n", index, value)
	}
}
