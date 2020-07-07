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
		log.Fatalln("Failed to read input file:", err)
	}

	signal := parseOutputList(string(data))

	basePattern := []int{0, 1, 0, -1}

	for i := 0; i < 100; i++ {
		signal = runPhase(signal, basePattern)
	}

	log.Println("Part one:", signal[:8])

	signal10000 := parseOutputList(strings.Repeat(string(data), 10000))
	offset, err := strconv.Atoi(string(data)[:7])
	if err != nil {
		log.Fatalln("Failed to parse offset:", err)
	}
	signal10000 = signal10000[offset:]

	for i := 0; i < 100; i++ {
		signal10000 = runPhaseOffset(signal10000)
	}

	log.Println("Part two:", signal10000[:8])
}

func parseOutputList(input string) outputList {
	var signal outputList
	for _, r := range input {
		signal = append(signal, int(r-'0')) // booyakasha
	}
	return signal
}

type outputList []int

func (o outputList) String() string {
	var out strings.Builder
	for _, i := range o {
		out.WriteString(fmt.Sprintf("%d", i))
	}
	return out.String()
}

func runPhase(input outputList, basePattern []int) outputList {
	output := make(outputList, len(input))
	copy(output, input)
	patternLen := len(basePattern)
	for i := 0; i < len(input); i++ {
		var value int
		for j, d := range input {
			value += d * basePattern[patternIdx(j, i, patternLen)]
		}
		value = value % 10
		if value < 0 {
			value = value * -1
		}
		output[i] = value
	}
	return output
}

func patternIdx(inIdx, outIdx, patternLen int) int {
	return (inIdx + 1) / (outIdx + 1) % patternLen
}

func runPhaseOffset(input outputList) outputList {
	output := make(outputList, len(input))
	var sum int
	for i := len(input) - 1; i >= 0; i-- {
		sum += input[i]
		output[i] = sum % 10
		if output[i] < 0 {
			output[i] = output[i] * -1
		}
	}
	return output
}
