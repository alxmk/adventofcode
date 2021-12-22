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

	positions, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", play(positions))

	positions, err = parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	w := playDirac(positions, [2]int{})
	winner := w[0]
	if w[1] > winner {
		winner = w[1]
	}
	log.Println("Part two:", winner)
}

func parse(input string) ([2]int, error) {
	var positions [2]int
	for i, line := range strings.Split(input, "\n") {
		var dummy int
		if _, err := fmt.Sscanf(line, "Player %d starting position: %d", &dummy, &positions[i]); err != nil {
			return positions, fmt.Errorf("failed to parse line %s: %s", line, err)
		}
	}
	return positions, nil
}

type deterministic struct {
	value int
	rolls int
}

func (d *deterministic) Roll() int {
	d.value++
	d.rolls++
	if d.value > 100 {
		d.value = 1
	}
	return d.value
}

func play(positions [2]int) int {
	d := &deterministic{}
	var scores [2]int
	for {
		for j := 0; j < 2; j++ {
			positions[j] = newPosition(positions[j], d.Roll()+d.Roll()+d.Roll())
			scores[j] += positions[j]
			if scores[j] >= 1000 {
				if j == 0 {
					return scores[1] * d.rolls
				}
				return scores[0] * d.rolls
			}
		}
	}
}

func newPosition(p, roll int) int {
	if np := (p + roll) % 10; np != 0 {
		return np
	}
	return 10
}

func playDirac(positions, scores [2]int) [2]int64 {
	var wins [2]int64
	for roll, count := range diracOutcomes {
		np := newPosition(positions[0], roll)
		ns := scores[0] + np
		if ns >= 21 {
			wins[0] += count
			continue
		}
		newWins := playDirac([2]int{positions[1], np}, [2]int{scores[1], ns})
		wins[0] += newWins[1] * count
		wins[1] += newWins[0] * count
	}
	return wins
}

// i.e. in 1 universe out of 27 after the 3 rolls, the score was 3, in 3 universes it's 4, etc
var diracOutcomes = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}
