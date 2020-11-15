package main

import (
	"fmt"
)

func getTriangleType(a int, b int, c int) string {
	triangleType := ""

	if a == b || b == c || a == c {
		triangleType = "isoscele"
	}

	if a == b && b == c {
		triangleType = "equilateral"
	}

	if a != b && b != c && a != c {
		triangleType = "scalene"
	}

	return triangleType
}

func main() {
	fmt.Println("Triangle")
	fmt.Println("")
	fmt.Println("(1,1,1) ", getTriangleType(1, 1, 1))
	fmt.Println("(1,1,2) ", getTriangleType(1, 1, 2))
	fmt.Println("(1,2,3) ", getTriangleType(1, 2, 3))
}
