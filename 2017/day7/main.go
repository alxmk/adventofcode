package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	parttwo(partone())
}

type process struct {
	name       string
	weight     int
	children   []*process
	childnames []string
}

func (p *process) Weight() int {
	weight := p.weight

	for _, c := range p.children {
		weight += c.Weight()
	}

	return weight
}

func getNames(input string) []string {
	if input == "" {
		return []string{}
	}

	return strings.Split(input, ", ")
}

var procRegexp = regexp.MustCompile(`(?P<procname>[a-z]+) \((?P<weight>[0-9]+)\)( -> (?P<children>[a-z, ]+))?`)

func parttwo(root string) {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Error reading input file", err)
	}

	processesByName := make(map[string]*process)

	for _, line := range strings.Split(string(data), "\n") {
		matches := procRegexp.FindStringSubmatch(line)
		w, _ := strconv.Atoi(matches[2])
		p := &process{
			name:       matches[1],
			weight:     w,
			childnames: getNames(matches[4]),
		}
		processesByName[p.name] = p
	}

	for _, proc := range processesByName {
		for _, c := range proc.childnames {
			if p, ok := processesByName[c]; ok {
				proc.children = append(proc.children, p)
			} else {
				fmt.Println("Couldn't find process", c)
			}
		}
	}

	current := root
	for {
		proc := processesByName[current]

		weights := make(map[int][]string)
		for _, c := range proc.children {
			w := c.Weight()
			weights[w] = append(weights[w], c.name)
			fmt.Println(c.name, w)
		}

		if len(weights) != 1 {
			for _, c := range weights {
				if len(c) == 1 {
					current = c[0]
					break
				}
			}
		} else {
			fmt.Printf("Done at %#v", proc)
			break
		}
	}
}

func partone() string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	processes := strings.Split(string(input), "\n")

	holders := []string{}
	holdees := make(map[string]struct{})

	for _, p := range processes {
		procName, children := parse(p)
		if len(children) != 0 {
			holders = append(holders, procName)
		}

		for _, c := range children {
			holdees[c] = struct{}{}
		}
	}

	for _, h := range holders {
		if _, ok := holdees[h]; !ok {
			log.Println("The answer is", h)
			return h
		}
	}

	return ""
}

func parse(line string) (string, []string) {
	processName := strings.Fields(line)[0]
	children := []string{}

	if parts := strings.Split(line, "->"); len(parts) == 2 {
		children = strings.Fields(strings.Replace(parts[1], ",", "", -1))
	}

	return processName, children
}
