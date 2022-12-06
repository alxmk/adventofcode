package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	log.Println("Part one:", partOne())
	log.Println("Part two:", partTwo())

	// n, _ := thisProgramme.Exec(0, 0)
	// log.Println("Part one alt:", n)
}

func partOne() int64 {
	cacheA, cacheB := make(map[int]int64), make(map[int]int64)
	for x := int64(999999); x >= 111111; x-- {
		ok, i := six(strconv.FormatInt(x, 10))
		if ok {
			// We are counting down so just store the first time we hit any value of z because it
			// will be the maximum
			if _, ok := cacheA[i]; !ok {
				cacheA[i] = x
			}
		}
	}
	for x := int64(9); x >= 1; x-- {
		for r, n := range cacheA {
			ok, i := seven(strconv.FormatInt(x, 10), r)
			if ok {
				// If it's the first time we've hit this z value then store, otherwise if the
				// input is a larger number then it's a better route to this z value so overwrite
				if s, ok := cacheB[i]; !ok || (n*10+x) > s {
					cacheB[i] = n*10 + x
				}
			}
		}
	}
	cacheA = make(map[int]int64)
	for x := int64(99); x >= 11; x-- {
		for r, n := range cacheB {
			ok, i := nine(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheA[i]; !ok || (n*100+x) > s {
					cacheA[i] = n*100 + x
				}
			}
		}
	}
	cacheB = make(map[int]int64)
	for x := int64(99); x >= 11; x-- {
		for r, n := range cacheA {
			ok, i := eleven(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheB[i]; !ok || (n*100+x) > s {
					cacheB[i] = n*100 + x
				}
			}
		}
	}
	cacheA = make(map[int]int64)
	for x := int64(9); x >= 1; x-- {
		for r, n := range cacheB {
			ok, i := twelve(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheA[i]; !ok || (n*10+x) > s {
					cacheA[i] = n*10 + x
				}
			}
		}
	}
	cacheB = make(map[int]int64)
	for x := int64(9); x >= 1; x-- {
		for r, n := range cacheA {
			ok, i := thirteen(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheB[i]; !ok || (n*10+x) > s {
					cacheB[i] = n*10 + x
				}
			}
		}
	}
	cacheA = make(map[int]int64)
	for x := int64(9); x >= 1; x-- {
		for r, n := range cacheB {
			ok, i := fourteen(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheA[i]; !ok || (n*10+x) > s {
					cacheA[i] = n*10 + x
				}
			}
		}
	}
	return cacheA[0]
}

func partTwo() int64 {
	cacheA, cacheB := make(map[int]int64), make(map[int]int64)
	for x := int64(111111); x <= 999999; x++ {
		ok, i := six(strconv.FormatInt(x, 10))
		if ok {
			// We are counting up so just store the first time we hit any value of z because it
			// will be the minimum
			if _, ok := cacheA[i]; !ok {
				cacheA[i] = x
			}
		}
	}
	for x := int64(1); x <= 9; x++ {
		for r, n := range cacheA {
			ok, i := seven(strconv.FormatInt(x, 10), r)
			if ok {
				// If it's the first time we've hit this z value then store, otherwise if the
				// input is a smaller number then it's a better route to this z value so overwrite
				if s, ok := cacheB[i]; !ok || (n*10+x) < s {
					cacheB[i] = n*10 + x
				}
			}
		}
	}
	cacheA = make(map[int]int64)
	for x := int64(11); x <= 99; x++ {
		for r, n := range cacheB {
			ok, i := nine(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheA[i]; !ok || (n*100+x) < s {
					cacheA[i] = n*100 + x
				}
			}
		}
	}
	cacheB = make(map[int]int64)
	for x := int64(11); x <= 99; x++ {
		for r, n := range cacheA {
			ok, i := eleven(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheB[i]; !ok || (n*100+x) < s {
					cacheB[i] = n*100 + x
				}
			}
		}
	}
	cacheA = make(map[int]int64)
	for x := int64(1); x <= 9; x++ {
		for r, n := range cacheB {
			ok, i := twelve(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheA[i]; !ok || (n*10+x) < s {
					cacheA[i] = n*10 + x
				}
			}
		}
	}
	cacheB = make(map[int]int64)
	for x := int64(1); x <= 9; x++ {
		for r, n := range cacheA {
			ok, i := thirteen(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheB[i]; !ok || (n*10+x) < s {
					cacheB[i] = n*10 + x
				}
			}
		}
	}
	cacheA = make(map[int]int64)
	for x := int64(1); x <= 9; x++ {
		for r, n := range cacheB {
			ok, i := fourteen(strconv.FormatInt(x, 10), r)
			if ok {
				if s, ok := cacheA[i]; !ok || (n*10+x) < s {
					cacheA[i] = n*10 + x
				}
			}
		}
	}
	return cacheA[0]
}

