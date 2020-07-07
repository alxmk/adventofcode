package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file:", err)
	}

	rs, err := parseRecipes(string(data))
	if err != nil {
		log.Fatalln("Failed to parse input as recipes:", err)
	}

	rq := rs.getMaterials(chemical{"FUEL", 1}, make(map[string]int64))
	log.Println("Part one:", rq[0].amount)
	f := 1000000000000 / rq[0].amount
	log.Println("Part two:", binarySearch(f*2, f, 1000000000000, func(fuel int64) int64 {
		return rs.getMaterials(chemical{"FUEL", fuel}, make(map[string]int64))[0].amount
	}))
}

type recipe struct {
	ingredients []chemical
	result      chemical
}

type chemical struct {
	name   string
	amount int64
}

func (c chemical) String() string {
	return fmt.Sprintf("%d %s", c.amount, c.name)
}

type recipes map[string]*recipe

func (r recipe) String() string {
	var ingredients []string
	for _, i := range r.ingredients {
		ingredients = append(ingredients, i.String())
	}
	return fmt.Sprintf("%s => %s", strings.Join(ingredients, ", "), r.result)
}

func (r recipes) String() string {
	var out []string
	for _, rec := range r {
		out = append(out, rec.String())
	}
	return strings.Join(out, "\n")
}

func parseRecipes(input string) (recipes, error) {
	rs := make(recipes)
	for _, line := range strings.Split(input, "\n") {
		r, err := parseRecipe(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse recipes: %v", err)
		}
		rs[r.result.name] = r
	}
	return rs, nil
}

func parseRecipe(line string) (*recipe, error) {
	parts := strings.Split(line, " => ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("malformed line: %s", line)
	}
	ingredients, result := parts[0], parts[1]
	var r recipe
	if _, err := fmt.Sscanf(result, "%d %s", &r.result.amount, &r.result.name); err != nil {
		return nil, fmt.Errorf("failed to parse \"%s\": %s", result, err)
	}

	for _, ingredient := range strings.Split(ingredients, ", ") {
		var i chemical
		if _, err := fmt.Sscanf(ingredient, "%d %s", &i.amount, &i.name); err != nil {
			return nil, fmt.Errorf("failed to parse \"%s\": %s", ingredient, err)
		}
		r.ingredients = append(r.ingredients, i)
	}
	return &r, nil
}

func (r recipes) getMaterials(chem chemical, stock map[string]int64) []chemical {
	// Ore we have infinite so just return whatever is needed
	if chem.name == "ORE" {
		return []chemical{chem}
	}

	// Check the stock first
	if s := stock[chem.name]; s > 0 {
		// If we have enough in stock then reduce the stock and return
		if chem.amount < s {
			stock[chem.name] -= chem.amount
			return []chemical{}
		}
		// Otherwise use up the stock and we have to make the rest
		chem.amount -= s
		stock[chem.name] = 0
	}

	// Otherwise we need to run a recipe
	recipe := r[chem.name]

	// How many times do we need to make the recipe to get enough?
	repetitions := getRepetitions(chem.amount, recipe.result.amount)
	var materials []chemical
	for _, ingredient := range recipe.ingredients {
		materials = append(materials, r.getMaterials(chemical{ingredient.name, ingredient.amount * repetitions}, stock)...)
	}

	stock[chem.name] += recipe.result.amount*repetitions - chem.amount

	return dedupe(materials)
}

func getRepetitions(amount, produced int64) int64 {
	if amount%produced == 0 {
		return amount / produced
	}
	return 1 + amount/produced
}

func dedupe(in []chemical) []chemical {
	var deduped []chemical
	dedupeMap := make(map[string]int64)
	for _, c := range in {
		dedupeMap[c.name] += c.amount
	}
	for k, v := range dedupeMap {
		deduped = append(deduped, chemical{k, v})
	}

	return deduped
}

func binarySearch(upper, lower, target int64, f func(int64) int64) int64 {
	resUpper, resLower := f(upper), f(lower)
	if resUpper == target {
		return upper
	}
	if resLower == target {
		return lower
	}

	midpoint := lower + ((upper - lower) / 2)

	resMP := f(midpoint)

	if resMP == target || resMP == resUpper || resMP == resLower {
		return midpoint
	}

	if resMP > target {
		return binarySearch(midpoint, lower, target, f)
	}

	return binarySearch(upper, midpoint, target, f)
}
