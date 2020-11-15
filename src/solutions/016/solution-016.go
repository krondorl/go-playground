package main

import (
	"fmt"
)

func head(arr []int) int {
	return arr[0]
}

func last(arr []int) int {
	return arr[len(arr)-1]
}

func tail(arr []int) []int {
	return arr[1:]
}

func sum(arr []int) int {
	if len(arr) == 1 {
		return head(arr)
	}

	return head(arr) + sum(tail(arr))
}

func maximum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	x := head(arr)
	maxTail := maximum(tail(arr))

	if x > maxTail {
		return x
	}

	return maxTail
}

func factorial(n int) int {
	return n * factorial(n-1)
}

func applyTwice(n int, f func(int) int) int {
	return f(f(n))
}

func double(n int) int {
	return 2 * n
}

func main() {
	var exampleArray = []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Functions in FP style")
	fmt.Println("")
	fmt.Println("exampleArray = ", exampleArray)
	fmt.Println("head(exampleArray) = ", head(exampleArray))
	fmt.Println("last(exampleArray) = ", last(exampleArray))
	fmt.Println("tail(exampleArray) = ", tail(exampleArray))
	fmt.Println("sum(exampleArray) = ", sum(exampleArray))
	fmt.Println("maximum(exampleArray) = ", maximum(exampleArray))
	fmt.Println("applyTwice(1, double) = ", applyTwice(1, double))
	fmt.Println("applyTwice(2, double) = ", applyTwice(2, double))
	fmt.Println("applyTwice(3, double) = ", applyTwice(3, double))
}
