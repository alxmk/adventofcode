package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTile_Heuristic(t *testing.T) {
	// Heuristic is just the manhattan distance
	tests := []struct {
		from   tile
		to     tile
		expect int
	}{
		{
			from:   tile{x: 0, y: 0},
			to:     tile{x: 1, y: 1},
			expect: 2,
		},
		{
			from:   tile{x: 4, y: 1},
			to:     tile{x: 4, y: 4},
			expect: 3,
		},
		{
			from:   tile{x: 4, y: 2},
			to:     tile{x: 4, y: 1},
			expect: 1,
		},
		{
			from:   tile{x: 18, y: 6},
			to:     tile{x: 18, y: 6},
			expect: 0,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expect, tt.from.Heuristic(tt.to))
	}
}

func TestTile_Next(t *testing.T) {
	tests := []struct {
		t      tile
		expect []tile
	}{
		{
			t:      tile{x: 2, y: 1},
			expect: []tile{tile{x: 2, y: 0}, tile{x: 1, y: 1}, tile{x: 3, y: 1}, tile{x: 2, y: 2}},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expect, tt.t.Next())
	}
}

func TestGrid_FindPath(t *testing.T) {
	g := &grid{
		tiles: [][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, true, true, true, false},
			[]bool{false, true, true, true, false},
			[]bool{false, true, true, true, false},
			[]bool{false, false, false, false, false},
		},
	}

	from := tile{x: 2, y: 1}
	to := tile{x: 2, y: 3}

	path, next, ok := g.FindPath(from, to)

	fmt.Printf("%#v", path)

	require.True(t, ok)
	require.Equal(t, tile{x: 2, y: 2}, next)
	require.Len(t, path, 2)
}

func TestGrid_Passable(t *testing.T) {
	g := &grid{
		tiles: [][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, true, true, true, false},
			[]bool{false, true, true, true, false},
			[]bool{false, true, true, true, false},
			[]bool{false, false, false, false, false},
		},
	}

	for x := 0; x < 5; x++ {
		v, _ := g.Passable(x, 0)
		require.False(t, v)
		v, _ = g.Passable(x, 4)
		require.False(t, v)
	}

	for y := 1; y < 4; y++ {
		v, _ := g.Passable(0, y)
		require.False(t, v)
		v, _ = g.Passable(1, y)
		require.True(t, v)
		v, _ = g.Passable(2, y)
		require.True(t, v)
		v, _ = g.Passable(3, y)
		require.True(t, v)
		v, _ = g.Passable(4, y)
		require.False(t, v)
	}
}

func Test_PathingIssue(t *testing.T) {
	input := `#######
#GE.#E#
#E#..E#
#G.##.#
#..E#E#
#.....#
#######`

	g := load(strings.Split(input, "\n"), 3)

	path, next, ok := g.FindPath(tile{x: 5, y: 4}, tile{x: 2, y: 3})

	log.Printf("%#v", path)
	require.True(t, ok)
	require.Equal(t, tile{x: 5, y: 5}, next)
}

func Test_PathingIssue2(t *testing.T) {
	input := `#######
#GE.#E#
#E#..E#
#G.##.#
#.E.#E#
#.....#
#######`

	g := load(strings.Split(input, "\n"), 3)

	path, next, ok := g.FindPath(tile{x: 5, y: 4}, tile{x: 1, y: 4})

	log.Printf("%#v", path)
	require.True(t, ok)
	require.Equal(t, tile{x: 5, y: 5}, next)
	require.Len(t, path, 6)
}

