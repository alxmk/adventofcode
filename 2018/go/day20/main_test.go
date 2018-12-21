package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Scores(t *testing.T) {
	tests := []struct {
		input       string
		expectmap   string
		expectscore int
	}{
		{
			input: `^WNE$`,
			expectmap: `#####
#.|.#
#-###
#.|X#
#####
`,
			expectscore: 3,
		},
		{
			input: `^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`,
			expectmap: `###########
#.|.#.|.#.#
#-###-#-#-#
#.|.|.#.#.#
#-#####-#-#
#.#.#X|.#.#
#-#-#####-#
#.#.|.|.|.#
#-###-###-#
#.|.|.#.|.#
###########
`,
			expectscore: 18,
		},
		{
			input: `^ENWWW(NEEE|SSE(EE|N))$`,
			expectmap: `#########
#.|.|.|.#
#-#######
#.|.|.|.#
#-#####-#
#.#.#X|.#
#-#-#####
#.|.|.|.#
#########
`,
			expectscore: 10,
		},
		{
			input: `^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$`,
			expectmap: `#############
#.|.|.|.|.|.#
#-#####-###-#
#.#.|.#.#.#.#
#-#-###-#-#-#
#.#.#.|.#.|.#
#-#-#-#####-#
#.#.#.#X|.#.#
#-#-#-###-#-#
#.|.#.|.#.#.#
###-#-###-#-#
#.|.#.|.|.#.#
#############
`,
			expectscore: 23,
		},
		{
			input: `^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$`,
			expectmap: `###############
#.|.|.|.#.|.|.#
#-###-###-#-#-#
#.|.#.|.|.#.#.#
#-#########-#-#
#.#.|.|.|.|.#.#
#-#-#########-#
#.#.#.|X#.|.#.#
###-#-###-#-#-#
#.|.#.#.|.#.|.#
#-###-#####-###
#.|.#.|.|.#.#.#
#-#-#####-#-#-#
#.#.|.|.|.#.|.#
###############
`,
			expectscore: 31,
		},
	}

	for _, tt := range tests {
		w, score, _ := solve(tt.input)

		require.Equal(t, tt.expectmap, w)
		require.Equal(t, tt.expectscore, score)
	}
}
