package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", stable(bytes.Split(data, []byte{'\n'})))
}

func stable(seabed [][]byte) int {
	changed := true
	var i int
	for changed {
		seabed, changed = iterate(seabed)
		i++
		// log.Println("\n" + string(bytes.Join(seabed, []byte{'\n'})))
	}
	return i
}

func iterate(seabed [][]byte) ([][]byte, bool) {
	var changed bool
	for _, c := range []byte{'>', 'v'} {
		var subchanged bool
		if seabed, subchanged = subiterate(seabed, c); subchanged {
			changed = true
		}
	}
	return seabed, changed
}

func subiterate(seabed [][]byte, char byte) ([][]byte, bool) {
	var changed bool
	for y := 0; y < len(seabed); y++ {
		for x := 0; x < len(seabed[0]); x++ {
			switch seabed[y][x] {
			case '.':
				continue
			case char:
				switch char {
				case '>':
					if seabed[y][next(x, len(seabed[0]))] == '.' {
						seabed[y][x], seabed[y][next(x, len(seabed[0]))] = 'c', 'a'
						changed = true
					}
				case 'v':
					if seabed[next(y, len(seabed))][x] == '.' {
						seabed[y][x], seabed[next(y, len(seabed))][x] = 'c', 'b'
						changed = true
					}
				}
			}
		}
	}
	for y := 0; y < len(seabed); y++ {
		seabed[y] = bytes.ReplaceAll(bytes.ReplaceAll(bytes.ReplaceAll(seabed[y], []byte{'a'}, []byte{'>'}), []byte{'b'}, []byte{'v'}), []byte{'c'}, []byte{'.'})
	}
	return seabed, changed
}

func next(i, max int) int {
	if i+1 == max {
		return 0
	}
	return i + 1
}
