package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	tests := []struct {
		backpack string
		expect   int
	}{
		{
			backpack: "vJrwpWtwJgWrhcsFMMfFFhFp",
			expect:   16,
		},
		{
			backpack: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			expect:   38,
		},
	}

	for _, tt := range tests {
		t.Run(tt.backpack, func(t *testing.T) {
			assert.Equal(t, tt.expect, score(tt.backpack))
		})
	}
}

// var example = `vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw`

func TestCommon(t *testing.T) {
	tests := []struct {
		backpacks []string
		expect    rune
	}{
		{
			backpacks: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			expect: 'r',
		},
	}

	for _, tt := range tests {
		t.Run(tt.backpacks[0], func(t *testing.T) {
			assert.Equal(t, tt.expect, common(tt.backpacks))
		})
	}
}
