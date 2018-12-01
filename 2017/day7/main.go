package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file %v", err)
	}

	processes := strings.Split(string(input), "\n")

	holders := []string{}
	holdees := make(map[string]struct{})

	for _, p := range processes {
		procName, children := parse(p)
		if len(children) != 0 {
			holders = append(holders, procName)
		}

		for _, c := range children {
			holdees[c] = struct{}{}
		}
	}

	for _, h := range holders {
		if _, ok := holdees[h]; !ok {
			log.Println("The answer is", h)
			break
		}
	}
}

func parse(line string) (string, []string) {
	processName := strings.Fields(line)[0]
	children := []string{}

	if parts := strings.Split(line, "->"); len(parts) == 2 {
		children = strings.Fields(strings.Replace(parts[1], ",", "", -1))
	}

	return processName, children
}
