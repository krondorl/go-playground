package main

import (
	"fmt"
	"strings"
)

func abbreviate(expression string) string {
	var abbreviation strings.Builder
	split := strings.Split(expression, " ")

	for _, value := range split {
		abbreviation.WriteString(strings.ToUpper(string(value[0])))
	}

	return abbreviation.String()
}

func main() {
	fmt.Println("Abbreviations")
	fmt.Println("")
	fmt.Println("abbreviate(zero latency monitoring) -> ", abbreviate("zero latency monitoring"))
	fmt.Println("abbreviate(chicken feeder device) -> ", abbreviate("chicken feeder device"))
	fmt.Println("abbreviate(acoustic transducer company) -> ", abbreviate("acoustic transducer company"))
}
