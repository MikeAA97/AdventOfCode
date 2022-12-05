package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pairs := 0
	overlaps := 0
	for _, v := range file {

		left := strings.Split(v, ",")[0]
		left_min, _ := strconv.Atoi(strings.Split(left, "-")[0])
		left_max, _ := strconv.Atoi(strings.Split(left, "-")[1])
		// var left_list []int
		// for i := left_min; i <= left_max; i++ {
		// 	left_list = append(left_list, i)
		// }
		right := strings.Split(v, ",")[1]
		right_min, _ := strconv.Atoi(strings.Split(right, "-")[0])
		right_max, _ := strconv.Atoi(strings.Split(right, "-")[1])

		// //Check if left encloses right
		// if left_min <= right_min && left_max >= right_max {
		// 	pairs += 1
		// 	fmt.Printf("%v-%v , %v-%v", left_min, left_max, right_min, right_max)
		// 	fmt.Println()
		// 	continue
		// }
		// //Check if righr encloses left
		// if left_min >= right_min && left_max <= right_max {
		// 	pairs += 1
		// 	fmt.Printf("%v-%v , %v-%v", left_min, left_max, right_min, right_max)
		// 	fmt.Println()
		// }

		left_list := numberLine(left_min, left_max)
		right_list := numberLine(right_min, right_max)

		is_overlap := false
		for _, num := range left_list {
			for _, num2 := range right_list {
				if num == num2 {
					is_overlap = true
				}
			}
		}
		if is_overlap == true {
			overlaps += 1
		}

	}
	fmt.Println(pairs)
	fmt.Println(overlaps)
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

func numberLine(min, max int) (ans []int) {
	for i := min; i <= max; i++ {
		ans = append(ans, i)
	}
	return ans
}
