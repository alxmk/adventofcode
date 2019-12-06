package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	input := string(data)

	var foundOne, foundTwo bool

	for i := 0; ; i++ {
		h := md5.New()
		io.WriteString(h, input+fmt.Sprintf("%d", i))
		hash := fmt.Sprintf("%x", (h.Sum(nil)))
		if !foundOne && strings.HasPrefix(hash, "00000") {
			log.Printf("Part one: %d (%s)", i, hash)
			if foundTwo {
				break
			}
			foundOne = true
		}
		if !foundTwo && strings.HasPrefix(hash, "000000") {
			log.Printf("Part two: %d (%s)", i, hash)
			if foundOne {
				break
			}
			foundTwo = true
		}
	}
}
