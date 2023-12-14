package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(parseNumbers(data)))
	log.Println("Part two:", partTwo(parseNumbers(data)))
}

func partOne(numbers []number) int {
	num := numbers[0].Reduce()
	for _, n := range numbers {
		num = num.Add(n).Reduce()
	}
	return num.Magnitude()
}

func partTwo(numbers []number) int {
	var num int
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			num = max(num, numbers[i].Add(numbers[j]).Reduce().Magnitude())
			num = max(num, numbers[j].Add(numbers[i]).Reduce().Magnitude())
		}
	}
	return num
}

func parseNumbers(input []byte) []number {
	var numbers []number
	for _, line := range bytes.Split(input, []byte{'\n'}) {
		numbers = append(numbers, number(line))
	}
	return numbers
}

type number []byte

func (n number) Add(o number) number {
	return number(fmt.Sprintf("[%s,%s]", string(n), string(o)))
}

func (n number) Reduce() number {
	var ok bool
	for n, ok = n.reduce(); ok; n, ok = n.reduce() {
	}
	return n
}

var (
	regularPair      = regexp.MustCompile("[0-9]+,[0-9]+")
	splittableNumber = regexp.MustCompile("[0-9][0-9]+")
)

func (n number) reduce() (number, bool) {
	pair := regularPair.FindAllIndex(n, -1)
	for _, p := range pair {
		var depth int
		for i := 0; i <= p[0]; i++ {
			switch n[i] {
			case '[':
				depth++
			case ']':
				depth--
			}
		}
		if depth > 4 {
			return n.Explode([]int{p[0], p[1]}), true
		}
	}
	if p := splittableNumber.FindIndex(n); p != nil {
		return n.Split(p), true
	}

	return n, false
}

func (n number) Split(pair []int) number {
	match := n[pair[0]:pair[1]]
	v, _ := strconv.Atoi(string(match))
	first, second := v/2, v/2
	if second*2 < v {
		second++
	}
	n = append(n[:pair[0]], append([]byte(fmt.Sprintf("[%d,%d]", first, second)), n[pair[1]:]...)...)
	return n
}

func (n number) Explode(pair []int) number {
	// ps, pe := bytes.LastIndex(n[:index], []byte{'['}), index+bytes.Index(n[index:], []byte{']'})
	exploding := n[pair[0]:pair[1]]
	// log.Println(string(exploding), string(n))
	pl, _ := strconv.Atoi(strings.Split(string(exploding), ",")[0])
	pr, _ := strconv.Atoi(strings.Split(string(exploding), ",")[1])
	// log.Println(string(pair))
	n = append(n[:pair[0]-1], append(number{'0'}, n[pair[1]+1:]...)...)
	// Search right
	var found bool
	var right number
	ri := pair[0] + 1
	for ; ri < len(n); ri++ {
		if n[ri] >= '0' && n[ri] <= '9' {
			found = true
			right = append(right, n[ri])
			continue
		}
		if found {
			break
		}
	}
	if found {
		b, _ := strconv.Atoi(string(right))
		// log.Println(string(n[:ri-len(right)]), strconv.Itoa(pr+b), string(n[ri:]))
		n = append(n[:ri-len(right)], append([]byte(strconv.Itoa(pr+b)), n[ri:]...)...)
	}
	// Search left
	found = false
	var left number
	li := pair[0] - 2
	for ; li >= 0; li-- {
		if n[li] >= '0' && n[li] <= '9' {
			found = true
			left = append(number{n[li]}, left...)
			continue
		}
		if found {
			break
		}
	}
	if found {
		b, _ := strconv.Atoi(string(left))
		n = append(n[:li+1], append([]byte(strconv.Itoa(pl+b)), n[li+len(left)+1:]...)...)
	}
	return n
}

func (n number) Magnitude() int {
	if len(n) == 1 {
		return int(n[0] - '0')
	}
	if bytes.HasPrefix(n, []byte{'['}) {
		if i := n.PairIdx(); i != -1 {
			return n[1:i-1].Magnitude()*3 + n[i+1:].Magnitude()*2
		}
		return n[1 : len(n)-1].Magnitude()
	}
	i := n.PairIdx()
	return n[:i].Magnitude()*3 + n[i+1:].Magnitude()*2
}

func (n number) PairIdx() int {
	// if !bytes.Contains(n, []byte{','}) {
	// 	return -1
	// }
	var depth int
	for i, r := range n {
		switch r {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			if depth == 0 {
				return i
			}
		}
	}
	// return n[1:len(n)-1].PairIdx()
	return -1
}
