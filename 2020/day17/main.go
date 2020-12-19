package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	u := parse3d(string(data))

	for i := 0; i < 6; i++ {
		u = u.Next()
	}

	log.Println("Part one:", u.Count())

	v := parse4d(string(data))
	for i := 0; i < 6; i++ {
		v = v.Next()
	}

	log.Println("Part two:", v.Count())
}

func parse3d(input string) *universe3d {
	u := &universe3d{
		cubes: make(map[int]map[int]map[int]bool),
	}
	for y, line := range strings.Split(input, "\n") {
		if y > u.ymax {
			u.ymax = y
		}
		for x, c := range line {
			if _, ok := u.cubes[x]; !ok {
				u.cubes[x] = map[int]map[int]bool{}
			}
			if _, ok := u.cubes[x][y]; !ok {
				u.cubes[x][y] = make(map[int]bool)
			}
			u.cubes[x][y][0] = c == '#'
			if x > u.xmax {
				u.xmax = x
			}
		}
	}
	return u
}

type universe3d struct {
	cubes                              map[int]map[int]map[int]bool
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func (u *universe3d) Next() *universe3d {
	next := &universe3d{
		cubes: make(map[int]map[int]map[int]bool),
		xmin:  u.xmin - 1,
		xmax:  u.xmax + 1,
		ymin:  u.ymin - 1,
		ymax:  u.ymax + 1,
		zmin:  u.zmin - 1,
		zmax:  u.zmax + 1,
	}
	for x := next.xmin; x <= next.xmax; x++ {
		if _, ok := next.cubes[x]; !ok {
			next.cubes[x] = make(map[int]map[int]bool)
		}
		for y := next.ymin; y <= next.ymax; y++ {
			if _, ok := next.cubes[x][y]; !ok {
				next.cubes[x][y] = make(map[int]bool)
			}
			for z := next.zmin; z <= next.zmax; z++ {
				var activeNeighbours int
				for i := x - 1; i <= x+1; i++ {
					for j := y - 1; j <= y+1; j++ {
						for k := z - 1; k <= z+1; k++ {
							// Skip self
							if i == x && j == y && k == z {
								continue
							}
							if u.IsActive(i, j, k) {
								// log.Println(x, y, z, "neighbour", i, j, k, "is active")
								activeNeighbours++
							}
						}
					}
				}
				if u.IsActive(x, y, z) {
					if activeNeighbours == 2 || activeNeighbours == 3 {
						next.cubes[x][y][z] = true
					}
				} else {
					if activeNeighbours == 3 {
						next.cubes[x][y][z] = true
					}
				}
			}
		}
	}
	return next
}

func (u universe3d) IsActive(x, y, z int) bool {
	if _, ok := u.cubes[x]; ok {
		if _, ok := u.cubes[x][y]; ok {
			return u.cubes[x][y][z]
		}
	}
	return false
}

func (u universe3d) Count() int {
	var count int
	for x := range u.cubes {
		for y := range u.cubes[x] {
			for z := range u.cubes[x][y] {
				if u.IsActive(x, y, z) {
					count++
				}
			}
		}
	}
	return count
}

func parse4d(input string) *universe4d {
	u := &universe4d{
		cubes: make(map[int]map[int]map[int]map[int]bool),
	}
	for y, line := range strings.Split(input, "\n") {
		if y > u.ymax {
			u.ymax = y
		}
		for x, c := range line {
			if _, ok := u.cubes[x]; !ok {
				u.cubes[x] = map[int]map[int]map[int]bool{}
			}
			if _, ok := u.cubes[x][y]; !ok {
				u.cubes[x][y] = make(map[int]map[int]bool)
			}
			if _, ok := u.cubes[x][y][0]; !ok {
				u.cubes[x][y][0] = make(map[int]bool)
			}
			u.cubes[x][y][0][0] = c == '#'
			if x > u.xmax {
				u.xmax = x
			}
		}
	}
	return u
}

type universe4d struct {
	cubes                                          map[int]map[int]map[int]map[int]bool
	xmin, xmax, ymin, ymax, zmin, zmax, wmin, wmax int
}

func (u *universe4d) Next() *universe4d {
	next := &universe4d{
		cubes: make(map[int]map[int]map[int]map[int]bool),
		xmin:  u.xmin - 1,
		xmax:  u.xmax + 1,
		ymin:  u.ymin - 1,
		ymax:  u.ymax + 1,
		zmin:  u.zmin - 1,
		zmax:  u.zmax + 1,
		wmin:  u.wmin - 1,
		wmax:  u.wmax + 1,
	}
	for x := next.xmin; x <= next.xmax; x++ {
		if _, ok := next.cubes[x]; !ok {
			next.cubes[x] = make(map[int]map[int]map[int]bool)
		}
		for y := next.ymin; y <= next.ymax; y++ {
			if _, ok := next.cubes[x][y]; !ok {
				next.cubes[x][y] = make(map[int]map[int]bool)
			}
			for z := next.zmin; z <= next.zmax; z++ {
				if _, ok := next.cubes[x][y][z]; !ok {
					next.cubes[x][y][z] = make(map[int]bool)
				}
				for w := next.wmin; w <= next.wmax; w++ {
					var activeNeighbours int
				outer:
					for i := x - 1; i <= x+1; i++ {
						for j := y - 1; j <= y+1; j++ {
							for k := z - 1; k <= z+1; k++ {
								for l := w - 1; l <= w+1; l++ {
									// Skip self
									if i == x && j == y && k == z && l == w {
										continue
									}
									if u.IsActive(i, j, k, l) {
										activeNeighbours++
										if activeNeighbours == 4 {
											break outer
										}
									}
								}
							}
						}
					}
					if u.IsActive(x, y, z, w) {
						if activeNeighbours == 2 || activeNeighbours == 3 {
							next.cubes[x][y][z][w] = true
						}
					} else {
						if activeNeighbours == 3 {
							next.cubes[x][y][z][w] = true
						}
					}
				}
			}
		}
	}
	return next
}

func (u universe4d) IsActive(x, y, z, w int) bool {
	if _, ok := u.cubes[x]; ok {
		if _, ok := u.cubes[x][y]; ok {
			if _, ok := u.cubes[x][y][z]; ok {
				return u.cubes[x][y][z][w]
			}
		}
	}
	return false
}

func (u universe4d) Count() int {
	var count int
	for x := range u.cubes {
		for y := range u.cubes[x] {
			for z := range u.cubes[x][y] {
				for w := range u.cubes[x][y][z] {
					if u.IsActive(x, y, z, w) {
						count++
					}
				}
			}
		}
	}
	return count
}
