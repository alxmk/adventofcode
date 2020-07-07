package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.ibm.com/alexmk/adventofcode/2019/day2/intcode"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	programme, err := intcode.Parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input as intcode programme:", err)
	}

	log.Println("Part one:", len(newBot(programme).Run(0).contents))

	log.Printf("Part two:\n%s", newBot(programme).Run(1))
}

type bot struct {
	in, out   chan int64
	pos       coordinate
	direction coordinate
	brain     intcode.Programme
}

type coordinate struct {
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func (c coordinate) turnLeft() coordinate {
	switch c {
	case up:
		return left
	case left:
		return down
	case down:
		return right
	case right:
		return up
	default:
		panic("invalid coordinate to turn left")
	}
}

func (c coordinate) turnRight() coordinate {
	switch c {
	case up:
		return right
	case left:
		return up
	case down:
		return left
	case right:
		return down
	default:
		panic("invalid coordinate to turn right")
	}
}

var (
	up    coordinate = coordinate{0, 1}
	down  coordinate = coordinate{0, -1}
	left  coordinate = coordinate{-1, 0}
	right coordinate = coordinate{1, 0}
)

type surface struct {
	contents               map[coordinate]int64
	xmax, xmin, ymax, ymin int
}

func (s surface) String() string {
	var out strings.Builder
	for y := s.ymax; y >= s.ymin; y-- {
		for x := s.xmin; x <= s.xmax; x++ {
			if o, ok := s.contents[coordinate{x, y}]; !ok || o == 0 {
				out.WriteRune('.')
			} else {
				out.WriteRune('#')
			}
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func newBot(brain intcode.Programme) *bot {
	return &bot{
		in:        make(chan int64, 1),
		out:       make(chan int64),
		pos:       coordinate{0, 0},
		direction: up,
		brain:     brain.Copy(),
	}
}

func (b *bot) Run(startingColour int64) *surface {
	go func() {
		if err := b.brain.Run(b.in, b.out); err != nil {
			panic("intcode computer failed:" + err.Error())
		}
	}()

	s := &surface{
		contents: make(map[coordinate]int64),
	}
	s.contents[b.pos] = startingColour
	b.in <- s.contents[b.pos]

	var isDirection bool
	for o := range b.out {
		if !isDirection {
			s.contents[b.pos] = o
			isDirection = !isDirection
			continue
		}
		switch o {
		case 0:
			b.direction = b.direction.turnLeft()
		case 1:
			b.direction = b.direction.turnRight()
		default:
			panic(fmt.Sprintf("unexpected bot response code %d", o))
		}
		b.move(s)
		b.in <- s.contents[b.pos]
		isDirection = !isDirection
	}

	close(b.in)

	return s
}

func (b *bot) move(s *surface) {
	b.pos.x += b.direction.x
	b.pos.y += b.direction.y

	if b.pos.x < s.xmin {
		s.xmin = b.pos.x
	}
	if b.pos.x > s.xmax {
		s.xmax = b.pos.x
	}
	if b.pos.y < s.ymin {
		s.ymin = b.pos.y
	}
	if b.pos.y > s.ymax {
		s.ymax = b.pos.y
	}
}
