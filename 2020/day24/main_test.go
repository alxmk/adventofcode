package main

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlip(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "Ex1",
			input:  "nwwswee",
			expect: 1,
		},
		{
			name: "Ex2",
			input: `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`,
			expect: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(tt.name)
			g := make(grid)
			for _, line := range strings.Split(tt.input, "\n") {
				g.Flip(line)
			}
			assert.Equal(t, tt.expect, g.Count())
		})
	}
}

func TestIterate(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		iterations int
		expect     int
	}{
		// 		{
		// 			name: "1 iter",
		// 			input: `sesenwnenenewseeswwswswwnenewsewsw
		// neeenesenwnwwswnenewnwwsewnenwseswesw
		// seswneswswsenwwnwse
		// nwnwneseeswswnenewneswwnewseswneseene
		// swweswneswnenwsewnwneneseenw
		// eesenwseswswnenwswnwnwsewwnwsene
		// sewnenenenesenwsewnenwwwse
		// wenwwweseeeweswwwnwwe
		// wsweesenenewnwwnwsenewsenwwsesesenwne
		// neeswseenwwswnwswswnw
		// nenwswwsewswnenenewsenwsenwnesesenew
		// enewnwewneswsewnwswenweswnenwsenwsw
		// sweneswneswneneenwnewenewwneswswnese
		// swwesenesewenwneswnwwneseswwne
		// enesenwswwswneneswsenwnewswseenwsese
		// wnwnesenesenenwwnenwsewesewsesesew
		// nenewswnwewswnenesenwnesewesw
		// eneswnwswnwsenenwnwnwwseeswneewsenese
		// neswnwewnwnwseenwseesewsenwsweewe
		// wseweeenwnesenwwwswnew`,
		// 			iterations: 1,
		// 			expect:     15,
		// 		},
		{
			name: "10 iter",
			input: `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`,
			iterations: 10,
			expect:     37,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := make(grid)
			for _, line := range strings.Split(tt.input, "\n") {
				g.Flip(line)
			}
			for i := 0; i < tt.iterations; i++ {
				g = g.Iterate()
			}
			assert.Equal(t, tt.expect, g.Count())
		})
	}
}
