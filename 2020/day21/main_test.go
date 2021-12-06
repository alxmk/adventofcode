package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHypoallergenics(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name: "Ex1",
			input: `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`,
			expect: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, i := parse(tt.input)
			assert.Equal(t, tt.expect, findHypoallergenics(c, i))
		})
	}
}

func TestFindAllergenics(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name: "Ex1",
			input: `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`,
			expect: "mxmxvkd,sqjhc,fvjkl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := parse(tt.input)
			assert.Equal(t, tt.expect, findAllergenics(c))
		})
	}
}
