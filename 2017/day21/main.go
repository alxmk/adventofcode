package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	ruleLines := strings.Split(string(input), "\n")

	rules := []Rule{}

	for _, r := range ruleLines {
		rules = append(rules, parseRule(r))
	}

	grid := map[int]map[int]bool{
		0: map[int]bool{
			0: false,
			1: true,
			2: false,
		},
		1: map[int]bool{
			0: false,
			1: false,
			2: true,
		},
		2: map[int]bool{
			0: true,
			1: true,
			2: true,
		},
	}

	printGrid(grid)

	for i := 0; i < 5; i++ {
		gridLets := splitGrid(grid)

		newGridLets := [][]bool{}

		for _, g := range gridLets {
			output, err := matchRules(g, rules)
			if err != nil {
				log.Fatalln("Couldn't match rule")
			}
			newGridLets = append(newGridLets, output)
		}

		grid = assembleGrid(newGridLets)

		printGrid(grid)
	}

	// Count the trues
	trueCount := 0
	for _, row := range grid {
		for _, elem := range row {
			if elem {
				trueCount++
			}
		}
	}

	log.Println("Part one answer is", trueCount)
}

func matchRules(input []bool, rules []Rule) ([]bool, error) {
	for _, r := range rules {
		if r.Match(input) {
			return r.output, nil
		}
	}

	return nil, fmt.Errorf("Didn't match any rules")
}

func printGrid(grid map[int]map[int]bool) {
	fmt.Println("=============================")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			output := "."
			if grid[i][j] {
				output = "#"
			}
			fmt.Print(output)
		}
		fmt.Print("\n")
	}
	fmt.Println("=============================")
}

func assembleGrid(gridLets [][]bool) map[int]map[int]bool {
	gridletSideLength := int(math.Sqrt(float64(len(gridLets[0]))))
	gridLetsPerSide := int(math.Sqrt(float64(len(gridLets))))
	sideLength := gridLetsPerSide * gridletSideLength

	toReturn := make(map[int]map[int]bool)

	for i := 0; i < len(gridLets); i++ {
		startX := i % (sideLength / gridletSideLength) * gridletSideLength
		// This looks weird but it works because it's integer division
		startY := i / gridLetsPerSide * gridletSideLength

		elem := 0

		for y := 0; y < gridletSideLength; y++ {
			for x := 0; x < gridletSideLength; x++ {
				if _, ok := toReturn[startY+y]; !ok {
					toReturn[startY+y] = make(map[int]bool)
				}
				toReturn[startY+y][startX+x] = gridLets[i][elem]
				elem++
			}
		}
	}

	// Check squareness
	for _, row := range toReturn {
		if len(row) != len(toReturn) {
			log.Println(gridLets)
			log.Fatalln("Row length is not the same as grid height, grid is not square", len(row), len(toReturn), toReturn)
		}
	}

	return toReturn
}

func splitGrid(grid map[int]map[int]bool) [][]bool {
	sideLength := len(grid)
	elementSize := 3
	if sideLength%2 == 0 {
		elementSize = 2
	}
	numGrids := sideLength / elementSize

	toReturn := [][]bool{}

	for i := 0; i < int(math.Pow(float64(numGrids), 2)); i++ {
		startX := i % (sideLength / elementSize) * elementSize
		// This looks weird but it works because it's integer division
		startY := i / numGrids * elementSize

		fmt.Println("Start:", i, startX, startY)

		thisSubGrid := []bool{}

		for y := 0; y < elementSize; y++ {
			for x := 0; x < elementSize; x++ {
				thisSubGrid = append(thisSubGrid, grid[startY+y][startX+x])
			}
		}

		toReturn = append(toReturn, thisSubGrid)
	}

	return toReturn
}

func printGridlet(gridlet []bool) {
	sidelength := int(math.Sqrt(float64(len(gridlet))))

	for i := 0; i < len(gridlet); i++ {
		out := "."
		if gridlet[i] {
			out = "#"
		}
		fmt.Print(out)

		if (i+1)%sidelength == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")
}

type Rule struct {
	size   int
	match  []bool
	output []bool
}

func (r *Rule) Match(input []bool) bool {
	toTry := make([]bool, len(input))
	copy(toTry, input)

	// Three times rotate and try again
	for i := 0; i < 4; i++ {
		if match(toTry, r.match) || match(flipped(toTry), r.match) {
			r.Print()
			return true
		}

		toTry = rotate(toTry)
	}

	// If we got this far then we'll never match
	return false
}

func (r *Rule) Print() {
	fmt.Print("Matched rule: ")
	for i, elem := range r.match {
		if i != 0 && i%r.size == 0 {
			fmt.Print("/")
		}
		output := "."
		if elem {
			output = "#"
		}
		fmt.Print(output)
	}

	fmt.Print(" => ")

	for i, elem := range r.output {
		if i != 0 && i%(r.size+1) == 0 {
			fmt.Print("/")
		}
		output := "."
		if elem {
			output = "#"
		}
		fmt.Print(output)
	}
	fmt.Print("\n")
}

func flipped(input []bool) []bool {
	output := []bool{}

	if len(input)%2 == 0 {
		output = append(output, input[2])
		output = append(output, input[3])
		output = append(output, input[0])
		output = append(output, input[1])
		return output
	}

	output = append(output, input[6])
	output = append(output, input[7])
	output = append(output, input[8])
	output = append(output, input[3])
	output = append(output, input[4])
	output = append(output, input[5])
	output = append(output, input[0])
	output = append(output, input[1])
	output = append(output, input[2])

	return output
}

func rotate(input []bool) []bool {
	output := []bool{}

	if len(input)%2 == 0 {
		output = append(output, input[2])
		output = append(output, input[0])
		output = append(output, input[3])
		output = append(output, input[1])
		return output
	}

	output = append(output, input[6])
	output = append(output, input[3])
	output = append(output, input[0])
	output = append(output, input[7])
	output = append(output, input[4])
	output = append(output, input[1])
	output = append(output, input[8])
	output = append(output, input[5])
	output = append(output, input[2])

	return output
}

func match(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i, val := range a {
		if b[i] != val {
			return false
		}
	}

	return true
}

func parseRule(input string) Rule {
	// Split into match and output
	parts := strings.Split(input, " => ")

	// Create the rule assuming it's of size 3
	rule := Rule{
		size: 3,
	}

	// If it isn't, correct that
	if len(parts[0]) == 5 {
		rule.size = 2
	}

	// Parse match
	for _, part := range strings.Split(parts[0], "/") {
		for _, c := range []rune(part) {
			rule.match = append(rule.match, c == '#')
		}
	}

	// Parse output
	for _, part := range strings.Split(parts[1], "/") {
		for _, c := range []rune(part) {
			rule.output = append(rule.output, c == '#')
		}
	}

	return rule
}
