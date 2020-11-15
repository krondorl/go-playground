package main

import (
	"fmt"
)

func fibonacciLoop(items int) []int {
	fibonumbers := []int{1, 1}

	for i := 2; i <= items-1; i++ {
		prev2 := fibonumbers[i-2]
		prev := fibonumbers[i-1]
		sum := prev2 + prev
		fibonumbers = append(fibonumbers, sum)
	}

	return fibonumbers
}

func fibonacciRecursive(n int) []int {
	if n < 2 {
		return []int{1}
	}

	if n < 3 {
		return []int{1, 1}
	}

	array := fibonacciRecursive(n - 1)
	array = append(array, array[n-2]+array[n-3])

	return array
}

func main() {
	fmt.Println("Fibonacci")
	fmt.Println("")
	fmt.Println("fibonacciLoop(10) ", fibonacciLoop(10))
	fmt.Println("fibonacciRecursive(10) ", fibonacciRecursive(10))
}
