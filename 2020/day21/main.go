package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/golang-collections/collections/set"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("Failed to read input:", err)
	}
	candidates, ingredients := parse(string(data))
	log.Println("Part one:", findHypoallergenics(candidates, ingredients))
	log.Println("Part two:", findAllergenics(candidates))
}

func parse(input string) (map[string]*set.Set, map[string]int) {
	candidates := make(map[string]*set.Set)
	allIngredients := make(map[string]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " (")
		newSet := set.New()
		for _, ingredient := range strings.Fields(parts[0]) {
			newSet.Insert(ingredient)
			allIngredients[ingredient]++
		}
		for _, allergen := range strings.Split(strings.TrimSuffix(strings.TrimPrefix(parts[1], "contains "), ")"), ", ") {
			if s, ok := candidates[allergen]; ok {
				candidates[allergen] = s.Intersection(newSet)
			} else {
				candidates[allergen] = newSet
			}
		}
	}
	return candidates, allIngredients
}

func findHypoallergenics(candidates map[string]*set.Set, ingredients map[string]int) int {
	inc := make(map[string]int)
	for k, v := range ingredients {
		inc[k] = v
	}
	for _, s := range candidates {
		s.Do(func(i interface{}) {

			delete(inc, i.(string))
		})
	}

	var total int
	for _, n := range inc {
		total += n
	}
	return total
}

func findAllergenics(candidates map[string]*set.Set) string {
	var allergenics, allergens []string
	for len(candidates) > 0 {
		for i, s := range candidates {
			if s.Len() == 1 {
				s.Do(func(v interface{}) {
					allergenics = append(allergenics, v.(string))
					allergens = append(allergens, i)
					for j, t := range candidates {
						if i != j {
							t.Remove(v)
						}
					}
				})
				delete(candidates, i)
			}
		}
	}
	a := as{
		allergenics: allergenics,
		allergens:   allergens,
	}
	sort.Sort(byAllergen(a))
	return strings.Join(a.allergenics, ",")
}

type as struct {
	allergens   []string
	allergenics []string
}

type byAllergen as

func (b byAllergen) Swap(i, j int) {
	b.allergenics[i], b.allergenics[j] = b.allergenics[j], b.allergenics[i]
	b.allergens[i], b.allergens[j] = b.allergens[j], b.allergens[i]
}

func (b byAllergen) Len() int           { return len(b.allergenics) }
func (b byAllergen) Less(i, j int) bool { return b.allergens[i] < b.allergens[j] }
