package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", partOne(data))
	log.Println("Part two:", partTwo(data))
}

func partOne(input []byte) int {
	var sum int
	for _, step := range bytes.Split(input, []byte{','}) {
		sum += hash(step)
	}
	return sum
}

func partTwo(input []byte) int {
	boxes := make([]box, 256)
	for _, step := range bytes.Split(input, []byte{','}) {
		func() {
			if step[len(step)-1] == '-' {
				n := hash(step[:len(step)-1])
				for i, c := range boxes[n] {
					if bytes.HasPrefix(c, step[:len(step)-1]) {
						boxes[n] = append(boxes[n][:i], boxes[n][i+1:]...)
						return
					}
				}
				return
			}
			n := hash(step[:len(step)-2])
			for i, c := range boxes[n] {
				if bytes.HasPrefix(c, step[:len(step)-2]) {
					boxes[n] = append(boxes[n][:i], append(box{step}, boxes[n][i+1:]...)...)
					return
				}
			}
			boxes[n] = append(boxes[n], step)
		}()
	}
	var sum int
	for n, b := range boxes {
		for i, s := range b {
			sum += (n + 1) * (i + 1) * int(s[len(s)-1]-'0')
		}
	}
	return sum
}

func hash(step []byte) int {
	var value int
	for _, r := range step {
		value = (value + int(r)) * 17 % 256
	}
	return value
}

type box [][]byte
