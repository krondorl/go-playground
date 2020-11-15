package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func randomNumbers(count int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	var p []int = rand.Perm(maxValue - 1)
	var output = make([]int, count)

	for i := 0; i < count; i++ {
		output[i] = p[i] + 1
	}

	sort.Ints(output)

	return output
}

func lottery5() []int {
	return randomNumbers(5, 90)
}

func lottery6() []int {
	return randomNumbers(6, 45)
}

func lottery7() []int {
	return randomNumbers(7, 35)
}

func main() {
	fmt.Println("Random lottery")
	fmt.Println("")
	fmt.Println("lottery5() ", lottery5())
	fmt.Println("lottery6() ", lottery6())
	fmt.Println("lottery7() ", lottery7())
}
