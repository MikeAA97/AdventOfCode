package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	values := make(map[string]int)
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for idx, char := range letters {
		values[string(char)] = idx + 1
	}
	fmt.Println(values)

	file, err := open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, v := range file {
		midpoint := len(v) / 2
		firststring := v[0:midpoint]
		secondstring := v[midpoint:]

		var duplicates []string

		for _, x := range firststring {
			for _, y := range secondstring {
				if x == y {
					duplicates = append(duplicates, string(x))
				}
			}
		}
		fmt.Println(duplicates)
		sum += values[duplicates[0]]
	}
	fmt.Println(sum)

	sum2 := 0
	for i := 0; i < len(file); i += 3 {
		fmt.Println(file[i : i+3])
		test := file[i : i+3]
		fmt.Println("Zero Value:", test[0])

		var duplicates2 []string
		for _, x := range file[i] {
			for _, y := range file[i+1] {
				for _, z := range file[i+2] {
					if x == y && x == z {
						duplicates2 = append(duplicates2, string(x))
					}
				}
			}
		}
		sum2 += values[duplicates2[0]]
	}
	fmt.Println(sum2)
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
