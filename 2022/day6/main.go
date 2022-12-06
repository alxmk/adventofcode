package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", find(data, 4))
	log.Println("Part two:", find(data, 14))
}

func find(data []byte, size int) int {
	for i := 0; i < len(data)-size; i++ {
		packet := make(map[byte]struct{})
		for j := 0; j < size; j++ {
			packet[data[i+j]] = struct{}{}
		}
		if len(packet) == size {
			return i + size
		}
	}
	return -1
}
