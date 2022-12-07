package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expect        *dir
		expectSize    int
		expectPartOne int
	}{
		{
			name:  "Example",
			input: example,
			expect: &dir{
				name:   "/",
				parent: nil,
				subdirs: map[string]*dir{
					"a": {
						name:   "a",
						parent: nil,
						subdirs: map[string]*dir{
							"e": {
								name:    "e",
								parent:  nil,
								subdirs: map[string]*dir{},
								files: map[string]int{
									"i": 584,
								},
							},
						},
						files: map[string]int{
							"f":     29116,
							"g":     2557,
							"h.lst": 62596,
						},
					},
					"d": {
						name:    "d",
						parent:  nil,
						subdirs: map[string]*dir{},
						files: map[string]int{
							"j":     4060174,
							"d.log": 8033020,
							"d.ext": 5626152,
							"k":     7214296,
						},
					},
				},
				files: map[string]int{
					"b.txt": 14848514,
					"c.dat": 8504156,
				},
			},
			expectSize:    48381165,
			expectPartOne: 95437,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := parse(tt.input)
			assert.Equal(t, tt.expectSize, actual.Size())
			assert.Equal(t, tt.expectPartOne, partOne(actual))
		})
	}
}

var example = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
