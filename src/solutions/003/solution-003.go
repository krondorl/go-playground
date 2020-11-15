package main

import "fmt"

func factorial(n int) int {
	if n > 1 {
		return n * factorial(n-1)
	}

	return 1
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Factorial and Fibonacci")
	fmt.Println("factorial(3) = ", factorial(3))
	fmt.Println("fibonacci(3) = ", fibonacci(3))
}
