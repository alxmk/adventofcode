package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	partOne()

	value := "077201"

	start := &score{value: 3}
	end := &score{value: 7, next: start}
	start.next = end

	elves := elves{&elf{current: start}, &elf{current: end}}

	var lastsix []rune
	i := 2
	for {
		sum := elves.Sum()

		// start.Print(end, elves)

		for _, r := range fmt.Sprintf("%d", sum) {
			val, _ := strconv.Atoi(string(r))
			end.next = &score{value: val, next: start}
			end = end.next
			lastsix = append(lastsix, r)
			if len(lastsix) > 6 {
				lastsix = lastsix[1:]
			}

			i++

			if string(lastsix) == value {
				log.Println(i - len(value))
				return
			}
		}

		// start.Print(end, elves)

		for _, e := range elves {
			// log.Printf("Elf %d moving %d", i, e.current.value+1)
			moves := e.current.value + 1
			for j := 0; j < moves; j++ {
				// log.Println(j)
				e.current = e.current.next
				// start.Print(end, elves)
			}
		}
	}
}

func partOne() {
	n := 77201

	start := &score{value: 3}
	end := &score{value: 7, next: start}
	start.next = end

	elves := elves{&elf{current: start}, &elf{current: end}}

	var lastten []rune
	i := 2
	for i < n+10 {
		sum := elves.Sum()

		// start.Print(end, elves)

		for _, r := range fmt.Sprintf("%d", sum) {
			val, _ := strconv.Atoi(string(r))
			end.next = &score{value: val, next: start}
			end = end.next
			lastten = append(lastten, r)
			if len(lastten) > 10 {
				lastten = lastten[1:]
			}
			i++
		}

		// start.Print(end, elves)

		for _, e := range elves {
			// log.Printf("Elf %d moving %d", i, e.current.value+1)
			moves := e.current.value + 1
			for j := 0; j < moves; j++ {
				// log.Println(j)
				e.current = e.current.next
				// start.Print(end, elves)
			}
		}
	}

	// start.Print(end, elves)

	log.Println(string(lastten))
}

type elves []*elf

func (e elves) Sum() int {
	var sum int

	for _, elf := range e {
		sum += elf.current.value
	}

	return sum
}

type score struct {
	value int
	next  *score
}

type elf struct {
	current *score
}

func (s *score) Print(last *score, e elves) {
	var here bool
	for _, elf := range e {
		if elf.current == s {
			here = true
			break
		}
	}
	if here {
		fmt.Printf("(%d) ", s.value)
	} else {
		fmt.Printf(" %d  ", s.value)
	}
	if s != last {
		s.next.Print(last, e)
	} else {
		fmt.Print("\n")
	}
}
