package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var caseDiff = 'a' - 'A'

func main() {
	d, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input", err)
	}

	data := string(d)

	shortestLen := len(data)

	for i := 'a'; i <= 'z'; i++ {
		stripped := strings.Replace(strings.Replace(data, string(i), "", -1), string(i-caseDiff), "", -1)

		reactedLen := len(react(stripped))
		if reactedLen < shortestLen {
			shortestLen = reactedLen
		}
	}

	log.Println(len(react(data)), shortestLen)
}

func react(data string) string {
	var index int
	runes := []rune(data)
	for {
		diff := runes[index] - runes[index+1]
		if (diff == caseDiff) || (diff == -1*caseDiff) {
			if index+2 == len(runes) {
				runes = runes[:index]
				break
			}
			runes = append(runes[:index], runes[index+2:]...)
			index--
			if index < 0 {
				index = 0
			}
			continue
		}

		index++

		if index == len(runes)-1 {
			break
		}
	}

	return string(runes)
}
