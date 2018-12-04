package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	lines := strings.Split(string(data), "\n")
	sort.Sort(sort.StringSlice(lines))

	var currentGuard int
	var wasAsleep int
	sleepPattern := make(map[int][]int)

	for _, line := range lines {
		minute, err := strconv.Atoi(strings.TrimSuffix(strings.Split(strings.Fields(line)[1], ":")[1], "]"))
		if err != nil {
			log.Fatalln("Failed to parse minute", line, err)
		}

		switch {
		case strings.Contains(line, "Guard"):
			var err error
			currentGuard, err = strconv.Atoi(strings.TrimPrefix(strings.Fields(line)[3], "#"))
			if err != nil {
				log.Fatalln("Failed to parse guard ID", line, err)
			}
			if _, ok := sleepPattern[currentGuard]; !ok {
				sleepPattern[currentGuard] = make([]int, 60)
			}
			continue
		case strings.Contains(line, "falls asleep"):
			wasAsleep = minute
		case strings.Contains(line, "wakes up"):
			for i := wasAsleep; i < minute; i++ {
				sleepPattern[currentGuard][i]++
			}
			wasAsleep = -1
		}
	}

	// log.Printf("%#v", sleepPattern)

	mostSleepy := getMostSleepy(sleepPattern)

	sleepiestMinute := getSleepiestMinute(sleepPattern[mostSleepy])

	guard, minute := getSameMinuteMostAsleep(sleepPattern)

	log.Println(mostSleepy, sleepiestMinute, mostSleepy*sleepiestMinute, ";", guard, minute, guard*minute)
}

func getMostSleepy(sleepPattern map[int][]int) int {
	mostSleepy := -1
	recordSleepiness := -1
	for guard, sleeps := range sleepPattern {
		var sleepCount int
		for _, sleep := range sleeps {
			sleepCount += sleep
		}
		if sleepCount > recordSleepiness {
			recordSleepiness = sleepCount
			mostSleepy = guard
		}
	}

	return mostSleepy
}

func getSleepiestMinute(sleeps []int) int {
	sleepiest := -1
	record := -1

	for i, s := range sleeps {
		if s > record {
			sleepiest = i
			record = s
		}
	}

	return sleepiest
}

func getSameMinuteMostAsleep(sleepPattern map[int][]int) (int, int) {
	guard := -1
	minute := -1
	record := -1

	for i := 0; i < 60; i++ {
		for g, sleeps := range sleepPattern {
			if sleeps[i] > record {
				guard = g
				record = sleeps[i]
				minute = i
			}
		}
	}

	return guard, minute
}
