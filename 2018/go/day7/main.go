package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file", err)
	}

	reqs := make(requirements)

	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Fields(line)

		reqs[parts[7]] = append(reqs[parts[7]], parts[1])

		// Discover stuff with no prereqs
		if _, ok := reqs[parts[1]]; !ok {
			reqs[parts[1]] = []string{}
		}
	}

	reqs2 := copyMap(reqs)

	log.Printf("%#v", reqs)

	var complete []string
	for {
		log.Println(strings.Join(complete, ""))
		next, ok := reqs.Next(complete)
		if !ok {
			break
		}
		complete = append(complete, next)
		delete(reqs, next)
	}

	log.Println("Part one:", strings.Join(complete, ""))

	var t int
	var complete2 []string
	var workers []*worker
	for i := 0; i < 5; i++ {
		workers = append(workers, &worker{})
	}

	for {
		allDone := true
		// Mark complete tasks at the start of the second
		for _, w := range workers {
			if w.currentTask != "" && w.remainingTime == 0 {
				complete2 = append(complete2, w.currentTask)
			}
		}
		// Now get tasks we can start
		for _, w := range workers {
			if w.Ready() {
				next, ok := reqs2.Next(complete2)
				if ok {
					w.currentTask = next
					delete(reqs2, next)
					w.remainingTime = 60 + durationFor([]rune(next)[0])
				} else {
					w.currentTask = ""
				}
			}
			if w.remainingTime != 0 {
				w.remainingTime--
			}
			if w.currentTask != "" {
				allDone = false
			}
		}
		log.Printf("%d '%s' '%s' '%s' '%s' '%s' '%s'", t, workers[0].currentTask, workers[1].currentTask, workers[2].currentTask, workers[3].currentTask, workers[4].currentTask, strings.Join(complete2, ""))
		if allDone {
			break
		}
		t++
	}

	log.Println("Part two:", t)
}

type requirements map[string][]string

func (r requirements) Next(complete []string) (string, bool) {
	var ready []string

	for step, prereqs := range r {
		stepReady := true
		for _, p := range prereqs {
			var matched bool
			for _, c := range complete {
				if p == c {
					matched = true
					break
				}
			}
			if !matched {
				stepReady = false
			}
		}

		if stepReady {
			ready = append(ready, step)
		}
	}

	sort.Sort(sort.StringSlice(ready))

	if len(ready) == 0 {
		return "", false
	}

	return ready[0], true
}

func durationFor(step rune) int {
	return int(step%'A' + 1)
}

func copyMap(r requirements) requirements {
	out := make(requirements)
	for k, v := range r {
		out[k] = v
	}
	return out
}

type worker struct {
	currentTask   string
	remainingTime int
}

func (w *worker) Ready() bool {
	return w.remainingTime == 0 || w.currentTask == ""
}
