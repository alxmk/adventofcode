package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	log.Println("Part one:", parseAnyone(string(data)).Length())
	log.Println("Part two:", parseEveryone(string(data)).Length())
}

type form map[rune]int

type forms []form

func (f forms) Length() int {
	var l int
	for _, form := range f {
		l += len(form)
	}
	return l
}

func parseAnyone(input string) forms {
	var forms forms
	for _, section := range strings.Split(input, "\n\n") {
		thisForm := make(form)
		for _, line := range strings.Split(section, "\n") {
			for _, r := range line {
				thisForm[r]++
			}
		}
		forms = append(forms, thisForm)
	}
	return forms
}

func parseEveryone(input string) forms {
	var forms forms
	for _, section := range strings.Split(input, "\n\n") {
		thisForm := make(form)
		lines := strings.Split(section, "\n")
		for i, line := range lines {
			for _, r := range line {
				if _, ok := thisForm[r]; i == 0 || ok {
					thisForm[r]++
				}
			}
		}
		for r, count := range thisForm {
			if count != len(lines) {
				delete(thisForm, r)
			}
		}
		forms = append(forms, thisForm)
	}
	return forms
}
