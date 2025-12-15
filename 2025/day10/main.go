package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var machines []machine
	for line := range strings.SplitSeq(input, "\n") {
		machines = append(machines, parse(line))
	}
	var partOne int
	var partTwo int
	for _, m := range machines {
		fmt.Println("Solving", m.joltage)
		partOne += m.solveLights()
		partTwo += m.solveJoltage()
	}
	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}

var maxVars int

func parse(line string) machine {
	var m machine
	for field := range strings.FieldsSeq(line) {
		switch field[0] {
		case '[':
			var n uint16
			for i := 1; i < len(field)-1; i++ {
				var bit uint16
				if field[i] == '#' {
					bit = 1
				}
				m.lights += bit << n
				n++
			}
		case '(':
			var b uint16
			var jb []int
			for id := range strings.SplitSeq(field[1:len(field)-1], ",") {
				v, _ := strconv.Atoi(id)
				b += 1 << v
				jb = append(jb, v)
			}
			m.buttons = append(m.buttons, b)
		case '{':
			for joltage := range strings.SplitSeq(field[1:len(field)-1], ",") {
				v, _ := strconv.Atoi(joltage)
				m.joltage = append(m.joltage, v)
			}
		}
		maxVars = max(maxVars, len(m.buttons))
	}
	return m
}

type machine struct {
	lights  uint16
	buttons []uint16
	joltage []int
}

func (m machine) solveLights() int {
	states := map[uint16]uint16{0: math.MaxUint16}
	for count := 0; ; count++ {
		next := make(map[uint16]uint16)
		for state, bmask := range states {
			for i, button := range m.buttons {
				if (1<<i)&bmask == 0 {
					continue
				}
				if ns := state ^ button; ns != m.lights {
					next[ns] = bmask - (1 << i)
					continue
				}
				return count + 1
			}
		}
		states = next
	}
}

func (m machine) solveJoltage() int {
	// Credit for this solution to Martin Benda https://github.com/bendiscz/aoc/blob/master/2025/10/2025-10.go
	// I couldn't solve it myself so I adapted his code.
	vars := make([]variable, len(m.buttons))
	for i := range vars {
		vars[i].max = math.MaxInt
	}

	eqs := make([]linear, len(m.joltage))
	for i, jolt := range m.joltage {
		eq := linear{b: float64(-jolt), a: make([]float64, maxVars)}
		for j, b := range m.buttons {
			if b&(1<<i) != 0 {
				eq.a[j] = 1
				vars[j].max = min(vars[j].max, jolt)
			}
		}
		eqs[i] = eq
	}

	for i := range vars {
		vars[i].free = true

		for _, eq := range eqs {
			if expr, ok := extract(eq, i); ok {
				vars[i].free = false
				vars[i].expr = expr

				for j := range eqs {
					eqs[j] = substitute(eqs[j], i, expr)
				}

				break
			}
		}
	}

	free := []int(nil)
	for i, v := range vars {
		if v.free {
			free = append(free, i)
		}
	}

	best, _ := evalRecursive(vars, free, 0)
	return best
}

const (
	errorMargin = 1e-8
)

// a_0*x_0 + a_1*x_1 + ... + a_N*x_N + b = 0
type linear struct {
	a []float64
	b float64
}

type variable struct {
	expr linear
	free bool
	val  int
	max  int
}

func extract(lin linear, index int) (linear, bool) {
	a := -lin.a[index]
	if math.Abs(a) < errorMargin {
		return linear{}, false
	}

	r := linear{b: lin.b / a, a: make([]float64, maxVars)}
	for i := range maxVars {
		if i != index {
			r.a[i] = lin.a[i] / a
		}
	}
	return r, true
}

func substitute(lin linear, index int, expr linear) linear {
	r := linear{a: make([]float64, maxVars)}

	a := lin.a[index]
	lin.a[index] = 0

	for i := range maxVars {
		r.a[i] = lin.a[i] + a*expr.a[i]
	}
	r.b = lin.b + a*expr.b
	return r
}

func eval(v variable, vals []int) float64 {
	if v.free {
		return float64(v.val)
	}

	x := v.expr.b
	for i := 0; i < maxVars; i++ {
		x += v.expr.a[i] * float64(vals[i])
	}
	return x
}

func evalRecursive(vars []variable, free []int, index int) (int, bool) {
	if index == len(free) {
		vals := make([]int, maxVars)
		total := 0

		for i := len(vars) - 1; i >= 0; i-- {
			x := eval(vars[i], vals)
			if x < -errorMargin || math.Abs(x-math.Round(x)) > errorMargin {
				return 0, false
			}
			vals[i] = int(math.Round(x))
			total += vals[i]
		}

		return total, true
	}

	best, found := math.MaxInt, false
	for x := 0; x <= vars[free[index]].max; x++ {
		vars[free[index]].val = x
		total, ok := evalRecursive(vars, free, index+1)

		if ok {
			found = true
			best = min(best, total)
		}
	}

	return best, found
}
