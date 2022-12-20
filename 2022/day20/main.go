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

	log.Println("Part one:", partOne(parse(string(data))))
	log.Println("Part two:", partTwo(parse(string(data))))
}

func partOne(first *element, all []*element) int64 {
	return score(mix(first, all))
}

func partTwo(first *element, all []*element) int64 {
	for _, a := range all {
		a.value *= 811589153
	}
	var current *element
	for i := 0; i < 10; i++ {
		current = mix(first, all)
	}
	return score(current)
}

func score(e *element) int64 {
	var sum int64
	for i := 0; i < 3000; i++ {
		e = e.next
		switch i {
		case 999, 1999, 2999:
			sum += e.value
		}
	}
	return sum
}

func mix(first *element, all []*element) *element {
	var zero *element
	for _, a := range all {
		if a.value == 0 {
			zero = a
			continue
		}
		if a == first {
			first = a.next
		}
		a.previous.next = a.next
		a.next.previous = a.previous
		insert := a
		if a.value > 0 {
			for i := int64(0); i < a.value%int64(len(all)-1); i++ {
				insert = insert.next
			}
		} else {
			for i := int64(0); i >= a.value%int64(len(all)-1); i-- {
				insert = insert.previous
			}
		}
		a.next = insert.next
		a.next.previous, insert.next, a.previous = a, a, insert
	}
	return zero
}

type element struct {
	value          int64
	previous, next *element
}

func parse(input string) (*element, []*element) {
	var first, current *element
	var all []*element
	for i, line := range strings.Split(input, "\n") {
		v, _ := strconv.ParseInt(line, 10, 64)
		this := &element{value: v}
		if i == 0 {
			first = this
		}
		if current != nil {
			this.previous, current.next = current, this
		}
		current = this
		all = append(all, this)
	}
	last := all[len(all)-1]
	first.previous, last.next = last, first
	return first, all
}

func (e element) String() string {
	first := e.value
	var b strings.Builder
	for next, i := &e, 0; next.value != first || i == 0; next, i = next.next, i+1 {
		b.WriteString(strconv.FormatInt(next.value, 10))
		b.WriteString(",")
	}
	return b.String()
}
