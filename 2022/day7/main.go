package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	root := parse(string(data))

	log.Println("Part one:", partOne(root))
	log.Println("Part two:", partTwo(root))
}

func partTwo(d *dir) int {
	minSize := d.Size() - 40000000
	best := math.MaxInt
	d.Walk(func(d *dir) {
		s := d.Size()
		if s > minSize && s < best {
			best = s
		}
	})
	return best
}

func partOne(d *dir) int {
	sizes := make(map[string]int)
	d.Walk(func(d *dir) {
		sizes[d.Path()] = d.Size()
	})
	var partOne int
	for _, s := range sizes {
		if s <= 100000 {
			partOne += s
		}
	}
	return partOne
}

func (d *dir) Walk(f func(d *dir)) {
	f(d)
	for _, s := range d.subdirs {
		s.Walk(f)
	}
}

func (d dir) Path() string {
	path := []string{d.name}
	for s := d.parent; s != nil; s = s.parent {
		path = append(path, s.name)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return strings.Join(path, "/")
}

type dir struct {
	name    string
	parent  *dir
	subdirs map[string]*dir
	files   map[string]int
}

func (d dir) Size() int {
	var size int
	for _, s := range d.subdirs {
		size += s.Size()
	}
	for _, fs := range d.files {
		size += fs
	}
	return size
}

func parse(input string) *dir {
	lines := strings.Split(input, "\n")
	root := &dir{name: "/", subdirs: make(map[string]*dir), files: make(map[string]int)}
	currentDir := root
	for i := 0; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		switch fields[0] {
		case "$":
			switch fields[1] {
			case "cd":
				switch fields[2] {
				case "..":
					currentDir = currentDir.parent
				case "/":
					currentDir = root
				default:
					if _, ok := currentDir.subdirs[fields[2]]; !ok {
						currentDir.subdirs[fields[2]] = &dir{name: fields[2], parent: currentDir, subdirs: make(map[string]*dir), files: make(map[string]int)}
					}
					currentDir = currentDir.subdirs[fields[2]]
				}
			case "ls":
			default:
				panic(lines[i])
			}
		default:
			switch fields[0] {
			case "dir":
			default:
				size, _ := strconv.Atoi(fields[0])
				currentDir.files[fields[1]] = size
			}
		}
	}
	return root
}
