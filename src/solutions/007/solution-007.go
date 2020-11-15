package main

import (
	"fmt"
)

func partition(arr []int, low int, high int) int {
	var pivot int = arr[high]
	var i int = low - 1

	for j := low; j <= high-1; j++ {

		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return (i + 1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		var pi int = partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	fmt.Println("Quick sort")
	sampleArray := []int{3, 1, 4, 2, 10, 8, 6}
	fmt.Println("before ", sampleArray)

	quickSort(sampleArray, 0, 6)
	fmt.Println("after ", sampleArray)
}
