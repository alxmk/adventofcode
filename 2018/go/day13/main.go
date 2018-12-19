package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("amf.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	lines := strings.Split(string(data), "\n")
	xmax, ymax := len(lines[0]), len(lines)

	track := make(circuit, xmax)

	var cars []*car

	for y, line := range lines {
		for x, r := range line {
			if y == 0 {
				track[x] = make([]rail, ymax)
			}
			switch r {
			case '-':
				track[x][y] = straightlat
			case '|':
				track[x][y] = straightlong
			case '\\':
				track[x][y] = lcurve
			case '/':
				track[x][y] = rcurve
			case '+':
				track[x][y] = intersection
			case '^':
				track[x][y] = straightlong
				cars = append(cars, &car{x: x, y: y, direction: north, nextTurn: left})
			case 'v':
				track[x][y] = straightlong
				cars = append(cars, &car{x: x, y: y, direction: south, nextTurn: left})
			case '>':
				track[x][y] = straightlat
				cars = append(cars, &car{x: x, y: y, direction: east, nextTurn: left})
			case '<':
				track[x][y] = straightlat
				cars = append(cars, &car{x: x, y: y, direction: west, nextTurn: left})
			}
		}
	}

	sort.Sort(byXY(cars))

	// track.Print(cars)

	// var t int

	for {
		// track.Print(cars)
		var toremove []int
		for i, car := range cars {
			car.Tick(track)
			for j, c := range cars {
				if j == i {
					continue
				}
				if car.x == c.x && car.y == c.y {
					log.Println("Collision at", c.x, c.y)
					toremove = append(toremove, i, j)
				}
			}
		}
		// for i, car := range cars {
		// 	fmt.Print(t, i, car.x, car.y, "\n")
		// }
		sort.Sort(sort.Reverse(sort.IntSlice(toremove)))
		for _, i := range toremove {
			cars = append(cars[:i], cars[i+1:]...)
		}
		if len(cars) == 1 {
			// fmt.Println(t)
			fmt.Println("Last car", cars[0].x, cars[0].y)
			break
		}
		sort.Sort(byXY(cars))
		// t++
	}
}

type circuit [][]rail

func (c circuit) Print(cars []*car) {
	for y := range c[0] {
		for x := range c {
			var carHere bool
			for _, car := range cars {
				if car.x == x && car.y == y {
					fmt.Print(car.String())
					carHere = true
				}
			}
			if !carHere {
				fmt.Print(c[x][y].String())
			}
		}
		fmt.Print("\n")
	}
}

type rail int

const (
	none rail = iota
	straightlong
	straightlat
	intersection
	rcurve
	lcurve
)

func (r rail) String() string {
	switch r {
	case straightlong:
		return "|"
	case straightlat:
		return "-"
	case lcurve:
		return "\\"
	case rcurve:
		return "/"
	case intersection:
		return "+"
	}
	return " "
}

type car struct {
	x         int
	y         int
	direction dir
	nextTurn  turn
}

type byXY []*car

func (a byXY) Len() int      { return len(a) }
func (a byXY) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byXY) Less(i, j int) bool {
	if a[i].y == a[j].y {
		return a[i].x < a[j].x
	}
	return a[i].y < a[j].y
}

func (c *car) String() string {
	switch c.direction {
	case north:
		return "^"
	case south:
		return "v"
	case east:
		return ">"
	case west:
		return "<"
	}
	return " "
}

func (c *car) Tick(track circuit) {
	switch c.direction {
	case north:
		c.y--
	case south:
		c.y++
	case east:
		c.x++
	case west:
		c.x--
	}

	switch track[c.x][c.y] {
	case straightlong, straightlat:
		// Nothing to do, we're going straight
	case intersection:
		c.direction = c.direction.Resultant(c.nextTurn)
		c.nextTurn = c.nextTurn.Next()
	case rcurve:
		switch c.direction {
		case north, south:
			c.direction = c.direction.Resultant(right)
		case east, west:
			c.direction = c.direction.Resultant(left)
		}
	case lcurve:
		switch c.direction {
		case north, south:
			c.direction = c.direction.Resultant(left)
		case east, west:
			c.direction = c.direction.Resultant(right)
		}
	}
}

type dir int

const (
	north dir = iota
	east
	south
	west
)

func (d dir) Resultant(t turn) dir {
	switch t {
	case cont:
		return d
	case left:
		switch d {
		case north:
			return west
		default:
			return d - 1
		}
	case right:
		switch d {
		case west:
			return north
		default:
			return d + 1
		}
	}
	return d
}

type turn int

const (
	left turn = iota
	cont
	right
)

func (t turn) Next() turn {
	switch t {
	case left:
		return cont
	case cont:
		return right
	case right:
		return left
	}
	return -1
}
