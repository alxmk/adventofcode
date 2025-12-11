package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	nodes := make(map[string]*node)
	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Fields(line)
		id := parts[0][:len(parts[0])-1]
		if _, ok := nodes[id]; !ok {
			nodes[id] = &node{name: id}
		}
		for _, connected := range parts[1:] {
			if _, ok := nodes[connected]; !ok {
				nodes[connected] = &node{name: connected}
			}
			nodes[id].downstream = append(nodes[id].downstream, nodes[connected])
		}
	}
	fmt.Println("Part one:", nodes["you"].traverse("out", true, true, make(map[state]int)))
	fmt.Println("Part two:", nodes["svr"].traverse("out", false, false, make(map[state]int)))
}

type node struct {
	name       string
	downstream []*node
}

type state struct {
	node     string
	fft, dac bool
}

func (n *node) traverse(dest string, fft, dac bool, memos map[state]int) (paths int) {
	dac = dac || n.name == "dac"
	fft = fft || n.name == "fft"
	if n.name == dest {
		if fft && dac {
			return 1
		}
		return 0
	}
	for _, next := range n.downstream {
		s := state{next.name, fft, dac}
		if p, ok := memos[s]; ok {
			paths += p
			continue
		}
		memos[s] = next.traverse(dest, fft, dac, memos)
		paths += memos[s]
	}
	return paths
}
