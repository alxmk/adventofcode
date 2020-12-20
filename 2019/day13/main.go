package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/alxmk/adventofcode/2019/day2/intcode"
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

	out := make(chan int64)
	go func() {
		if err := programme.Copy().Run(nil, out); err != nil {
			log.Fatalln(err)
		}
	}()

	game := newScreen()
	var idx int
	var x, y int64
	for o := range out {
		switch idx % 3 {
		case 0:
			x = o
			if x > game.xsize {
				game.xsize = x
			}
		case 1:
			y = o
			if y > game.ysize {
				game.ysize = y
			}
		case 2:
			game.blocks[coordinate{x, y}] = o
		}
		idx++
	}

	var numBlocks int
	for _, tile := range game.blocks {
		if tile == 2 {
			numBlocks++
		}
	}

	log.Println("Part one:", numBlocks)

	// Free play woop woop
	programme[0] = 2

	log.Println("Part two:", play(programme))
}

type coordinate struct {
	x, y int64
}

type screen struct {
	blocks       map[coordinate]int64
	xsize, ysize int64
}

func newScreen() *screen {
	return &screen{blocks: make(map[coordinate]int64)}
}

func (s screen) String() string {
	var out strings.Builder
	for y := int64(0); y <= s.ysize; y++ {
		for x := int64(0); x <= s.xsize; x++ {
			switch s.blocks[coordinate{x, y}] {
			case 0:
				out.WriteRune(' ')
			case 1:
				out.WriteRune('#')
			case 2:
				out.WriteRune('X')
			case 3:
				out.WriteRune('=')
			case 4:
				out.WriteRune('O')
			default:
				panic(fmt.Sprintf("unrecognised block type"))
			}
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func play(programme intcode.Programme) int64 {
	// Set up the processor
	in, out := make(chan int64), make(chan int64)
	go func() {
		if err := programme.Copy().Run(in, out); err != nil {
			log.Fatalln(err)
		}
	}()

	// Set up a channel to send the moves down
	moves := make(chan int64)
	go func() {
		for m := range moves {
			in <- m
		}
	}()

	// Initialise the game state
	game := newScreen()

	var idx, ballX, paddleX, score int64
	var seenPaddle, seenBall bool
	var position coordinate
	for o := range out {
		switch idx % 3 {
		case 0:
			if position.x = o; position.x > game.xsize {
				game.xsize = position.x
			}
		case 1:

			if position.y = o; position.y > game.ysize {
				game.ysize = position.y
			}
		case 2:
			switch position {
			case coordinate{-1, 0}:
				score = o
			default:
				switch o {
				case 3:
					// If it's the paddle position update it
					paddleX = position.x
					seenPaddle = true
				case 4:
					// If it's the ball position update it
					ballX = position.x
					seenBall = true
				}
				game.blocks[position] = o
				// If we know where the paddle and ball are we can plot our next move
				if seenPaddle && seenBall {
					nextMove := getNextMove(ballX, paddleX)
					moves <- nextMove
					seenBall, seenPaddle = false, false // Reset for the next tick
					if nextMove == 0 {
						seenPaddle = true // If we don't move the paddle we won't get a position update, and we know where it is already
					}
				}
			}
		}
		idx++
	}

	return score
}

func getNextMove(ball, paddle int64) int64 {
	if ball > paddle {
		// Move right
		return 1
	}
	if ball < paddle {
		// Move left
		return -1
	}
	// Stay still
	return 0
}
