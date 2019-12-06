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

	objects, err := parseObjects(string(data))
	if err != nil {
		log.Fatalln("Failed to parse objects:", err)
	}

	log.Println("Part one:", objects.TraverseAllOrbits())
	log.Println("Part two:", objects.HopsBetween(objects["YOU"].orbits.name, objects["SAN"].orbits.name))
}

type object struct {
	orbits *object
	name   string
}

type universe map[string]*object

func parseObjects(input string) (universe, error) {
	objects := make(universe)
	for _, line := range strings.Split(input, "\n") {
		objs := strings.Split(line, ")")
		if len(objs) != 2 {
			return nil, fmt.Errorf("malformed line in input, expect XXX)YYY, got: %s", line)
		}
		if _, ok := objects[objs[0]]; !ok {
			objects[objs[0]] = &object{name: objs[0]}
		}
		if _, ok := objects[objs[1]]; !ok {
			objects[objs[1]] = &object{name: objs[1]}
		}
		objects[objs[1]].orbits = objects[objs[0]]
	}
	return objects, nil
}

func (u universe) TraverseAllOrbits() int {
	var totalOrbits int
	for obj := range u {
		totalOrbits += u.TraverseOrbits(obj)
	}
	return totalOrbits
}

// TraverseOrbits finds the number of orbits and suborbits of an object
func (u universe) TraverseOrbits(origin string) int {
	obj := u[origin]
	if obj.orbits == nil {
		return 0
	}
	return u.TraverseOrbits(obj.orbits.name) + 1
}

func (u universe) HopsBetween(start, end string) int {
	// Find paths to root
	startPath := u.PathToRoot(start, nil)
	endPath := u.PathToRoot(end, nil)

	// Find the closest object in both paths
	for i, startObj := range startPath {
		for j, endObj := range endPath {
			if startObj == endObj {
				return i + j
			}
		}
	}

	return -1
}

func (u universe) PathToRoot(start string, path []string) []string {
	if u[start].orbits == nil {
		return path
	}
	return u.PathToRoot(u[start].orbits.name, append(path, start))
}
