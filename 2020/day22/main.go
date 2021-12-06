package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", playCombat(parse(string(data))).Score())
	_, winningDeck := playRecursiveCombat(parse(string(data)))
	log.Println("Part two:", winningDeck.Score())
}

func parse(input string) (*queue, *queue) {
	players := strings.Split(input, "\n\n")
	return parseDeck(players[0]), parseDeck(players[1])
}

func parseDeck(input string) *queue {
	lines := strings.Split(input, "\n")
	q := newQueue(len(lines) - 1)
	for _, line := range lines {
		if strings.HasPrefix(line, "Player") {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse %s as int: %s", line, err)
		}
		q.Push(n)
	}
	return q
}

func playCombat(p1, p2 *queue) *queue {
	limit := 3000
	round := 1
	for p1.count != 0 && p2.count != 0 && limit != 0 {
		// log.Printf("-- Round %d --", round)
		a, b := p1.Pop(), p2.Pop()
		// log.Printf("Player 1 plays: %d", a)
		// log.Printf("Player 2 plays: %d", b)
		if a > b {
			// log.Println("Player 1 wins the round!")
			p1.Push(a)
			p1.Push(b)
		}
		if b > a {
			// log.Println("Player 2 wins the round!")
			p2.Push(b)
			p2.Push(a)
		}
		// log.Println("")
		limit--
		round++
	}
	if p1.count == 0 {
		return p2
	}
	return p1
}

var gameNumber = 1

func playRecursiveCombat(p1, p2 *queue) (rune, *queue) {
	cache := make(map[string]struct{})
	round := 1
	// thisGame := gameNumber
	//log.Printf("=== Game %d ===", thisGame)
	for p1.count != 0 && p2.count != 0 {
		//log.Printf("-- Round %d (Game %d) --", round, thisGame)
		d1, d2 := p1.String(), p2.String()
		//log.Printf("Player 1's deck: %s", d1)
		//log.Printf("Player 2's deck: %s", d2)
		if _, ok := cache[d1+d2]; ok {
			//log.Println("Player 1 wins the round due to infinite loop!")
			return '1', p1
		}
		cache[p1.String()+p2.String()] = struct{}{}
		a, b := p1.Pop(), p2.Pop()
		//log.Printf("Player 1 plays: %d", a)
		//log.Printf("Player 2 plays: %d", b)
		if a <= p1.count && b <= p2.count {
			//log.Println("Playing a sub-game to determine the winner...")
			gameNumber++
			winner, _ := playRecursiveCombat(p1.CopyLimit(a), p2.CopyLimit(b))
			//log.Printf("...anyway, back to game %d.", thisGame)
			switch winner {
			case '1':
				//log.Println("Player 1 wins the round!")
				p1.Push(a)
				p1.Push(b)
			case '2':
				//log.Println("Player 2 wins the round!")
				p2.Push(b)
				p2.Push(a)
			}
		} else {
			if a > b {
				//log.Println("Player 1 wins the round!")
				p1.Push(a)
				p1.Push(b)
			}
			if b > a {
				//log.Println("Player 2 wins the round!")
				p2.Push(b)
				p2.Push(a)
			}
		}
		//log.Println("")
		round++
	}
	if p1.count == 0 {
		//log.Printf("The winner of game %d is player 2!", thisGame)
		return '2', p2
	}
	//log.Printf("The winner of game %d is player 1!", thisGame)
	return '1', p1
}

func newQueue(size int) *queue {
	return &queue{
		elements: make([]int, size),
		size:     size,
	}
}

type queue struct {
	elements                []int
	size, head, tail, count int
}

func (q *queue) Copy() *queue {
	ne := make([]int, len(q.elements))
	copy(ne, q.elements)
	nq := &queue{
		elements: ne,
		size:     q.size,
		head:     q.head,
		tail:     q.tail,
		count:    q.count,
	}
	return nq
}

func (q *queue) CopyLimit(limit int) *queue {
	nq := newQueue(limit)
	cpy := q.Copy()
	for i := 0; i < limit; i++ {
		nq.Push(cpy.Pop())
	}
	return nq
}

func (q *queue) Push(i int) {
	if q.head == q.tail && q.count > 0 {
		elements := make([]int, len(q.elements)+q.size)
		copy(elements, q.elements[q.head:])
		copy(elements[len(q.elements)-q.head:], q.elements[:q.head])
		q.head = 0
		q.tail = len(q.elements)
		q.elements = elements
	}
	q.elements[q.tail] = i
	q.tail = (q.tail + 1) % len(q.elements)
	q.count++
}

func (q *queue) Pop() int {
	if q.count == 0 {
		return 0
	}
	node := q.elements[q.head]
	q.head = (q.head + 1) % len(q.elements)
	q.count--
	return node
}

func (q *queue) Score() int {
	var score int
	multiplier := q.count
	for e := q.Pop(); e != 0; e = q.Pop() {
		score += e * multiplier
		multiplier--
	}
	return score
}

func (q *queue) String() string {
	cpy := q.Copy()
	var cards []string
	for e := cpy.Pop(); e != 0; e = cpy.Pop() {
		cards = append(cards, fmt.Sprintf("%d", e))
	}
	return strings.Join(cards, ", ")
}
