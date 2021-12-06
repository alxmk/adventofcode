package main

import (
	"io/ioutil"
	"log"

	"github.com/alxmk/adventofcode/2019/day2/intcode"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	programme, err := intcode.Parse(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input as intcode programme:", err)
	}

	var maxPower int64
	for _, s := range generateSequences([]int64{0, 1, 2, 3, 4}) {
		output := runAmps(s, programme)
		if output > maxPower {
			maxPower = output
		}
	}

	log.Println("Part one:", maxPower)

	maxPower = 0
	for _, s := range generateSequences([]int64{5, 6, 7, 8, 9}) {
		output := runAmpsFeedback(s, programme)
		if output > maxPower {
			maxPower = output
		}
	}

	log.Println("Part two:", maxPower)
}

type seq []int64

func generateSequences(values []int64) []seq {
	resultChan := make(chan seq)
	go func() {
		generateSequence(values, len(values), resultChan)
		close(resultChan)
	}()

	var results []seq
	for result := range resultChan {
		results = append(results, result)
	}
	return results
}

func generateSequence(values []int64, n int, resultChan chan seq) {
	if n == 1 {
		tmp := make([]int64, len(values))
		copy(tmp, values)
		resultChan <- tmp
	} else {
		for i := 0; i < n; i++ {
			generateSequence(values, n-1, resultChan)
			if n%2 == 1 {
				tmp := values[i]
				values[i] = values[n-1]
				values[n-1] = tmp
			} else {
				tmp := values[0]
				values[0] = values[n-1]
				values[n-1] = tmp
			}
		}
	}
}

func runAmps(phases seq, programme intcode.Programme) int64 {
	var input int64
	for _, phase := range phases {
		in, out := make(chan int64), make(chan int64)
		go func() {
			if err := programme.Copy().Run(in, out); err != nil {
				log.Fatalln("Amp run failed:", err)
			}
		}()

		go func(phase int64) {
			in <- phase
			in <- input
			close(in)
		}(phase)

		for o := range out {
			input = o
		}
	}
	return input
}

func runAmpsFeedback(phases seq, programme intcode.Programme) int64 {
	inputA := make(chan int64)
	nextIn := inputA
	var out chan int64
	for _, phase := range phases {
		out = runAmp(programme.Copy(), phase, nextIn)
		// This out is next in
		nextIn = out
	}
	var input int64
	// Initial input
	inputA <- input
	// Last out is output E
	for input = range out {
		inputA <- input
	}
	return input
}

func runAmp(amp intcode.Programme, phase int64, input chan int64) chan int64 {
	in, out := make(chan int64), make(chan int64)
	go func() {
		if err := amp.Run(in, out); err != nil {
			panic("amp run failed:" + err.Error())
		}
	}()

	go func() {
		in <- phase
		for i := range input {
			in <- i
		}
		close(in)
	}()

	return out
}
