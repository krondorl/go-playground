package main

import (
	"fmt"
	"math"
)

func squareRoot(n int) int {
	return (int)(math.Sqrt(float64(n)))
}

func qMinusOne(n int) int {
	return (int)(math.Log(float64(n)) / math.Log(2))
}

func checkPerfect(n, pow int) bool {
	return n == int((math.Pow(2, float64(pow)) * (math.Pow(2, float64(pow)+1) - 1)))
}

func isPrime(n int) bool {
	var divisors int

	if n == 2 {
		return true
	}

	if n > 2 {
		for i := 2; i <= n; i++ {
			if n%i == 0 {
				divisors++
			}
		}
		if divisors == 1 {
			return true
		}
	}

	return false
}

func isEvenPerfect(n int) bool {
	root := squareRoot(n)
	pow := qMinusOne(root)

	if checkPerfect(n, pow) {
		if isPrime(pow + 1) {
			if isPrime(int(math.Pow(2, float64(pow+1)) - 1)) {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println("Even Perfect Number")
	fmt.Println("")
	fmt.Println("isEvenPerfect(6) = ", isEvenPerfect(6))
	fmt.Println("isEvenPerfect(156) = ", isEvenPerfect(156))
}
