package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	w, p := parse(data)

	log.Println("Part one:", partOne(w, p))
	log.Println("Part two:", partTwo(w))
}

func partOne(workflows map[string]workflow, parts []part) int {
	var accepted []part
	for _, p := range parts {
		name := "in"
		for w := workflows[name]; name != "R" && name != "A"; w = workflows[name] {
			for _, r := range w {
				if r.Match(p) {
					name = r.result
					break
				}
			}
		}
		if name == "A" {
			accepted = append(accepted, p)
		}
	}
	var sum int
	for _, a := range accepted {
		sum += a.x + a.m + a.a + a.s
	}
	return sum
}

func partTwo(workflows map[string]workflow) int64 {
	g := dfs(workflows)
	var sum int64
	for _, w := range g {
		min, max := w.Range()
		sum += int64(max.x-min.x+1) * int64(max.m-min.m+1) * int64(max.a-min.a+1) * int64(max.s-min.s+1)
	}
	return sum
}

func parse(input []byte) (map[string]workflow, []part) {
	blocks := bytes.Split(input, []byte{'\n', '\n'})
	workflows := make(map[string]workflow)
	for _, line := range bytes.Split(blocks[0], []byte{'\n'}) {
		var w workflow
		start := bytes.Index(line, []byte{'{'})
		name := string(line[:start])
		for _, part := range bytes.Split(line[start+1:], []byte{','}) {
			var r rule
			subparts := bytes.Split(part, []byte{':'})
			if len(subparts) == 2 {
				r.category, r.condition = string(subparts[0][:1]), subparts[0][1]
				r.comparator, _ = strconv.Atoi(string(subparts[0][2:]))
				r.result = strings.TrimSuffix(string(subparts[1]), "}")
			} else {
				r.result = strings.TrimSuffix(string(part), "}")
				r.comparator = ' '
			}
			w = append(w, r)

		}
		workflows[name] = w
	}
	workflows["A"] = workflow{{
		result:    "accepted",
		condition: ' ',
	}}
	workflows["R"] = workflow{{
		result:    "rejected",
		condition: ' ',
	}}
	var parts []part
	for _, line := range strings.Split(string(blocks[1]), "\n") {
		var p part
		fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &p.x, &p.m, &p.a, &p.s)
		parts = append(parts, p)
	}
	return workflows, parts
}

type part struct {
	x, m, a, s int
}

type workflow []rule

func (w workflow) String() string {
	var strs []string
	for _, r := range w {
		strs = append(strs, r.String())
	}
	return strings.Join(strs, ",")
}

func (w workflow) Range() (part, part) {
	minimum, maximum := part{1, 1, 1, 1}, part{4000, 4000, 4000, 4000}
	for _, r := range w {
		var upper, lower int
		switch r.condition {
		case '>':
			upper, lower = 4000, r.comparator+1
		case '<':
			upper, lower = r.comparator-1, 1
		}
		switch r.category {
		case "x":
			minimum.x, maximum.x = max(minimum.x, lower), min(maximum.x, upper)
		case "m":
			minimum.m, maximum.m = max(minimum.m, lower), min(maximum.m, upper)
		case "a":
			minimum.a, maximum.a = max(minimum.a, lower), min(maximum.a, upper)
		case "s":
			minimum.s, maximum.s = max(minimum.s, lower), min(maximum.s, upper)
		}
	}
	return minimum, maximum
}

type rule struct {
	category   string
	condition  byte
	comparator int
	result     string
}

func (r rule) String() string {
	return fmt.Sprintf("%s%s%d:%s", r.category, string(r.condition), r.comparator, r.result)
}

func (r rule) invert() rule {
	s := rule{category: r.category, result: r.result, condition: ' '}
	switch r.condition {
	case '>':
		s.condition = '<'
		s.comparator = r.comparator + 1
	case '<':
		s.condition = '>'
		s.comparator = r.comparator - 1
	}
	return s
}

func (r rule) Match(p part) bool {
	if r.category == "" {
		return true
	}
	var value int
	switch r.category {
	case "x":
		value = p.x
	case "m":
		value = p.m
	case "a":
		value = p.a
	case "s":
		value = p.s
	}
	switch r.condition {
	case '>':
		return value > r.comparator
	case '<':
		return value < r.comparator
	}
	panic("Undefined condition")
}

type ruleGraph struct {
	paths []workflow
}

func dfs(workflows map[string]workflow) []workflow {
	var path workflow
	r := &ruleGraph{}
	r.dfsInner(workflows["A"][0], "in", 0, workflows, path)
	return r.paths
}

func (r *ruleGraph) dfsInner(target rule, i string, j int, workflows map[string]workflow, path workflow) {
	if workflows[i][j] == target {
		path = append(path, workflows[i][j])
		cpy := make(workflow, len(path))
		copy(cpy, path)
		r.paths = append(r.paths, cpy)
		return
	}

	if _, ok := workflows[workflows[i][j].result]; ok {
		r.dfsInner(target, workflows[i][j].result, 0, workflows, append(path, workflows[i][j]))
	}
	if len(workflows[i]) > j+1 {
		r.dfsInner(target, i, j+1, workflows, append(path, workflows[i][j].invert()))
	}
}
