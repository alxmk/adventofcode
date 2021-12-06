package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}

	rules, messages := parse(string(data))
	rules.Compile()

	var count int
	for _, m := range messages {
		if rules.compiled["0"].Match(m) {
			count++
		}
	}
	log.Println("Part one:", count)

	newRules, _ := parse(string(data))
	newRules.basic["8"] = &rule{matchers: "42 +"}
	newRules.basic["11"] = &rule{matchers: "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31 "}

	newRules.Compile()

	count = 0
	for _, m := range messages {
		if newRules.compiled["0"].Match(m) {
			count++
		}
	}
	log.Println("Part two:", count)
}

func parse(input string) (*rules, []string) {
	parts := strings.Split(input, "\n\n")
	rules := &rules{basic: make(map[string]*rule), compiled: make(map[string]*rule)}
	for _, line := range strings.Split(parts[0], "\n") {
		ruleParts := strings.Split(line, ": ")
		if strings.Contains(ruleParts[1], "\"") {
			rules.compiled[ruleParts[0]] = &rule{matchers: strings.Trim(ruleParts[1], "\"")}
			continue
		}
		rules.basic[ruleParts[0]] = &rule{matchers: ruleParts[1]}
	}
	return rules, strings.Split(parts[1], "\n")
}

type rule struct {
	matchers string
	regex    *regexp.Regexp
}

func (r rule) IsCompiled() bool {
	return !strings.ContainsAny(r.matchers, "0123456789")
}

func (r rule) String() string {
	return r.matchers
}

func (r *rule) Match(text string) bool {
	if r.regex == nil {
		r.regex = regexp.MustCompile("^" + r.matchers + "$")
	}
	return r.regex.MatchString(text)
}

type rules struct {
	basic    map[string]*rule
	compiled map[string]*rule
}

func (r rules) String() string {
	var b strings.Builder
	b.WriteString("basic:\n")
	for n, rule := range r.basic {
		b.WriteString(fmt.Sprintf("%s: %s\n", n, rule))
	}
	b.WriteString("compiled:\n")
	for n, rule := range r.compiled {
		b.WriteString(fmt.Sprintf("%s: %s\n", n, rule))
	}
	return b.String()
}

func (r rules) Compile() {
	for {
		for i, b := range r.basic {
			for j, c := range r.compiled {
				x := strings.Split(b.matchers, " ")
				for k, e := range x {
					if e != j {
						continue
					}
					if strings.Contains(c.matchers, "|") {
						x[k] = "(" + c.matchers + ")"
					} else {
						x[k] = c.matchers
					}
				}
				b.matchers = strings.Join(x, " ")
			}
			if b.IsCompiled() {
				r.compiled[i] = b
				delete(r.basic, i)
			}
		}

		if len(r.basic) == 0 {
			for _, c := range r.compiled {
				c.matchers = strings.ReplaceAll(c.matchers, " ", "")
			}
			break
		}
	}
}
