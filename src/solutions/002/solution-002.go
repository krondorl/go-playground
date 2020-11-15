package main

import "fmt"

func mod3(x int) bool {
	if x%3 == 0 {
		return true
	}

	return false
}

func mod5(x int) bool {
	if x%5 == 0 {
		return true
	}

	return false
}

func multiples(x int) int {
	if x < 1 {
		return 0
	}

	var mults []int
	var sum int

	for i := 1; i < x; i++ {
		if mod3(i) || mod5(i) {
			mults = append(mults, i)
		}
	}

	for i := 1; i < len(mults); i++ {
		sum += mults[i]
	}

	return sum
}

func main() {
	fmt.Println("Multiples of 3 and 5")
	fmt.Println(multiples(1000))
}
