package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	log.Println(9, 25, playGame(9, 25))
	log.Println(10, 1618, playGame(10, 1618))
	log.Println(13, 7999, playGame(13, 7999))
	log.Println(17, 1104, playGame(17, 1104))
	log.Println(21, 6661, playGame(21, 6661))
	log.Println(30, 5807, playGame(30, 5807))
	log.Println(455, 71223, playGame(455, 71223))
	log.Println(455, 7122300, playGame(455, 7122300))
}

type marble struct {
	value    int
	next     *marble
	previous *marble
}

// playGame with specified players and maximum marble value, returning the highest score
func playGame(players, maxMarble int) int {
	// Scores
	scores := make([]int, players)

	original := &marble{value: 0}
	original.next, original.previous = original, original

	current := original

	for value := 1; value <= maxMarble; value++ {
		if value%23 == 0 {
			seventh := current.previous.previous.previous.previous.previous.previous.previous
			scores[value%players] += value + seventh.value
			seventh.previous.next = seventh.next
			seventh.next.previous = seventh.previous
			current = seventh.next
			continue
		}

		clockwise1 := current.next
		clockwise2 := current.next.next

		newMarble := &marble{value: value, next: clockwise2, previous: clockwise1}
		clockwise1.next = newMarble
		clockwise2.previous = newMarble

		current = newMarble
		// original.Print(current)
	}

	sort.Sort(sort.IntSlice(scores))

	return scores[len(scores)-1]
}

func (origin *marble) Print(current *marble) {
	var printedOrigin bool
	for marb := origin; ; marb = marb.next {
		if marb.value == origin.value {
			if !printedOrigin {
				printedOrigin = true
			} else {
				break
			}
		}
		if marb.value == current.value {
			fmt.Printf("(%d) ", marb.value)
			continue
		}
		fmt.Printf("%d ", marb.value)
	}
	fmt.Print("\n")
}