func Test_solve(t *testing.T) {
	tests := []struct {
		num         string
		input       string
		expectRound int
		expectHP    int
		expectScore int
	}{
		{
			num: "A",
			input: `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			expectRound: 37,
			expectHP:    982,
			expectScore: 36334,
		},
		{
			num: "B",
			input: `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			expectRound: 46,
			expectHP:    859,
			expectScore: 39514,
		},
		{
			num: "C",
			input: `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
			expectRound: 47,
			expectHP:    590,
			expectScore: 27730,
		},
		{
			num: "D",
			input: `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			expectRound: 35,
			expectHP:    793,
			expectScore: 27755,
		},
		{
			num: "E",
			input: `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			expectRound: 54,
			expectHP:    536,
			expectScore: 28944,
		},
		{
			num: "F",
			input: `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			expectRound: 20,
			expectHP:    937,
			expectScore: 18740,
		},
		{
			num: "G",
			input: `####
##E#
#GG#
####`,
			expectRound: 67,
			expectHP:    200,
			expectScore: 13400,
		},
		{
			num: "H",
			input: `#####
#GG##
#.###
#..E#
#.#G#
#.E##
#####`,
			expectRound: 71,
			expectHP:    197,
			expectScore: 13987,
		},
	}

	for _, tt := range tests {
		actualRound, actualHP, actualScore, _ := solve(tt.input, 3)

		assert.Equal(t, tt.expectRound, actualRound, tt.num)
		assert.Equal(t, tt.expectHP, actualHP, tt.num)
		assert.Equal(t, tt.expectScore, actualScore, tt.num)
	}
}

func TestGrid_NextMove(t *testing.T) {
	input := `#############
#.....GG#####
##..G...#.###
###....GGE..#
###.......#.#
###.........#
###....E....#
###...EGE.#.#
##.E........#
#############`

	g := load(strings.Split(input, "\n"), 3)

	dest, next := g.NextMove(tile{x: 4, y: 2}, g.elves)

	require.Equal(t, &tile{x: 5, y: 2}, next)
	require.Equal(t, &tile{x: 7, y: 5}, dest)
}

func TestGrid_NextMove2(t *testing.T) {
	input := `#######
#E..G.#
#...#.#
#.G.#G#
#######`

	g := load(strings.Split(input, "\n"), 3)

	dest, next := g.NextMove(tile{x: 1, y: 1}, g.goblins)

	require.Equal(t, &tile{x: 2, y: 1}, next)
	require.Equal(t, &tile{x: 3, y: 1}, dest)
}

func TestGrid_Attack(t *testing.T) {
	input := `#############
#.....GG#####
##..G...#.###
###....GGE..#
###.......#.#
###.........#
###....E....#
###...EGE.#.#
##.E........#
#############`

	g := load(strings.Split(input, "\n"), 3)

	var goblin *unit

	for _, gob := range g.goblins {
		if gob.x == 7 && gob.y == 7 {
			goblin = gob
			break
		}
	}

	require.NotNil(t, goblin)

	attacked, unitAttacked, killed := g.Attack(goblin, g.elves)

	require.True(t, attacked)
	require.Equal(t, &unit{x: 7, y: 6, kind: "elf", hp: 197, attack: 3}, unitAttacked)
	require.False(t, killed)
}

