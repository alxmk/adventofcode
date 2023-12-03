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

func parseSchematic(input string) schematic {
	s := schematic{
		data: make(map[xy]rune),
	}
	for y, line := range strings.Split(input, "\n") {
		s.max.x, s.max.y = len(line)-1, y
		for x, r := range line {
			s.data[xy{x, y}] = r
		}
	}
	return s
}

type xy struct {
	x, y int
}

type schematic struct {
	data map[xy]rune
	max  xy
}

func (s schematic) PartOne() int {
	var sum int
	for y := 0; y <= s.max.y; y++ {
	lx:
		for x := 0; x <= s.max.x; x++ {
			r := s.data[xy{x, y}]
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				pn := []rune{r}
			li:
				for i := x + 1; i <= s.max.x; i++ {
					ri := s.data[xy{i, y}]
					switch ri {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						pn = append(pn, ri)
					default:
						break li
					}
				}
				for j := y - 1; j <= y+1 && j <= s.max.y; j++ {
					for i := x - 1; i <= x+len(pn) && i <= s.max.x; i++ {
						switch s.data[xy{i, j}] {
						case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 0:
						default:
							v, _ := strconv.Atoi(string(pn))
							sum += v
							x += len(pn) - 1
							continue lx
						}
					}
				}
				x += len(pn) - 1
			}
		}
	}
	return sum
}

func (s schematic) PartTwo() int64 {
	var sum int64
	for y := 0; y <= s.max.y; y++ {
		for x := 0; x <= s.max.x; x++ {
			r := s.data[xy{x, y}]
			switch r {
			case '*':
			default:
				continue
			}
			numbers := make(map[xy]rune)
			for j := y - 1; j <= y+1; j++ {
				for i := x - 1; i <= x+1; i++ {
					rij := s.data[xy{i, j}]
					switch rij {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						numbers[xy{i, j}] = rij
					}
				}
			}
			if len(numbers) < 2 {
				continue
			}
			var values []int64
			for l := range numbers {
				var pn []rune
				var xys []xy
			lx:
				for i := l.x - 2; i <= l.x+2; i++ {
					rij := s.data[xy{i, l.y}]
					switch rij {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						pn = append(pn, rij)
						xys = append(xys, xy{i, l.y})
					default:
						if i < l.x {
							pn = []rune{}
							xys = []xy{}
							continue
						}
						break lx
					}
				}
				for _, t := range xys {
					delete(numbers, t)
				}
				v, _ := strconv.ParseInt(string(pn), 10, 64)
				values = append(values, v)
			}
			if len(values) != 2 {
				continue
			}
			sum += values[0] * values[1]
		}
	}
	return sum
}
