package main

import (
	"fmt"
)

var fibonacci = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

// map
func double(inputArray []int) []int {
	outputArray := inputArray

	for index, value := range outputArray {
		outputArray[index] = value * 2
	}

	return outputArray
}

// filter
func greaterThanFive(inputArray []int) []int {
	outputArray := make([]int, 0)

	for _, value := range inputArray {
		if value > 5 {
			outputArray = append(outputArray, value)
		}
	}

	return outputArray
}

// reduce
func sumValues(inputArray []int) int {
	sum := 0

	for _, value := range inputArray {
		sum += value
	}

	return sum
}

func main() {
	double := double(fibonacci)
	greaterThanFive := greaterThanFive(double)
	sum := sumValues(greaterThanFive)

	fmt.Println("Map, filter, reduce")
	fmt.Println("")
	fmt.Println(fibonacci)
	fmt.Println("")
	fmt.Println("double ", double)
	fmt.Println("greaterThanFive ", greaterThanFive)
	fmt.Println("sum ", sum)
}
