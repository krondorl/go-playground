package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func countWordsInFile(file *os.File) int {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
	count := 0

	for fileScanner.Scan() {
		count++
	}

	if err := fileScanner.Err(); err != nil {
		panic(err)
	}

	return count
}

func main() {
	fmt.Println("Wordcount")
	fmt.Println("")
	fmt.Println("count result: ", countWordsInFile(readFile("input.txt")))
}
