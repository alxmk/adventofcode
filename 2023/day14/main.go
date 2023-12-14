package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", parsePlatform(string(data)).Tilt(n).Load())
	log.Println("Part two:", partTwo(parsePlatform(string(data)), 1000000000))
}

func partTwo(p *platform, spins int) int {
	cache := make(map[string]int)
	var repeat bool
	for i := 0; i < spins; i++ {
		p.Spin()
		s := p.String()
		if j, ok := cache[s]; ok && !repeat {
			i = spins - j
			repeat = true
		}
		cache[s] = i
	}
	return p.Load()
}

func parsePlatform(input string) *platform {
	p := &platform{rocks: make(map[xy]rune)}
	for y, line := range strings.Split(input, "\n") {
		p.max.y = y
		for x, r := range line {
			p.rocks[xy{x, y}] = r
			p.max.x = x
		}
	}
	return p
}

var (
	n = xy{0, -1}
	s = xy{0, 1}
	e = xy{1, 0}
	w = xy{-1, 0}
)

type xy struct {
	x, y int
}

type platform struct {
	rocks map[xy]rune
	max   xy
}

func (p platform) String() string {
	var b strings.Builder
	for y := 0; y <= p.max.y; y++ {
		for x := 0; x <= p.max.x; x++ {
			b.WriteRune(p.rocks[xy{x, y}])
		}
		if y != p.max.y {
			b.WriteRune('\n')
		}
	}
	return b.String()
}

func (p *platform) Spin() *platform {
	for _, d := range []xy{n, w, s, e} {
		p.Tilt(d)
	}
	return p
}

func (p *platform) Tilt(dir xy) *platform {
	switch dir {
	case n, w:
		p.tiltInner(xy{0, 0}, xy{p.max.x + 1, p.max.y + 1}, xy{1, 1}, dir)
	case e, s:
		p.tiltInner(xy{p.max.x, p.max.y}, xy{-1, -1}, xy{-1, -1}, dir)
	}
	return p
}

func (p *platform) tiltInner(lower, upper, inc, dir xy) {
	for x := lower.x; x != upper.x; x += inc.x {
		for y := lower.y; y != upper.y; y += inc.y {
			if p.rocks[xy{x, y}] != 'O' {
				continue
			}
			p.rocks[xy{x, y}] = '.'
			for j := (xy{x, y}); ; j = (xy{j.x + dir.x, j.y + dir.y}) {
				if p.rocks[j] == '.' {
					continue
				}
				p.rocks[xy{j.x - dir.x, j.y - dir.y}] = 'O'
				break
			}
		}
	}
}

func (p *platform) Load() int {
	var load int
	for pos, r := range p.rocks {
		if r != 'O' {
			continue
		}
		load += p.max.y - pos.y + 1
	}
	return load
}
