package main

import (
	"fmt"
	"strconv"
)

func generateFizzBuzz(maxValue int, fizz, buzz int) []string {
	fizzBuzzSlice := []string{}

	for i := 1; i <= maxValue; i++ {
		if i%fizz == 0 && i%buzz == 0 {
			fizzBuzzSlice = append(fizzBuzzSlice, "fizzBuzz")
			continue
		} else if i%fizz == 0 {
			fizzBuzzSlice = append(fizzBuzzSlice, "fizz")
			continue
		} else if i%buzz == 0 {
			fizzBuzzSlice = append(fizzBuzzSlice, "buzz")
			continue
		}

		fizzBuzzSlice = append(fizzBuzzSlice, strconv.Itoa(i))
	}

	return fizzBuzzSlice
}

func main() {
	fmt.Println("Generalized Fizzbuzz")
	fmt.Println("")
	fmt.Println("generateFizzBuzz(45, 2, 7) ", generateFizzBuzz(45, 2, 7))
	fmt.Println("")
	fmt.Println("generateFizzBuzz(45, 3, 5) ", generateFizzBuzz(45, 3, 5))
	fmt.Println("")
	fmt.Println("generateFizzBuzz(100, 3, 11) ", generateFizzBuzz(100, 3, 11))
}
