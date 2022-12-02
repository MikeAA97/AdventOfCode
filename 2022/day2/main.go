package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Function to open and parse file into []string
	file, err := open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	score := 0

	//Part 1
	for _, v := range file {
		text := strings.Fields(v)

		fmt.Println(text)
		switch text[0] {
		case "A": //Rock
			switch text[1] {
			case "X":
				score += 4 //1 for rock, 3 for draw
			case "Y":
				score += 8 //2 for paper, 6 for win
			case "Z":
				score += 3 //3 for scissors, 0 for loss
			}
		case "B": //Paper
			switch text[1] {
			case "X":
				score += 1 //1 for rock, 0 for loss
			case "Y":
				score += 5 //2 for paper, 3 for draw
			case "Z":
				score += 9 //3 for scissors, 6 for win
			}
		case "C": //Scissors
			switch text[1] {
			case "X":
				score += 7 //1 for rock, 6 for win
			case "Y":
				score += 2 //2 for paper, 0 for loss
			case "Z":
				score += 6 //3 for scissors, 3 for draw
			}
		}
	}

	//Part 2
	score2 := 0
	for _, v := range file {
		text := strings.Fields(v)

		fmt.Println(text)
		switch text[0] {
		case "A": //Rock
			switch text[1] {
			case "X":
				score2 += 3 //0 for loss, 3 for scissors
			case "Y":
				score2 += 4 //3 for draw, 1 for rock
			case "Z":
				score2 += 8 //6 for win, 2 for paper
			}
		case "B": //Paper
			switch text[1] {
			case "X":
				score2 += 1 //0 for loss, 1 for rock
			case "Y":
				score2 += 5 //3 for draw, 2 for paper
			case "Z":
				score2 += 9 //6 for win, 3 for scissors
			}
		case "C": //Scissors
			switch text[1] {
			case "X":
				score2 += 2 //0 for loss, 2 for paper
			case "Y":
				score2 += 6 //3 for draw, 3 for scissors
			case "Z":
				score2 += 7 //6 for win, 1 for rock
			}
		}
	}
	fmt.Println("Part 1 Total Score:", score)
	fmt.Println("Part 2 Total Score:", score2)

	// Loop through adding in running total and resetting value at the prescense of newline
	// for _, v := range file {
	// 	if v != "" {
	// 		integer_value, err := strconv.Atoi(v)
	// 		if err != nil {
	// 			panic(err)
	// 		}

	// 		total += integer_value
	// 	}
	// 	if v == "" {
	// 		fmt.Println("Total:", total)
	// 		results = append(results, total)

	// 		total = 0
	// 	}
	// }
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
