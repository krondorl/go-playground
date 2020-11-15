package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	var arrayLength int = len(arr)
	var swapped bool

	for {
		swapped = false
		for i := 0; i < arrayLength-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return arr
}

func main() {
	sampleArray := []int{3, 5, 4, 9, 10, 4, 6}
	fmt.Println("Bubble sort")

	fmt.Println("before ", sampleArray)

	fmt.Println("after ", bubbleSort(sampleArray))
}
