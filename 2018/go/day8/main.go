package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Println("Failed to read input", err)
	}

	s := &Scanner{
		r: bytes.NewReader(data),
	}

	tree := parseNode(s)

	log.Println(sumMetadata(tree))

	log.Println(sumValue(tree))
}

type Scanner struct {
	r *bytes.Reader
}

func (s *Scanner) Read() (int, bool) {
	var chars []rune
	for {
		char, _, err := s.r.ReadRune()
		if err != nil || char == ' ' {
			n, err := strconv.Atoi(string(chars))
			if err != nil {
				return 0, false
			}
			return n, true
		}
		chars = append(chars, char)
	}
}

type Node struct {
	numChildren  int
	numMetadatas int
	children     []*Node
	metadatas    []int
}

func parseNode(s *Scanner) *Node {
	n, ok := s.Read()
	if !ok {
		return nil
	}

	node := &Node{
		numChildren: n,
	}

	m, ok := s.Read()
	if !ok {
		return nil
	}

	node.numMetadatas = m

	for i := 0; i < n; i++ {
		node.children = append(node.children, parseNode(s))
	}

	for i := 0; i < m; i++ {
		metadata, ok := s.Read()
		if !ok {
			return nil
		}
		node.metadatas = append(node.metadatas, metadata)
	}

	return node
}

func sumMetadata(n *Node) int {
	var sum int
	for _, c := range n.children {
		sum += sumMetadata(c)
	}
	for _, m := range n.metadatas {
		sum += m
	}
	return sum
}

func sumValue(n *Node) int {
	if n.numChildren == 0 {
		return sumMetadata(n)
	}
	var sum int
	for _, m := range n.metadatas {
		if m <= n.numChildren && m > 0 {
			sum += sumValue(n.children[m-1])
		}
	}

	return sum
}
