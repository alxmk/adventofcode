package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	p, err := parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input:", err)
	}

	log.Println("Part one:", sumVersions([]packet{*p}))
	log.Println("Part two:", p.Evaluate())
}

type packet struct {
	version    uint64
	ptype      uint64
	subpackets []packet
	value      uint64
}

func sumVersions(packets []packet) uint64 {
	var sum uint64
	for _, p := range packets {
		sum += p.version
		sum += sumVersions(p.subpackets)
	}
	return sum
}

func parse(input string) (*packet, error) {
	input = parseAsBinary(input)
	p, _, err := parsePacket(input, 0, false)
	if err != nil {
		return nil, fmt.Errorf("failed to parse packet: %s", err)
	}
	return p, nil
}

func parseAsBinary(input string) string {
	var b strings.Builder
	for _, r := range input {
		v, _ := strconv.ParseUint(string(r), 16, 4)
		b.WriteString(fmt.Sprintf("%04b", v))
	}
	return b.String()
}

func parsePacket(input string, i int, subpacket bool) (*packet, int, error) {
	var p packet
	var err error
	if p.version, err = strconv.ParseUint(input[i:i+3], 2, 3); err != nil {
		return nil, i, fmt.Errorf("failed to parse packet version: %s", err)
	}
	i += 3
	if p.ptype, err = strconv.ParseUint(input[i:i+3], 2, 3); err != nil {
		return nil, i, fmt.Errorf("failed to parse packet type: %s", err)
	}
	i += 3
	switch p.ptype {
	case 4:
		// Literal
		for {
			v, err := strconv.ParseUint(input[i:i+5], 2, 5)
			if err != nil {
				return nil, i, fmt.Errorf("failed to parse literal: %s", err)
			}
			i += 5
			p.value = (p.value << 4) + (v & 15)
			if v&(1<<4) == 0 {
				break
			}
		}
	default:
		// Operator
		switch input[i] {
		case '0':
			i++
			// Next 15 bits are subpackets by length
			len, err := strconv.ParseUint(input[i:i+15], 2, 15)
			if err != nil {
				return nil, i, fmt.Errorf("failed to parse length: %s", err)
			}
			i += 15
			end := i + int(len)
			for i < end {
				var q *packet
				if q, i, err = parsePacket(input, i, true); err != nil {
					return nil, i, fmt.Errorf("failed to parse subpacket: %s", err)
				}
				p.subpackets = append(p.subpackets, *q)
			}
		case '1':
			i++
			// Next 11 bits are subpackets by qty
			qty, err := strconv.ParseUint(input[i:i+11], 2, 11)
			if err != nil {
				return nil, i, fmt.Errorf("failed to parse quantity: %s", err)
			}
			i += 11
			for j := 0; j < int(qty); j++ {
				var q *packet
				if q, i, err = parsePacket(input, i, true); err != nil {
					return nil, i, fmt.Errorf("failed to parse subpacket: %s", err)
				}
				p.subpackets = append(p.subpackets, *q)
			}
		}
	}
	return &p, i, nil
}

func (p packet) Evaluate() uint64 {
	switch p.ptype {
	case 4:
		return p.value
	case 0:
		var sum uint64
		for _, s := range p.subpackets {
			sum += s.Evaluate()
		}
		return sum
	case 1:
		product := uint64(1)
		for _, s := range p.subpackets {
			product *= s.Evaluate()
		}
		return product
	case 2:
		min := uint64(math.MaxUint64)
		for _, s := range p.subpackets {
			if v := s.Evaluate(); v < min {
				min = v
			}
		}
		return min
	case 3:
		var max uint64
		for _, s := range p.subpackets {
			if v := s.Evaluate(); v > max {
				max = v
			}
		}
		return max
	case 5:
		if p.subpackets[0].Evaluate() > p.subpackets[1].Evaluate() {
			return 1
		}
		return 0
	case 6:
		if p.subpackets[0].Evaluate() < p.subpackets[1].Evaluate() {
			return 1
		}
		return 0
	case 7:
		if p.subpackets[0].Evaluate() == p.subpackets[1].Evaluate() {
			return 1
		}
		return 0
	}
	return 0
}
