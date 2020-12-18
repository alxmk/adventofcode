package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to load input:", err)
	}

	p1 := &programme{
		mem:          make(memory),
		instructions: strings.Split(string(data), "\n"),
	}

	if err := p1.Execute(1); err != nil {
		log.Fatalln("Programme failed to execute:", err)
	}

	log.Println("Part one:", p1.mem.Sum())

	p2 := &programme{
		mem:          make(memory),
		instructions: strings.Split(string(data), "\n"),
	}
	if err := p2.Execute(2); err != nil {
		log.Fatalln("Programme failed to execute:", err)
	}

	log.Println("Part two:", p2.mem.Sum())
}

type programme struct {
	mask         []rune
	mem          memory
	instructions []string
}

func (p *programme) Execute(version int) error {
	for _, line := range p.instructions {
		parts := strings.Split(line, " = ")
		switch parts[0] {
		case "mask":
			p.mask = []rune(parts[1])
			continue
		default:
			v, err := strconv.ParseUint(parts[1], 10, 36)
			if err != nil {
				return fmt.Errorf("failed to parse %s as 36 bit binary uint: %s", parts[1], err)
			}
			var baseAddress uint64
			fmt.Sscanf(parts[0], "mem[%d]", &baseAddress)
			var addresses []uint64
			v, addresses = applyMask(v, baseAddress, p.mask, version)
			for _, address := range addresses {
				p.mem[address] = v
			}
		}
	}
	return nil
}

type memory map[uint64]uint64

func (m memory) Sum() uint64 {
	var sum uint64
	for _, v := range m {
		sum += v
	}
	return sum
}

func applyValueMask(n uint64, mask []rune) uint64 {
	for i, r := range mask {
		switch r {
		case 'X':
			continue
		case '1':
			n |= (1 << uint64(35-i))
		case '0':
			n &= ^(1 << uint64(35-i))
		}
	}
	return n
}

func applyAddressMask(address uint64, mask []rune) []uint64 {
	addressMask := pad([]rune(strconv.FormatUint(address, 2)))
	for i, r := range mask {
		switch r {
		case 'X':
			addressMask[i] = 'X'
		case '1':
			addressMask[i] = '1'
		}
	}
	var addresses []uint64
	for _, perm := range permutations(addressMask) {
		v, err := strconv.ParseUint(string(perm), 2, 36)
		if err != nil {
			log.Fatalf("Failed to parse %s as a binary uint: %s", string(perm), err)
		}
		addresses = append(addresses, v)
	}
	return addresses
}

func pad(addressMask []rune) []rune {
	for len(addressMask) < 36 {
		addressMask = append([]rune{'0'}, addressMask...)
	}
	return addressMask
}

func permutations(addressMask []rune) [][]rune {
	var perms [][]rune
	for i, r := range addressMask {
		if r == 'X' {
			cpy1 := make([]rune, len(addressMask))
			copy(cpy1, addressMask)
			cpy1[i] = '0'
			perms = append(perms, permutations(cpy1)...)
			cpy2 := make([]rune, len(addressMask))
			copy(cpy2, addressMask)
			cpy2[i] = '1'
			perms = append(perms, permutations(cpy2)...)
			return perms
		}
	}
	return [][]rune{addressMask}
}

func applyMask(n, address uint64, mask []rune, version int) (uint64, []uint64) {
	switch version {
	case 1:
		return applyValueMask(n, mask), []uint64{address}
	case 2:
		return n, applyAddressMask(address, mask)
	default:
		panic(fmt.Sprintf("unsupported version %d", version))
	}
}