func Test_DeathOrdering(t *testing.T) {
	// 	Elves:
	// e 0: 2, 1 (200)
	// e 1: 5, 1 (200)
	// e 2: 1, 2 (2)
	// e 3: 2, 3 (200)
	// e 4: 1, 4 (200)
	// e 5: 5, 4 (200)
	// Goblins:
	// g 0: 1, 1 (5)
	// g 1: 1, 3 (32)
	// Round 33 complete
	//   0123456
	// 0 #######
	// 1 #GE.#E#
	// 2 #E#...#
	// 3 #GE##.#
	// 4 #E..#E#
	// 5 #.....#
	// 6 #######

	input := `#######
#GE.#E#
#E#...#
#GE##.#
#E..#E#
#.....#
#######`

	grid := load(strings.Split(input, "\n"), 3)

	grid.goblins = []*unit{
		&unit{
			x:      1,
			y:      1,
			kind:   "goblin",
			hp:     5,
			attack: 3,
		},
		&unit{
			x:      1,
			y:      3,
			kind:   "goblin",
			hp:     32,
			attack: 3,
		},
	}

	grid.elves = []*unit{
		&unit{
			x:      2,
			y:      1,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      5,
			y:      1,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      1,
			y:      2,
			kind:   "elf",
			hp:     2,
			attack: 3,
		},
		&unit{
			x:      2,
			y:      3,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      1,
			y:      4,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      5,
			y:      4,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
	}

	grid.Tick()

	require.Len(t, grid.goblins, 2)
	require.Len(t, grid.elves, 5)
}

func Test_PostDeathMove(t *testing.T) {
	// 	Elves:
	// e 0: 2, 1 (200)
	// e 1: 5, 1 (200)
	// e 2: 1, 2 (2)
	// e 3: 2, 3 (200)
	// e 4: 1, 4 (200)
	// e 5: 5, 4 (200)
	// Goblins:
	// g 0: 1, 1 (5)
	// g 1: 1, 3 (32)
	// Round 33 complete
	//   0123456
	// 0 #######
	// 1 #GE.#E#
	// 2 #E#...#
	// 3 #GE##.#
	// 4 #E..#E#
	// 5 #.....#
	// 6 #######

	input := `#######
#GE.#E#
#E....#
#GE##.#
#E..#E#
#.....#
#######`

	grid := load(strings.Split(input, "\n"), 3)

	grid.goblins = []*unit{
		&unit{
			x:      1,
			y:      1,
			kind:   "goblin",
			hp:     5,
			attack: 3,
		},
		&unit{
			x:      1,
			y:      3,
			kind:   "goblin",
			hp:     32,
			attack: 3,
		},
	}

	grid.elves = []*unit{
		&unit{
			x:      2,
			y:      1,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      5,
			y:      1,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      1,
			y:      2,
			kind:   "elf",
			hp:     2,
			attack: 3,
		},
		&unit{
			x:      2,
			y:      3,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      1,
			y:      4,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
		&unit{
			x:      5,
			y:      4,
			kind:   "elf",
			hp:     200,
			attack: 3,
		},
	}

	grid.Tick()

	require.Len(t, grid.goblins, 2)
	require.Len(t, grid.elves, 5)

	var found bool
	for _, e := range grid.elves {
		if e.x == 5 && e.y == 2 {
			found = true
		}
	}
	require.True(t, found)
}

func Test_FirstMove(t *testing.T) {
	input := `################################
#####...#.....#####..G...#######
######..##.#..#####...#..#######
####...G.G.....####......#######
####.G..........##.......#######
#####..##.###G..#......G.#######
########.G##....#........#######
########..###............#######
########.#####...G........######
########.G................######
#######........G.....E....######
########GE.................#####
################################`

	g := load(strings.Split(input, "\n"), 3)

	expected := `################################
#####...#.....#####......#######
######..##.#..#####..G#..#######
####....G.G....####......#######
####..G.........##.......#######
#####..##.###.G.#.....G..#######
########G.##....#........#######
########..###............#######
########.#####....G.......######
########.............E....######
#######..G......G.........######
########GE.................#####
################################`

	g2 := load(strings.Split(expected, "\n"), 3)

	g.Tick()

	require.Equal(t, g2.tiles, g.tiles)

	sort.Sort(readingOrder(g.elves))
	sort.Sort(readingOrder(g2.elves))
	sort.Sort(readingOrder(g.goblins))
	sort.Sort(readingOrder(g2.goblins))

	g.Print()

	for i, elf := range g.elves {
		elf2 := g2.elves[i]
		require.Equal(t, elf2.x, elf.x, "%s%s", elf.Print(), elf2.Print())
		require.Equal(t, elf2.y, elf.y, "%s%s", elf.Print(), elf2.Print())
	}

	for i, goblin := range g.goblins {
		goblin2 := g2.goblins[i]
		require.Equal(t, goblin2.x, goblin.x, "%s%s", goblin.Print(), goblin2.Print())
		require.Equal(t, goblin2.y, goblin.y, "%s%s", goblin.Print(), goblin2.Print())
	}
}

func Test_AnotherMoveBug(t *testing.T) {
	input := `################################
#####...#.....#####......#######
######..##.#..#####..G#..#######
####...G.G.....####......#######
####.G..........##.......#######
#####..##.###.G.#......G.#######
########.G##....#........#######
########..###............#######
########.#####...G........######
########.G................######
#######........G..........######
########GEG................#####
################################`

	g := load(strings.Split(input, "\n"), 3)

	dest, next := g.NextMove(tile{x: 7, y: 3}, g.elves)

	require.Equal(t, &tile{x: 9, y: 10}, dest)
	require.Equal(t, &tile{x: 8, y: 3}, next)
}

func Test_OtroMoveBug(t *testing.T) {
	input := `################################
#####...#.....#####..G...#######
######..##.#..#####...#..#######
####...G.G.....####......#######
####.G..........##.......#######
#####..##.###G..#......G.#######
########.G##....#........#######
########..###............#######
########.#####...G........######
########.G................######
#######........G..........######
########GEG................#####
################################`

	g := load(strings.Split(input, "\n"), 3)

	dest, next := g.NextMove(tile{x: 13, y: 5}, g.elves)

	require.Equal(t, &tile{x: 9, y: 10}, dest)
	require.Equal(t, &tile{x: 14, y: 5}, next)
}

func Test_YetAnotherMoveBug(t *testing.T) {
	input := `################################
##...G......#########.E...######
##G.........#########......#####
##...#.G....#########.#...######
##...#.......#######E.##########
####.#........#####...##########
#######............E..##########
####..#...........E#############
##...G#...........##############
##........#.......##############
#####G..###..E..################
################################`

	g := load(strings.Split(input, "\n"), 3)

	_, _, ok := g.FindPath(tile{x: 5, y: 1}, tile{x: 18, y: 6})

	require.True(t, ok)
	// require.Len(t, route, 18)

	dest, next := g.NextMove(tile{x: 5, y: 1}, g.elves)

	require.Equal(t, &tile{x: 13, y: 9}, dest)
	require.Equal(t, &tile{x: 6, y: 1}, next)
}

func Test_FirstMoveFull(t *testing.T) {
	input := `################################
#####...#.....#####..G...#######
######..##.#..#####...#..#######
####...G.G.....####......#######
####.G..........##.......#######
#####..##.###G..#......G.#######
########.G##....#........#######
########..###............#######
########.#####...G........######
########.G................######
#######........G.....E....######
########GE.................#####
########....G.#####.E.....######
##########..G#######....E.######
#########G..#########....#######
########....#########...########
#######G....#########...#..#####
####........#########......#####
####........#########.G....#####
##......G....#######....#.#.####
#..E..........#####..........###
#.#.........G.............E.#..#
####.........########..E.#....E#
##........E##########...######.#
##.##.....###########.##########
#..#...E.#######################
##G.G....#######################
##.....#########################
##.....#########################
##...###########################
####.###########################
################################`

	g := load(strings.Split(input, "\n"), 3)

	expect := `################################
#####...#.....#####......#######
######..##.#..#####..G#..#######
####....G.G....####......#######
####..G.........##.......#######
#####..##.###.G.#.....G..#######
########G.##....#........#######
########..###............#######
########.#####....G.......######
########.............E....######
#######..G......G.........######
########GE..G.......E......#####
########....G.#####.......######
##########...#######...E..######
#########.G.#########....#######
########....#########...########
#######.....#########...#..#####
####...G....#########......#####
####........#########..G...#####
##.E...G.....#######....#.#.####
#.............#####.......E..###
#.#........G...........E....#..#
####......E..########....#...E.#
##.........##########...######.#
##.##.....###########.##########
#..#G.E..#######################
##.G.....#######################
##.....#########################
##.....#########################
##...###########################
####.###########################
################################`

	g2 := load(strings.Split(expect, "\n"), 3)

	g.Tick()

	require.Equal(t, g2.tiles, g.tiles)

	sort.Sort(readingOrder(g.elves))
	sort.Sort(readingOrder(g2.elves))
	sort.Sort(readingOrder(g.goblins))
	sort.Sort(readingOrder(g2.goblins))

	g.Print()

	for i, elf := range g.elves {
		elf2 := g2.elves[i]
		require.Equal(t, elf2.x, elf.x, "%s%s", elf.Print(), elf2.Print())
		require.Equal(t, elf2.y, elf.y, "%s%s", elf.Print(), elf2.Print())
	}

	for i, goblin := range g.goblins {
		goblin2 := g2.goblins[i]
		require.Equal(t, goblin2.x, goblin.x, "%s%s", goblin.Print(), goblin2.Print())
		require.Equal(t, goblin2.y, goblin.y, "%s%s", goblin.Print(), goblin2.Print())
	}
}

func Test_Turn5(t *testing.T) {
	turn4 := `0 ################################
#####...#.....#####......#######
######..##.#..#####...#..#######
####...G...G...####......#######
####........G...##..G....#######
#####..##.###...#........#######
########..##....#........#######
########..###..G.....G...#######
########.#####......GE....######
########G.........G.E.....######
#######..G................######
########GEG................#####
########.G....#####.......######
##########...#######......######
#########G..#########....#######
########....#########...########
#######.....#########..E#..#####
####........#########......#####
####.G......#########....E.#####
##...EG......#######...G#.#.####
#.............#####....E.....###
#.#.......G................E#..#
####......E..########....#.....#
##.........##########...######.#
##.##.....###########.##########
#..#.GE..#######################
##....G..#######################
##.....#########################
##.....#########################
##...###########################
####.###########################
################################`

	grid4 := load(strings.Split(turn4, "\n"), 3)

	turn5 := `0 ################################
#####...#.....#####......#######
######..##.#..#####...#..#######
####....G...G..####......#######
####.........G..##...G...#######
#####..##.###...#........#######
########..##....#........#######
########..###...G....G...#######
########.#####......GE....######
########.G.........GE.....######
#######..G................######
########GEG................#####
########.G....#####.......######
##########...#######......######
#########...#########....#######
########.G..#########...########
#######.....#########...#..#####
####........#########..E...#####
####.G......#########...E..#####
##...EG......#######...G#.#.####
#.............#####....E.....###
#.#.......G...............E.#..#
####......E..########....#.....#
##.........##########...######.#
##.##.....###########.##########
#..#.GE..#######################
##....G..#######################
##.....#########################
##.....#########################
##...###########################
####.###########################
################################`

	grid5 := load(strings.Split(turn5, "\n"), 3)

	grid4.Tick()

	require.Equal(t, grid5.tiles, grid4.tiles)

	sort.Sort(readingOrder(grid4.elves))
	sort.Sort(readingOrder(grid5.elves))
	sort.Sort(readingOrder(grid4.goblins))
	sort.Sort(readingOrder(grid5.goblins))

	grid4.Print()

	for i, elf := range grid4.elves {
		elf2 := grid5.elves[i]
		require.Equal(t, elf2.x, elf.x, "%s%s", elf.Print(), elf2.Print())
		require.Equal(t, elf2.y, elf.y, "%s%s", elf.Print(), elf2.Print())
	}

	for i, goblin := range grid4.goblins {
		goblin2 := grid5.goblins[i]
		require.Equal(t, goblin2.x, goblin.x, "%s%s", goblin.Print(), goblin2.Print())
		require.Equal(t, goblin2.y, goblin.y, "%s%s", goblin.Print(), goblin2.Print())
	}
}

func Test_AnotherPathfindingIssue(t *testing.T) {
	input := `################################
#######.....#########...#..#####
####........#########..E...#####
####.G......#########...E..#####
##...EG......#######...G#.#.####
#.............#####....E.....###
#.#.......G................E#..#
####......E..########....#.....#
################################`

	grid := load(strings.Split(input, "\n"), 3)

	route, _, ok := grid.FindPath(tile{x: 27, y: 6}, tile{x: 22, y: 4})

	fmt.Printf("%#v\n", route)

	require.True(t, ok)
	require.Len(t, route, 7)
}

func Test_BillionthMoveBug(t *testing.T) {
	input := `################################
#####...#.....#####......#######
######..##.#..#####...#..#######
####...........####......#######
####............##.......#######
#####..##.###...#........#######
########..##....#........#######
########..###............#######
########.#####............######
########...........G......######
#######.....G.G...........######
########.....G..G..........#####
########......#####.......######
##########...#######......######
#########...#########....#######
########....#########GEG########
#######.....#########E.E#..#####
####..GGG..G#########......#####
####.....GG.#########......#####
##...........#######....#.#.####
#.........G...#####..........###
#.#.......G......G..........#..#
####.....G...########....#.....#
##.........##########...######.#
##.##.....###########.##########
#..#.....#######################
##.......#######################
##.....#########################
##.....#########################
##...###########################
####.###########################
################################`

	grid := load(strings.Split(input, "\n"), 3)

	route, _, ok := grid.FindPath(tile{x: 12, y: 10}, tile{x: 22, y: 14})

	fmt.Printf("Route: %#v\n", route)

	require.True(t, ok)

	dest, next := grid.NextMove(tile{x: 12, y: 10}, grid.elves)

	require.Equal(t, &tile{x: 22, y: 14}, dest)
	require.Equal(t, &tile{x: 12, y: 9}, next)
}
