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

	var numbers []int
	for _, line := range strings.Split(string(data), "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("Failed to convert to int", line, err)
		}
		numbers = append(numbers, i)
	}

	var found2, found3 bool
	for i, n := range numbers {
		for j, m := range numbers[i:] {
			if n+m == 2020 {
				log.Println("Two:", n*m)
				found2 = true
			}
			for _, o := range numbers[i+j:] {
				if n+m+o == 2020 {
					log.Println("Three:", n*m*o)
					found3 = true
					break
				}
			}
		}
		if found2 && found3 {
			break
		}
	}
}
