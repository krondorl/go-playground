package main

import "fmt"

func sumOfSquares() int {
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSums() int {
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum * sum
}

func main() {
	fmt.Println("Sum square difference")
	fmt.Println(squareOfSums() - sumOfSquares())
}
