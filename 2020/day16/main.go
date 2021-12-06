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
		log.Fatalln("Failed to read input:", err)
	}

	rules, myTicket, nearby := parse(string(data))

	var er int
	var validTickets []ticket
	for _, t := range nearby {
		ticketER := t.IsPossiblyValid(rules)
		if ticketER == 0 {
			validTickets = append(validTickets, t)
		}
		er += ticketER
	}
	log.Println("Part one:", er)

	orderedRules := determineFields(validTickets, rules)
	result := 1
	for i, r := range orderedRules {
		if strings.HasPrefix(r.name, "departure") {
			result *= myTicket.fields[i]
		}
	}
	log.Println("Part two:", result)
}

func parse(input string) ([]rule, ticket, []ticket) {
	parts := strings.Split(input, "\n\n")
	var rules []rule
	for _, line := range strings.Split(parts[0], "\n") {
		sections := strings.Split(line, ": ")
		r := rule{name: sections[0]}
		for _, f := range strings.Fields(sections[1]) {
			if f == "or" {
				continue
			}
			var vr validRange
			fmt.Sscanf(f, "%d-%d", &vr.min, &vr.max)
			r.params = append(r.params, vr)
		}
		rules = append(rules, r)
	}
	var myTicket ticket
	for _, line := range strings.Split(parts[1], "\n") {
		if !strings.Contains(line, ",") {
			continue
		}
		myTicket = parseTicket(line)
	}
	var nearbyTickets []ticket
	for _, line := range strings.Split(parts[2], "\n") {
		if !strings.Contains(line, ",") {
			continue
		}
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}
	return rules, myTicket, nearbyTickets
}

func parseTicket(line string) ticket {
	var t ticket
	for _, v := range strings.Split(line, ",") {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Failed to parse %s as int: %s", v, err)
		}
		t.fields = append(t.fields, num)
	}
	return t
}

type ticket struct {
	fields []int
}

func (t ticket) String() string {
	return fmt.Sprint(t.fields)
}

func (t ticket) IsPossiblyValid(rules []rule) int {
	var errorRate int
	for _, f := range t.fields {
		var found bool
		for _, r := range rules {
			if r.Match(f) {
				found = true
				break
			}
		}
		if !found {
			errorRate += f
		}
	}
	return errorRate
}

type rule struct {
	name   string
	params params
}

func (r rule) String() string {
	return fmt.Sprintf("%s: %s", r.name, r.params)
}

func (r rule) Match(num int) bool {
	for _, v := range r.params {
		if v.Match(num) {
			return true
		}
	}
	return false
}

type validRange struct {
	min, max int
}

func (v validRange) Match(num int) bool {
	return v.min <= num && v.max >= num
}

func (v validRange) String() string {
	return fmt.Sprintf("%d-%d", v.min, v.max)
}

type params []validRange

func (p params) String() string {
	var strs []string
	for _, v := range p {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, " or ")
}

func determineFields(validTickets []ticket, rules []rule) []rule {
	possibilities := make(map[string]map[int]struct{})
	rulesMap := make(map[string]rule)
	for _, r := range rules {
		possibilities[r.name] = map[int]struct{}{
			0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {},
			10: {}, 11: {}, 12: {}, 13: {}, 14: {}, 15: {}, 16: {}, 17: {}, 18: {}, 19: {},
		}
		rulesMap[r.name] = r
	}

	for _, t := range validTickets {
		for i, f := range t.fields {
			for _, r := range rules {
				if !r.Match(f) {
					delete(possibilities[r.name], i)
				}
			}
		}
	}

	correctRules := make([]rule, len(rules))

	for found := 0; found < len(rules)-1; {
		for r, indexes := range possibilities {
			if len(indexes) == 1 {
				found++
				for n := range indexes {
					correctRules[n] = rulesMap[r]
					for _, ix := range possibilities {
						delete(ix, n)
					}
				}
			}
		}
		log.Println(found)
	}
	return correctRules
}
