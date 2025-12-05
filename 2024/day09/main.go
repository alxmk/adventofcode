package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part one:", partOne(input))
	fmt.Println("Part two:", partTwo(input))
}

func partOne(input string) int64 {
	var compressed files
	i, j := 0, len(input)-1
	for {
		if i >= len(input) || j < i {
			break
		}
		// We're a file
		if i%2 == 0 {
			if size := int(input[i] - '0'); size != 0 {
				compressed = append(compressed, file{i / 2, size})
			}
			i++
			continue
		}
		// Otherwise we're free space, take from the end
		if gap := int(input[i] - '0'); gap != 0 {
			for {
				lastFile := int(input[j] - '0')
				if lastFile != 0 {
					compressed = append(compressed, file{j / 2, min(gap, lastFile)})
				}
				if lastFile >= gap {
					input = input[:j] + string(input[j]-byte(gap))
					break
				}
				gap -= lastFile
				j -= 2
				input = input[:j+1]
				if j <= i {
					break
				}
			}
		}
		i++
	}
	// fmt.Println(compressed.String())
	return compressed.checksum()
}

type file [2]int //id,length

type files []file

func (fs files) String() string {
	var b strings.Builder
	for _, f := range fs {
		for i := 0; i < f[1]; i++ {
			b.WriteRune('0' + rune(f[0]))
		}
	}
	return b.String()
}

func (fs files) checksum() int64 {
	var sum int64
	var b int
	for _, f := range fs {
		for i := b; i < b+f[1]; i++ {
			sum += int64(i * f[0])
		}
		b += f[1]
	}
	return sum
}

func partTwo(input string) int64 {
	var idx int
	var files [][3]int
	var free [][2]int
	for i, r := range input {
		length := int(r - '0')
		if i%2 == 1 {
			free = append(free, [2]int{idx, length})
		} else {
			files = append(files, [3]int{idx, length, i / 2})
		}
		idx += length
	}
	var compressed [][3]int
outer:
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		for j := 0; j < len(free); j++ {
			dest := free[j]
			// If we've gone past the index of the file itself then
			// it can't be compressed
			if dest[0] > file[0] {
				break
			}
			// If it's smaller than the file then we can't move it here
			if dest[1] < file[1] {
				continue
			}
			// Add the compressed file at the index of the destination space
			compressed = append(compressed, [3]int{dest[0], file[1], file[2]})
			// Update the destination space to remove the newly allocated block
			free[j] = [2]int{dest[0] + file[1], dest[1] - file[1]}
			// Consolidate the space where we removed the file
			for k := 0; k < len(free); k++ {
				elem := free[k]
				// If this isn't the space that finishes where our compressed
				// file starts then keep looking
				if file[0] != elem[0]+elem[1] {
					continue
				}
				if k == len(free)-1 {
					// If it's the last element in the free space then just expand it
					free[k] = [2]int{elem[0], elem[1] + file[1]}
				} else {
					// otherwise add it between the two either side
					free[k] = [2]int{elem[0], elem[1] + file[1] + free[k+1][1]}
					free = append(free[:k+1], free[k+2:]...)
				}
				break
			}
			continue outer
		}
		// Can't compress this one, so just add it as is
		compressed = append(compressed, file)
	}
	var checksum int64
	for _, entry := range compressed {
		for i := 0; i < entry[1]; i++ {
			checksum += int64((i + entry[0]) * entry[2])
		}
	}
	return checksum
}
