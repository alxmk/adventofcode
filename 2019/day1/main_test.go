package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fuelRecurse(t *testing.T) {
	tests := []struct {
		name       string
		mass       int
		expectFuel int
	}{
		{
			name:       "Ex1",
			mass:       1969,
			expectFuel: 966,
		},
		{
			name:       "Ex2",
			mass:       100756,
			expectFuel: 50346,
		},
		{
			name:       "Ex3",
			mass:       14,
			expectFuel: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fuelFor(tt.mass)
			log.Println(f)
			r := fuelRecurse(f)
			log.Println(r)
			assert.Equal(t, f+r, tt.expectFuel)
		})
	}
}
