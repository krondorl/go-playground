package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	split := strings.Split(str, "")
	len := len(split)
	reversed := make([]string, 0)

	for i := len - 1; i >= 0; i-- {
		reversed = append(reversed, split[i])
	}

	return strings.Join(reversed, "")
}

func main() {
	fmt.Println("Reverse string")
	fmt.Println("")
	fmt.Println("loco dice")
	fmt.Println(reverseString("loco dice"))
}
