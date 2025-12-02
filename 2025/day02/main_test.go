package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeFactors(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, primeFactors(6))
	assert.Equal(t, []int{1, 3}, primeFactors(9))
	assert.Equal(t, []int{1, 11}, primeFactors(11))
}

// func TestIsValid(t *testing.T) {
// 	assert.Equal(t, true, isValid(824824821))
// 	assert.Equal(t, false, isValid(824824824))
// 	assert.Equal(t, true, isValid(1138912))
// 	assert.Equal(t, false, isValid(1111111))
// }
