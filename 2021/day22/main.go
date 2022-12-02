package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(string(data)))
	log.Println("Part two:", partTwo(string(data)))
}

func partOne(input string) int {
	r := make(reactor)

	for _, line := range strings.Split(input, "\n") {
		execute(parse(line), r)
	}

	return r.Count()
}

func partTwo(input string) int {
	var boxes []transform
	for _, line := range strings.Split(input, "\n") {
		this := parse(line)
		var newBoxes []transform
		// Find the intersections of this box with all the boxes already parsed to figure out
		// the resultant operation to get the correct number later
		for _, b := range boxes {
			if i := b.intersection(this); i != nil {
				newBoxes = append(newBoxes, *i)
			}
		}
		// If we're switching stuff on then track it, switching off is done because it's been
		// handled by the intersections already, and if there isn't an intersection it's a no-op
		if this.operation {
			newBoxes = append(newBoxes, this)
		}
		boxes = append(boxes, newBoxes...)
	}

	var sum int
	for _, b := range boxes {
		// If the operation is "on" then we add the volume of cubes
		if b.operation {
			sum += b.volume()
			continue
		}
		// Otherwise remove
		sum -= b.volume()
	}

	return sum
}

type cube struct {
	x, y, z int
}

type transform struct {
	x, y, z   mrange
	operation bool
}

func (t transform) String() string {
	opstr := "off"
	if t.operation {
		opstr = "on"
	}
	return fmt.Sprintf("%s x=%d..%d,y=%d..%d,z=%d..%d", opstr, t.x.min, t.x.max, t.y.min, t.y.max, t.z.min, t.z.max)
}

func (t transform) intersection(r transform) *transform {
	rrange := []mrange{r.x, r.y, r.z}
	for i, v := range []mrange{t.x, t.y, t.z} {
		w := rrange[i]
		// If min is less than max, or max less than min we have no intersection
		if v.min > w.max || v.max < w.min {
			return nil
		}
	}
	// The latter operation wins if they are different (i.e. switching "off" after "on" means we need to remove the original
	// "on"s to get the correct number in total, and vice versa).
	op := r.operation
	// If the operations are the same, we need to offset the effects of doing the same operation twice, i.e. if we
	// turn on 30 cubes, of which 15 are already on, then we've only effectively turned on 15 cubes, so we need
	// to remove 15 from the total number of "on"s, and vice versa.
	if t.operation == r.operation {
		op = !t.operation
	}
	return &transform{
		mrange{max(t.x.min, r.x.min), min(t.x.max, r.x.max)},
		mrange{max(t.y.min, r.y.min), min(t.y.max, r.y.max)},
		mrange{max(t.z.min, r.z.min), min(t.z.max, r.z.max)},
		op,
	}
}

func (t *transform) volume() int {
	return (t.x.max - t.x.min + 1) * (t.y.max - t.y.min + 1) * (t.z.max - t.z.min + 1)
}

type mrange struct {
	min, max int
}

type reactor map[cube]bool

func (r reactor) Count() int {
	var sum int
	for _, v := range r {
		if v {
			sum++
		}
	}
	return sum
}

func parse(line string) transform {
	var t transform
	var opstr string
	fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &opstr, &t.x.min, &t.x.max, &t.y.min, &t.y.max, &t.z.min, &t.z.max)
	switch opstr {
	case "on":
		t.operation = true
	}
	return t
}

func execute(t transform, r reactor) {
	for x := max(-50, t.x.min); x <= min(50, t.x.max); x++ {
		for y := max(-50, t.y.min); y <= min(50, t.y.max); y++ {
			for z := max(-50, t.z.min); z <= min(50, t.z.max); z++ {
				r[cube{x, y, z}] = t.operation
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
