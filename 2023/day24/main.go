package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parseHailstones(string(data)), [2]float64{200000000000000, 200000000000000}, [2]float64{400000000000000, 400000000000000}))
	log.Println("Part two:", partTwo(parseHailstones(string(data))))
}

func partOne(hailstones []hailstone, min, max [2]float64) int {
	var count int
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if is, ok := hailstones[i].intersects2d(hailstones[j]); ok &&
				is[0] >= min[0] &&
				is[0] <= max[0] &&
				is[1] >= min[0] &&
				is[1] <= max[0] {
				count++
			}
		}
	}
	return count
}

func partTwo(h []hailstone) int64 {
	// Method inspired by https://www.reddit.com/r/adventofcode/comments/18pnycy/comment/kepu26z/
	s1 := h[1].p.cross(h[1].v).minus(h[0].p.cross(h[0].v))
	s2 := h[2].p.cross(h[2].v).minus(h[0].p.cross(h[0].v))

	rhs := [6]float64{s1[0], s1[1], s1[2], s2[0], s2[1], s2[2]}

	m00 := h[0].v.crossMatrix().minus(h[1].v.crossMatrix())
	m03 := h[0].v.crossMatrix().minus(h[2].v.crossMatrix())
	m30 := h[1].p.crossMatrix().minus(h[0].p.crossMatrix())
	m33 := h[2].p.crossMatrix().minus(h[0].p.crossMatrix())

	m := mat.NewDense(6, 6, []float64{
		m00[0][0], m00[0][1], m00[0][2], m30[0][0], m30[0][1], m30[0][2],
		m00[1][0], m00[1][1], m00[1][2], m30[1][0], m30[1][1], m30[1][2],
		m00[2][0], m00[2][1], m00[2][2], m30[2][0], m30[2][1], m30[2][2],
		m03[0][0], m03[0][1], m03[0][2], m33[0][0], m33[0][1], m33[0][2],
		m03[1][0], m03[1][1], m03[1][2], m33[1][0], m33[1][1], m33[1][2],
		m03[2][0], m03[2][1], m03[2][2], m33[2][0], m33[2][1], m33[2][2],
	})
	m.Inverse(m)
	b := bigmatrix{
		{m.At(0, 0), m.At(0, 1), m.At(0, 2), m.At(0, 3), m.At(0, 4), m.At(0, 5)},
		{m.At(1, 0), m.At(1, 1), m.At(1, 2), m.At(1, 3), m.At(1, 4), m.At(1, 5)},
		{m.At(2, 0), m.At(2, 1), m.At(2, 2), m.At(2, 3), m.At(2, 4), m.At(2, 5)},
		{m.At(3, 0), m.At(3, 1), m.At(3, 2), m.At(3, 3), m.At(3, 4), m.At(3, 5)},
		{m.At(4, 0), m.At(4, 1), m.At(4, 2), m.At(4, 3), m.At(4, 4), m.At(4, 5)},
		{m.At(5, 0), m.At(5, 1), m.At(5, 2), m.At(5, 3), m.At(5, 4), m.At(5, 5)},
	}
	r := b.mul(rhs)
	return int64(r[0] + r[1] + r[2])
}

func parseHailstones(input string) []hailstone {
	var hailstones []hailstone
	for _, line := range strings.Split(input, "\n") {
		var h hailstone
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &h.p[0], &h.p[1], &h.p[2], &h.v[0], &h.v[1], &h.v[2])
		hailstones = append(hailstones, h)
	}
	return hailstones
}

type hailstone struct {
	p, v vector
}

func (h hailstone) intersects2d(i hailstone) ([3]float64, bool) {
	var intersection [3]float64
	x1, y1, x2, y2 := h.p[0], h.p[1], h.v[0]+h.p[0], h.v[1]+h.p[1]
	x3, y3, x4, y4 := i.p[0], i.p[1], i.v[0]+i.p[0], i.v[1]+i.p[1]

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) /
		((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))
	u := ((x1-x3)*(y1-y2) - (y1-y3)*(x1-x2)) /
		((x1-x2)*(y3-y4) - (y1-y2)*(x3-x4))

	intersection[0] = x1 + t*(x2-x1)
	intersection[1] = y1 + t*(y2-y1)

	return intersection, t > 0 && u > 0
}

type vector [3]float64

func (a vector) cross(b vector) vector {
	return vector{a[1]*b[2] - a[2]*b[1], a[2]*b[0] - a[0]*b[2], a[0]*b[1] - a[1]*b[0]}
}

func (v vector) minus(w vector) vector {
	return vector{v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}

type matrix [3]vector

func (v vector) crossMatrix() matrix {
	return matrix{
		{0, -v[2], v[1]},
		{v[2], 0, -v[0]},
		{-v[1], v[0], 0},
	}
}

func (m matrix) minus(n matrix) matrix {
	return matrix{
		{m[0][0] - n[0][0], m[0][1] - n[0][1], m[0][2] - n[0][2]},
		{m[1][0] - n[1][0], m[1][1] - n[1][1], m[1][2] - n[1][2]},
		{m[2][0] - n[2][0], m[2][1] - n[2][1], m[2][2] - n[2][2]},
	}
}

type bigmatrix [6][6]float64

func (a bigmatrix) mul(b [6]float64) [6]float64 {
	return [6]float64{
		a[0][0]*b[0] + a[0][1]*b[1] + a[0][2]*b[2] + a[0][3]*b[3] + a[0][4]*b[4] + a[0][5]*b[5],
		a[1][0]*b[0] + a[1][1]*b[1] + a[1][2]*b[2] + a[1][3]*b[3] + a[1][4]*b[4] + a[1][5]*b[5],
		a[2][0]*b[0] + a[2][1]*b[1] + a[2][2]*b[2] + a[2][3]*b[3] + a[2][4]*b[4] + a[2][5]*b[5],
		a[3][0]*b[0] + a[3][1]*b[1] + a[3][2]*b[2] + a[3][3]*b[3] + a[3][4]*b[4] + a[3][5]*b[5],
		a[4][0]*b[0] + a[4][1]*b[1] + a[4][2]*b[2] + a[4][3]*b[3] + a[4][4]*b[4] + a[4][5]*b[5],
		a[5][0]*b[0] + a[5][1]*b[1] + a[5][2]*b[2] + a[5][3]*b[3] + a[5][4]*b[4] + a[5][5]*b[5],
	}
}
