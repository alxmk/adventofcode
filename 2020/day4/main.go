package main

import (
	"encoding/hex"
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

	passports := parse(string(data))

	var allFields, valid int
	for _, p := range passports {
		if p.AllFields() {
			allFields++
		}
		if p.Valid() {
			valid++
		}
	}
	log.Println("Part one:", allFields)
	log.Println("Part two:", valid)
}

func parse(input string) []passport {
	var passports []passport
	current := make(passport)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			passports = append(passports, current)
			current = make(passport)
			continue
		}
		for _, pairs := range strings.Fields(line) {
			kv := strings.Split(pairs, ":")
			current[kv[0]] = kv[1]
		}
	}
	return passports
}

type passport map[string]string

var validators = map[string]func(string) bool{
	"byr": func(i string) bool {
		return validateNumber(i, 1920, 2002)
	},
	"iyr": func(i string) bool {
		return validateNumber(i, 2010, 2020)
	},
	"eyr": func(i string) bool {
		return validateNumber(i, 2020, 2030)
	},
	"hgt": func(i string) bool {
		switch {
		case strings.HasSuffix(i, "in"):
			return validateNumber(strings.TrimSuffix(i, "in"), 59, 76)
		case strings.HasSuffix(i, "cm"):
			return validateNumber(strings.TrimSuffix(i, "cm"), 150, 193)
		}
		return false
	},
	"hcl": func(i string) bool {
		if strings.HasPrefix(i, "#") {
			_, err := hex.DecodeString(strings.TrimPrefix(i, "#"))
			return err == nil
		}
		return false
	},
	"ecl": func(i string) bool {
		switch i {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		}
		return false
	},
	"pid": func(i string) bool {
		_, err := strconv.Atoi(i)
		return err == nil && len(i) == 9
	},
	// "cid",
}

func validateNumber(i string, min, max int) bool {
	if v, err := strconv.Atoi(i); err == nil {
		return v >= min && v <= max
	}
	return false
}

func (p passport) AllFields() bool {
	for n := range validators {
		if _, ok := p[n]; !ok {
			return false
		}
	}
	return true
}

func (p passport) Valid() bool {
	for n, f := range validators {
		if v, ok := p[n]; !ok || !f(v) {
			return false
		}
	}
	return true
}
