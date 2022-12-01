package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	// Function to open and parse file into []string
	file, err := open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	// Slice to capture each sum chunk
	var results []int

	// Loop through adding in running total and resetting value at the prescense of newline
	for _, v := range file {
		if v != "" {
			integer_value, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			total += integer_value
		}
		if v == "" {
			fmt.Println("Total:", total)
			results = append(results, total)

			total = 0
		}
	}

	//Sort the summed values into reverse order
	sort.Sort(sort.Reverse(sort.IntSlice(results)))

	//Get the "max"
	fmt.Println(results[0])

	//Get the top 3 values then sum them
	top3 := results[0:3]
	top3_total := 0

	for _, v := range top3 {
		top3_total += v
	}

	fmt.Println(top3_total)
}

func open(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, err
}
