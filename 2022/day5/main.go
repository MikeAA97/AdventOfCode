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

	file, _ := open("input.txt")

	// fmt.Println(file[0:10])
	instr_list := file[10:]

	// Got lazy with parsing initial list
	values := map[int][]string{
		1: {"P", "V", "Z", "W", "D", "T"},
		2: {"D", "J", "F", "V", "W", "S", "L"},
		3: {"H", "B", "T", "V", "S", "L", "M", "Z"},
		4: {"J", "S", "R"},
		5: {"W", "L", "M", "F", "G", "B", "Z", "C"},
		6: {"B", "G", "R", "Z", "H", "V", "W", "Q"},
		7: {"N", "D", "B", "C", "P", "J", "V"},
		8: {"Q", "B", "T", "P"},
		9: {"C", "R", "Z", "G", "H"},
	}
	values2 := map[int][]string{
		1: {"P", "V", "Z", "W", "D", "T"},
		2: {"D", "J", "F", "V", "W", "S", "L"},
		3: {"H", "B", "T", "V", "S", "L", "M", "Z"},
		4: {"J", "S", "R"},
		5: {"W", "L", "M", "F", "G", "B", "Z", "C"},
		6: {"B", "G", "R", "Z", "H", "V", "W", "Q"},
		7: {"N", "D", "B", "C", "P", "J", "V"},
		8: {"Q", "B", "T", "P"},
		9: {"C", "R", "Z", "G", "H"},
	}

	for _, v := range instr_list {
		move(v, values)
	}

	for _, v := range instr_list {
		move2(v, values2)
	}

	fmt.Print("Part 1: ")
	for i := 0; i < len(values); i++ {
		fmt.Print(values[i+1][0])
	}
	fmt.Println()
	fmt.Print("Part 2: ")
	for i := 0; i < len(values); i++ {
		fmt.Print(values2[i+1][0])
	}

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

func removeElement(input []string, val string) []string {
	j := 0
	for _, v := range input {
		if v != val {
			input[j] = v
			j++
		}
	}
	return input[:j]
}

func move(instructions string, grid map[int][]string) {

	//Parse movement statement to target correct stacks
	breakdown := strings.Fields(instructions)
	from, _ := strconv.Atoi(breakdown[3])
	to, _ := strconv.Atoi(breakdown[5])
	max, _ := strconv.Atoi(breakdown[1])

	//Ensure we can move at the most the number of elements in the list
	if len(grid[from]) < max {
		max = len(grid[from])
	}

	//Iterate through each set of instructions, pop value from "from" stack into "to" stack
	for i := 0; i <= max-1; i++ {
		popped := []string{grid[from][0]}
		grid[from] = grid[from][1:]
		grid[to] = append(popped, grid[to]...)
	}
}

func move2(instructions string, grid map[int][]string) {

	//Parse movement statement to target correct stacks
	breakdown := strings.Fields(instructions)
	from, _ := strconv.Atoi(breakdown[3])
	to, _ := strconv.Atoi(breakdown[5])
	max, _ := strconv.Atoi(breakdown[1])

	//Ensure we can move at the most the number of elements in the list
	if len(grid[from]) < max {
		max = len(grid[from])
	}

	//New Variable to hold the values to be moved, had issues overwriting the stack in question
	move_section := make([]string, max)
	move_section = append(move_section, grid[from][0:max]...)

	//Removing whitespace from appended move_section
	move_section = removeElement(move_section, "")

	//Removing chunk from the "from" stack
	grid[from] = grid[from][max:]

	//Prepending "to" stack with new values
	grid[to] = append(move_section, grid[to]...)
}
