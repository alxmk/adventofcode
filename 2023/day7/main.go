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
	log.Println("Part two:", solve(jokers(h), true))
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
		h := hand{cmap: make(map[rune]int)}
		fmt.Sscanf(line, "%s %d", &h.cards, &h.bid)
		for _, r := range h.cards {
			h.cmap[r]++
		}
		hands = append(hands, h)
	}
	return hands
}

func jokers(hands []hand) []hand {
	var j []hand
	for _, h := range hands {
		j = append(j, h.Joker())
	}
	return j
}

type hand struct {
	cards string
	cmap  map[rune]int
	bid   int
}

func (h hand) Joker() hand {
	nj, ok := h.cmap['J']
	if !ok {
		return h
	}
	delete(h.cmap, 'J')
	var pairs, singles []rune
	for c, n := range h.cmap {
		switch n {
		case 4, 3:
			// This is always the best option to bump
			h.cmap[c] += nj
			return h
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
		h.cmap[pairs[0]] += nj
		return h
	}
	if len(singles) != 0 {
		// Find the highest single and bump it
		sort.Slice(singles, func(i, j int) bool {
			return power(singles[i], false) < power(singles[j], false)
		})
		h.cmap[singles[0]] += nj
		return h
	}
	// They were all Jokers
	h.cmap['A'] += nj
	return h
}

func (h hand) Beats(j hand, jokers bool) bool {
	lh, lj := len(h.cmap), len(j.cmap)
	if lh < lj {
		return true
	}
	if lh > lj {
		return false
	}
	switch lh {
	case 1, 4, 5:
		// Five of a kind || One Pair || High card
		return tiebreak(h.cards, j.cards, jokers)
	case 2:
		// Four of a kind or full house
		for _, l := range h.cmap {
			switch l {
			case 1, 4:
				// Four of a kind
				for _, k := range j.cmap {
					switch k {
					case 1, 4:
						return tiebreak(h.cards, j.cards, jokers)
					default:
						// Four of a kind beats full house
						return true
					}
				}
			case 2, 3:
				// Full house
				for _, k := range j.cmap {
					switch k {
					case 2, 3:
						return tiebreak(h.cards, j.cards, jokers)
					default:
						// Four of a kind beats full house
						return false
					}
				}
			}
		}
	case 3:
		// Three of a kind or two pair
		// Four of a kind or full house
		for _, l := range h.cmap {
			switch l {
			case 3:
				// Three of a kind
				for _, k := range j.cmap {
					switch k {
					case 3:
						return tiebreak(h.cards, j.cards, jokers)
					case 2:
						// Three of a kind beats two pair
						return true
					case 1:
						// Could be either
						continue
					}
				}
			case 2:
				// Two pair
				for _, k := range j.cmap {
					switch k {
					case 2:
						return tiebreak(h.cards, j.cards, jokers)
					case 3:
						// Three of a kind beats two pair
						return false
					case 1:
						// Could be either
						continue
					}
				}
			case 1:
				// Could be either
				continue
			}
		}
	}
	panic(fmt.Sprintf("%s vs %s is a draw", h.cards, j.cards))
}

var cardPower = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
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
		if ph > pj {
			return true
		}
		if ph < pj {
			return false
		}
	}
	panic(fmt.Sprintf("%s vs %s is a draw", a, b))
}
