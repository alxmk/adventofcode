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
		log.Fatalln("Failed to read input file:", err)
	}

	moons := parseMoons(string(data))

	log.Println("Part one:", simulate(moons, 1000))
	log.Println("Part two:", searchForRepeat(moons))
}

func simulate(moons []*moon, steps int) int {
	pairs := generatePairs(moons)
	for t := 0; t < steps; t++ {
		for _, p := range pairs {
			p.ApplyGravity()
		}
		for _, m := range moons {
			m.Tick()
		}
	}

	var totalEnergy int
	for _, m := range moons {
		totalEnergy += m.KineticEnergy() * m.PotentialEnergy()
	}
	return totalEnergy
}

func searchForRepeat(moons []*moon) int {
	xstates, ystates, zstates := make(map[string]struct{}), make(map[string]struct{}), make(map[string]struct{})
	var xrepeats, yrepeats, zrepeats int
	pairs := generatePairs(moons)
	for t := 0; ; t++ {
		for _, p := range pairs {
			p.ApplyGravity()
		}
		for _, m := range moons {
			m.Tick()
		}
		xs, ys, zs := axisStates(moons)
		if _, ok := xstates[xs]; ok && xrepeats == 0 {
			xrepeats = len(xstates)
		}
		if _, ok := ystates[ys]; ok && yrepeats == 0 {
			yrepeats = len(ystates)
		}
		if _, ok := zstates[zs]; ok && zrepeats == 0 {
			zrepeats = len(zstates)
		}
		if xrepeats != 0 && yrepeats != 0 && zrepeats != 0 {
			return lowestCommonMultiple(xrepeats, yrepeats, zrepeats)
		}
		xstates[xs], ystates[ys], zstates[zs] = struct{}{}, struct{}{}, struct{}{}
	}
}

func lowestCommonMultiple(x, y, z int) int {
	pfX, pfY, pfZ := primeFactors(x), primeFactors(y), primeFactors(z)

	factors := make(map[int]int)
	for k, n := range pfX {
		if factors[k] < n {
			factors[k] = n
		}
	}
	for k, n := range pfY {
		if factors[k] < n {
			factors[k] = n
		}
	}
	for k, n := range pfZ {
		if factors[k] < n {
			factors[k] = n
		}
	}

	lcm := 1
	for k, n := range factors {
		for i := 0; i < n; i++ {
			lcm *= k
		}
	}

	return lcm
}

func primeFactors(n int) map[int]int {
	pf := make(map[int]int)
	for i := n % 2; i == 0; i = n % 2 {
		pf[2]++
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for j := n % i; j == 0; j = n % i {
			pf[i]++
			n /= i
		}
		if n == 1 {
			return pf
		}
	}
	if n > 2 {
		pf[n]++
	}
	return pf
}

func axisStates(moons []*moon) (string, string, string) {
	var x, y, z strings.Builder
	for _, m := range moons {
		x.WriteString(fmt.Sprintf("p%dv%d", m.x, m.vx))
		y.WriteString(fmt.Sprintf("p%dv%d", m.y, m.vy))
		z.WriteString(fmt.Sprintf("p%dv%d", m.z, m.vz))
	}
	return x.String(), y.String(), z.String()
}

type moon struct {
	x, y, z    int
	vx, vy, vz int
}

func (m moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>", m.x, m.y, m.z, m.vx, m.vy, m.vz)
}

func (m moon) KineticEnergy() int {
	return absint(m.vx) + absint(m.vy) + absint(m.vz)
}

func (m moon) PotentialEnergy() int {
	return absint(m.x) + absint(m.y) + absint(m.z)
}

func (m *moon) Tick() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func absint(i int) int {
	if i >= 0 {
		return i
	}
	return -1 * i
}

func parseMoons(input string) []*moon {
	var moons []*moon
	for _, line := range strings.Split(input, "\n") {
		var m moon
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &m.x, &m.y, &m.z)
		moons = append(moons, &m)
	}
	return moons
}

type pair [2]*moon

func (p pair) ApplyGravity() {
	p[0].vx, p[1].vx = applyGravity(p[0].x, p[1].x, p[0].vx, p[1].vx)
	p[0].vy, p[1].vy = applyGravity(p[0].y, p[1].y, p[0].vy, p[1].vy)
	p[0].vz, p[1].vz = applyGravity(p[0].z, p[1].z, p[0].vz, p[1].vz)
}

func applyGravity(a, b, va, vb int) (int, int) {
	if a > b {
		return va - 1, vb + 1
	}
	if a < b {
		return va + 1, vb - 1
	}
	return va, vb
}

func generatePairs(moons []*moon) []pair {
	var pairs []pair
	for i, m := range moons {
		for _, n := range moons[i+1:] {
			pairs = append(pairs, [2]*moon{m, n})
		}
	}
	return pairs
}