func six(input string) (bool, int) {
	if strings.Contains(input, "0") {
		return false, -1
	}
	var w, x, z int
	w = int(input[0] - '0')
	z += w + 2
	w = int(input[1] - '0')
	z *= 26
	z += w + 4
	w = int(input[2] - '0')
	z *= 26
	z += w + 8
	w = int(input[3] - '0')
	z *= 26
	z += w + 7
	w = int(input[4] - '0')
	z *= 26
	z += w + 12
	w = int(input[5] - '0')
	x = (z % 26) - 14
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

func seven(input string, z int) (bool, int) {
	var w, x int
	w = int(input[0] - '0')
	x = z % 26
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

func nine(input string, z int) (bool, int) {
	if strings.Contains(input, "0") {
		return false, -1
	}
	var w, x int
	w = int(input[0] - '0')
	z *= 26
	z += w + 14
	w = int(input[1] - '0')
	x = (z % 26) - 10
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

func eleven(input string, z int) (bool, int) {
	if strings.Contains(input, "0") {
		return false, -1
	}
	var w, x int
	w = int(input[0] - '0')
	z *= 26
	z += w + 6
	w = int(input[1] - '0')
	x = (z % 26) - 12
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

func twelve(input string, z int) (bool, int) {
	var w, x int
	w = int(input[0] - '0')
	x = (z % 26) - 3
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

func thirteen(input string, z int) (bool, int) {
	var w, x int
	w = int(input[0] - '0')
	x = (z % 26) - 11
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

func fourteen(input string, z int) (bool, int) {
	var w, x int
	w = int(input[0] - '0')
	x = (z % 26) - 2
	z /= 26
	if x != w {
		return false, z
	}
	return true, z
}

type block struct {
	zdiv, xdiff, ydiff int
}

func (p programme) Exec(depth, z int) (bool, int) {
	if c, ok := p.cache[depth]; ok {
		if _, ok := c[z]; ok {
			log.Println("Cache exit", depth)
			return false, z
		}
	}
	var ok bool
	var newz int
	for input := 9; input >= 1; input-- {
		log.Println(depth, input)
		if ok, newz = p.blocks[depth].Run(input, z); ok {
			if depth == 5 {
				// Winner winner chicken dinner, maybe
				return true, z
			}
			if ok, newz = p.Exec(depth+1, newz); ok {
				return ok, newz
			}
		}
	}
	log.Println("Exit", depth)
	// This doesn't work, cache big values only
	if depth < 10 {
		if _, ok := p.cache[depth]; !ok {
			p.cache[depth] = make(map[int]struct{})
		}
		p.cache[depth][z] = struct{}{}
	}
	return false, z
}

func (b *block) Run(w, z int) (bool, int) {
	x := (z % 26) + b.xdiff
	if b.zdiv != 1 && x != w {
		return false, z
	}
	z /= b.zdiv
	z += w + b.ydiff
	return true, z
}

type programme struct {
	blocks []block
	cache  map[int]map[int]struct{} // map of depth to zs that don't work
}

var thisProgramme = programme{
	blocks: []block{
		{1, 10, 2},
		{1, 10, 4},
		{1, 14, 8},
		{1, 11, 7},
		{1, 14, 12},
		{26, -14, 7},
		{26, 0, 10},
		{1, 10, 14},
		{26, -10, 2},
		{1, 13, 6},
		{26, -12, 8},
		{26, -3, 11},
		{26, -11, 5},
		{26, -2, 11},
	},
	cache: map[int]map[int]struct{}{},
}
