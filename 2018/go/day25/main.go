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
		log.Fatalln("Failed to read input", err)
	}

	ps := parsePoints(strings.Split(string(data), "\n"))

	log.Println(ps)

	cons := ps.Constellations(3)

	log.Println("Part one:", len(cons))

}

func parsePoints(input []string) points {
	var ps points
	for _, line := range input {
		var p point
		fmt.Sscanf(line, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)
		ps = append(ps, &p)
	}
	return ps
}

type point struct {
	x, y, z, t int
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", p.x, p.y, p.z, p.t)
}

func (p *point) Distance(o *point) int {
	return absInt(p.x-o.x) + absInt(p.y-o.y) + absInt(p.z-o.z) + absInt(p.t-o.t)
}

type points []*point

func (p points) String() string {
	var out strings.Builder
	for _, pt := range p {
		out.WriteString(fmt.Sprintf("%v\n", pt))
	}

	return out.String()
}

type constellations []points

func (c constellations) String() string {
	var out strings.Builder

	for _, co := range c {
		out.WriteString(fmt.Sprintf("[%v]\n", co))
	}

	return out.String()
}

func (p points) Constellations(dist int) constellations {
	var cons constellations

	for i, pa := range p {
		if pa == nil {
			continue
		}

		constellation := points{pa}

		for j := 0; j < len(p); j++ {
			pb := p[j]
			if i == j || pb == nil {
				continue
			}

			var inconst bool
			for _, e := range constellation {
				if e.Distance(pb) <= dist {
					inconst = true
					break
				}
			}
			if inconst {
				constellation = append(constellation, pb)
				p[j] = nil
				// Loop back around
				j = -1
			}
		}
		p[i] = nil

		cons = append(cons, constellation)
	}

	return cons
}

func absInt(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}
