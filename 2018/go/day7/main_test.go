package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_durationFor(t *testing.T) {
	require.Equal(t, 1, durationFor('A'))
	require.Equal(t, 2, durationFor('B'))
	require.Equal(t, 26, durationFor('Z'))
}
