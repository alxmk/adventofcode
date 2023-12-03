package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	s := parseSchematic(string(data))

	log.Println("Part one:", s.PartOne())
	log.Println("Part two:", s.PartTwo())
}

type xy struct {
	x, y int
}

func (x xy) adjacent() []xy {
	return []xy{
		{x.x - 1, x.y},
		{x.x + 1, x.y},
		{x.x, x.y - 1},
		{x.x + 1, x.y - 1},
		{x.x - 1, x.y - 1},
		{x.x, x.y + 1},
		{x.x + 1, x.y + 1},
		{x.x - 1, x.y + 1},
	}
}

type schematic struct {
	numbers []number
	nmap    map[xy]*number
	cogs    map[xy]struct{}
}

type number struct {
	value    int
	location map[xy]struct{}
}

func parseSchematic(input string) schematic {
	s := schematic{cogs: make(map[xy]struct{}), nmap: make(map[xy]*number)}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			r := line[x]
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				var pn []byte
				if n, valid := func() (*number, bool) {
					n := &number{location: make(map[xy]struct{})}
					defer func() {
						n.value, _ = strconv.Atoi(string(pn))
					}()
					var valid bool
					for i := x; i < len(line); i++ {
						ri := line[i]
						switch ri {
						case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
							c := xy{i, y}
							pn = append(pn, ri)
							n.location[c] = struct{}{}
							valid = valid || func() bool {
								for _, c := range c.adjacent() {
									switch lines[min(max(0, c.y), len(lines)-1)][min(max(0, c.x), len(line)-1)] {
									case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
										continue
									default:
										return true
									}
								}
								return false
							}()
						default:
							return n, valid
						}
					}
					return n, valid
				}(); valid {
					s.numbers = append(s.numbers, *n)
					for c := range n.location {
						s.nmap[c] = n
					}
				}
				x += len(pn) - 1
			case '*':
				s.cogs[xy{x, y}] = struct{}{}
			}
		}
	}
	return s
}

func (s schematic) PartOne() int {
	var sum int
	for _, n := range s.numbers {
		sum += n.value
	}
	return sum
}

func (s schematic) PartTwo() int {
	var sum int
	for c := range s.cogs {
		adj := make(map[*number]struct{})
		for _, a := range c.adjacent() {
			if n, ok := s.nmap[a]; ok {
				adj[n] = struct{}{}
			}
		}
		if len(adj) != 2 {
			continue
		}
		product := 1
		for n := range adj {
			product *= n.value
		}
		sum += product
	}
	return sum
}
