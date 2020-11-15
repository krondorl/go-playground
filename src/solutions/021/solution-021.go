package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "bufio"
	// "flag"
	// "io"
	// "strings"
	// "time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	fmt.Println("Quiz Game")
	fmt.Println("")
	records := readCsvFile("problem.csv")
	fmt.Println(records)
}
