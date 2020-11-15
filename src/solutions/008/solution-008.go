package main

import (
	"fmt"
)

// https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
func primeFactors(n int) (pfs []int) {
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	if n > 2 {
		pfs = append(pfs, n)
	}

	pfs = unique(pfs)

	return
}

// https://www.cloudhadoop.com/2018/12/golang-example-program-to-check-number.html
func isPrime(number int) bool {
	isPrime := true

	if number == 0 || number == 1 {
		isPrime = false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}

	return isPrime
}

// https://www.cloudhadoop.com/2018/12/golang-example-calculate-sum-of-digits.html
func sumDigits(number int) int {
	remainder := 0
	sumResult := 0

	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}

	return sumResult
}

// https://www.cloudhadoop.com/2018/12/golang-example-program-to-sum-elements.html
func sum(array []int) int {
	result := 0

	for _, v := range array {
		result += v
	}

	return result
}

func sumDigitsOfArray(array []int) int {
	result := 0

	for _, v := range array {
		result += sumDigits(v)
	}

	return result
}

func isHoax(number int) bool {
	if !isPrime(number) {
		var sumDigitsOfPrimeFactors = sumDigitsOfArray(primeFactors(number))
		var sumDigits = sumDigits(number)
		if sumDigitsOfPrimeFactors == sumDigits {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Hoax number")
	fmt.Println("22", isHoax(22))
	fmt.Println("84", isHoax(84))
	fmt.Println("19", isHoax(19))
}
