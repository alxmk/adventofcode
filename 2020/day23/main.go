package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	cups := parse(string(data))
	for i := 0; i < 100; i++ {
		cups.Move()
	}
	log.Println("Part one:", cups.FromOne())
	cups = parsePartTwo(string(data))
	for i := 0; i < 10000000; i++ {
		cups.Move()
	}
	log.Println("Part two:", cups.FindStars())
}

type cup struct {
	number int
	next   *cup
}

type cups struct {
	list    map[int]*cup
	pointer int
	max     int
}

func parse(input string) *cups {
	cups := cups{
		list: make(map[int]*cup),
	}
	var prev, last *cup
	for i := len(input) - 1; i >= 0; i-- {
		n := int(input[i] - '0')
		cups.list[n] = &cup{number: n, next: prev}
		if i == len(input)-1 {
			last = cups.list[n]
		}
		prev = cups.list[n]
		if n > cups.max {
			cups.max = n
		}
	}
	last.next = prev
	cups.pointer = prev.number
	return &cups
}

func parsePartTwo(input string) *cups {
	cups := cups{
		list: make(map[int]*cup),
	}
	var prev, last *cup
	for i := 1000000 - 1; i >= 0; i-- {
		var n int
		if i >= len(input) {
			n = i + 1
		} else {
			n = int(input[i] - '0')
		}
		cups.list[n] = &cup{number: n, next: prev}
		if i == 1000000-1 {
			last = cups.list[n]
		}
		prev = cups.list[n]
		if n > cups.max {
			cups.max = n
		}
	}
	last.next = prev
	cups.pointer = prev.number
	return &cups
}

func (c cups) String() string {
	var b strings.Builder
	b.WriteRune(rune(c.pointer) + '0')
	for cup := c.list[c.pointer].next; ; cup = cup.next {
		if cup.number == c.pointer {
			break
		}
		b.WriteRune(rune(cup.number) + '0')
	}
	return b.String()
}

func (c cups) StringLimit(limit int) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%d", c.pointer))
	for cup, i := c.list[c.pointer].next, 0; i < limit; cup, i = cup.next, i+1 {
		if cup.number == c.pointer {
			break
		}
		b.WriteString(fmt.Sprintf("%d", cup.number))
	}
	return b.String()
}

func (c cups) FromOne() string {
	var b strings.Builder
	pointer := 1
	for cup := c.list[pointer].next; ; cup = cup.next {
		if cup.number == pointer {
			break
		}
		b.WriteRune(rune(cup.number) + '0')
	}
	return b.String()
}

func (c cups) FromOneLimit(limit int) string {
	var b strings.Builder
	pointer := 1
	for cup, i := c.list[pointer].next, 0; i < limit; cup, i = cup.next, i+1 {
		if cup.number == pointer {
			break
		}
		b.WriteString(fmt.Sprintf("%d\n", cup.number))
	}
	return b.String()
}

func (c cups) FindStars() int64 {
	return int64(c.list[1].next.number) * int64(c.list[1].next.next.number)
}

func (c *cups) Move() {
	head, middle, tail := c.list[c.pointer].next, c.list[c.pointer].next.next, c.list[c.pointer].next.next.next
	destination := c.pointer
	for {
		destination--
		if destination == 0 {
			destination = c.max
		}
		if destination == head.number || destination == middle.number || destination == tail.number {
			continue
		}
		break
	}
	insertion := c.list[destination]
	c.list[c.pointer].next = tail.next
	tail.next = insertion.next
	insertion.next = head

	c.pointer = c.list[c.pointer].next.number
}
