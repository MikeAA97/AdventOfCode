package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, _ := open("input.txt")

	text := file[0]

	fmt.Println("Part 1:", checkUnique(text, 4))
	fmt.Println("Part 2:", checkUnique(text, 14))

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

func checkUnique(input string, num int) int {

	//Make slice of X capacity
	prev_x := make([]string, 0, num)

	//Iterate through string
	for idx, value := range input {

		//Fill up to X amount
		if len(prev_x) != num {
			prev_x = append(prev_x, string(value))
		}

		//After X amount check for uniqueness, or add next value to queue
		if len(prev_x) == num {

			prev_x = prev_x[1:]
			prev_x = append(prev_x, string(value))

			values := make(map[string]int)
			for _, char := range prev_x {
				values[string(char)] += 1
			}
			if len(values) == num {
				return idx + 1
			}
		}
	}
	return -1
}
