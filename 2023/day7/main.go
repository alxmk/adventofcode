package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	h := parseHands(string(data))

	log.Println("Part one:", solve(h, false))
	log.Println("Part two:", solve(h, true))
}

func solve(hands []hand, joker bool) int {
	sort.Slice(hands, func(i, j int) bool {
		return hands[j].Beats(hands[i], joker)
	})
	var sum int
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	return sum
}

func parseHands(input string) []hand {
	var hands []hand
	for _, line := range strings.Split(input, "\n") {
		var h hand
		fmt.Sscanf(line, "%s %d", &h.cards, &h.bid)
		h.jackPower, h.jokerPower = calculatePowers(h.cards)
		hands = append(hands, h)
	}
	return hands
}

type hand struct {
	cards      string
	jackPower  int
	jokerPower int
	bid        int
}

func (h hand) Power(jokers bool) int {
	if jokers {
		return h.jokerPower
	}
	return h.jackPower
}

func calculatePowers(hand string) (int, int) {
	hmap := make(map[rune]int)
	for _, r := range hand {
		hmap[r]++
	}
	return handPower(hmap), handPower(convertJokers(hmap))
}

func convertJokers(hmap map[rune]int) map[rune]int {
	nj, ok := hmap['J']
	if !ok {
		return hmap
	}
	delete(hmap, 'J')
	var pairs, singles []rune
	for c, nc := range hmap {
		switch nc {
		case 4, 3:
			// This is always the best option to bump
			hmap[c] += nj
			return hmap
		case 2:
			pairs = append(pairs, c)
		case 1:
			singles = append(singles, c)
		}
	}
	if len(pairs) != 0 {
		// Find the highest pair and bump it
		sort.Slice(pairs, func(i, j int) bool {
			return power(pairs[i], false) < power(pairs[j], false)
		})
		hmap[pairs[0]] += nj
		return hmap
	}
	if len(singles) != 0 {
		// Find the highest single and bump it
		sort.Slice(singles, func(i, j int) bool {
			return power(singles[i], false) < power(singles[j], false)
		})
		hmap[singles[0]] += nj
		return hmap
	}
	// They were all Jokers
	hmap['A'] += nj
	return hmap
}

func (h hand) Beats(j hand, jokers bool) bool {
	ph, pj := h.Power(jokers), j.Power(jokers)
	if ph == pj {
		return tiebreak(h.cards, j.cards, jokers)
	}
	return ph > pj
}

var cardPower = map[rune]int{
	'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12,
}

func power(c rune, jokers bool) int {
	if jokers && c == 'J' {
		return -1
	}
	return cardPower[c]
}

func tiebreak(a, b string, jokers bool) bool {
	for i, ch := range a {
		ph, pj := power(ch, jokers), power(rune(b[i]), jokers)
		if ph == pj {
			continue
		}
		return ph > pj
	}
	panic(fmt.Sprintf("%s vs %s is a draw", a, b))
}

func handPower(h map[rune]int) int {
	switch len(h) {
	case 1: // Five of a kind
		return 6
	case 2: // Four of a kind or full house
		for _, l := range h {
			switch l {
			case 1, 4: // Four of a kind
				return 5
			}
			return 4 // Full house
		}
	case 3: // Three of a kind or two pair
		for _, l := range h {
			switch l {
			case 3: // Three of a kind
				return 3
			}
		}
		return 2 // Two pair
	case 4: // Pair
		return 1
	case 5: // High card
		return 0
	}
	panic("unparseable hand")
}
