package main

import (
	"fmt"
	"math"
)

func solveQuadraticEquation(a, b, c float64) (float64, float64) {
	result := (-1*b + math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	result2 := (-1*b - math.Sqrt(math.Pow(b, 2)-(4*a*c))) / (2 * a)
	return result, result2
}

func main() {
	fmt.Println("Quadratic formula")
	fmt.Println("")
}
