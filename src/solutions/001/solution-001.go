package main

import "fmt"

func main() {
	fmt.Println("Fizzbuzz app")
	fmt.Println("")
    for n := 1; n <= 100; n++ {
		if n % 15 == 0 {
			fmt.Println("fizzbuzz")
		} else {
			if n % 3 == 0 {
				fmt.Println("fizz")
			} else if n % 5 == 0 {
				fmt.Println("buzz")
			} else {
				fmt.Println(n)
			}	
		}
    }
}