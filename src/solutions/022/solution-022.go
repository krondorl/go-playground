package main

import (
	"fmt"
	"math"
)

func dogAge(ageYears float64) float64 {
	return 16*math.Log(ageYears) + 31
}

func main() {
	fmt.Println("Dog Age Calculator")
	fmt.Println("")
	fmt.Println("Alfy dog's human age is 13")
	fmt.Println("So his dog age is ", dogAge(13))
	fmt.Println("")
	fmt.Println("Clara dog's human age is 5")
	fmt.Println("So her dog age is ", dogAge(5))
}
