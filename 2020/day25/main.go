package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	lines := strings.Split(string(data), "\n")
	doorPK, err := strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse door public key %s as int64: %s", lines[0], err)
	}

	doorLS := determineLoopSize(doorPK)
	cardPK, err := strconv.ParseInt(lines[1], 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse card public key %s as int64: %s", lines[1], err)
	}

	log.Println("Part one:", getEncryptionKey(cardPK, doorLS))
}

func determineLoopSize(pk int64) int64 {
	var sn int64 = 7
	var v int64 = 1
	for i := int64(1); ; i++ {
		v *= sn
		v %= 20201227
		if v == pk {
			return i
		}
	}
}

func getEncryptionKey(pk int64, ls int64) int64 {
	var v int64 = 1
	for i := int64(0); i < ls; i++ {
		v *= pk
		v %= 20201227
	}
	return v
}
