package knot

import (
	"fmt"
	"math"
)

var (
	standardSuffix = []int{17, 31, 73, 47, 23}
)

func KnotHash(input string) string {
	lengthsASCII := []int{}

	for _, c := range []rune(input) {
		lengthsASCII = append(lengthsASCII, int(c))
	}

	lengthsASCII = append(lengthsASCII, standardSuffix...)

	Loop := NewLoop(256)
	skip := 0
	pos := 0

	for i := 0; i < 64; i++ {
		Loop, skip, pos = OneRound(Loop, skip, pos, lengthsASCII)
	}

	// Will never error in this context
	hash, _ := Loop.DenseHash()

	return hash
}

func OneRound(l Loop, skip, pos int, lengths []int) (Loop, int, int) {
	for _, length := range lengths {
		l.Reverse(pos, length)

		pos += length + skip

		if pos > len(l.Contents) {
			pos = pos - len(l.Contents)
		}

		skip++
	}

	return l, skip, pos
}

type Loop struct {
	Contents map[int]int
}

func NewLoop(length int) Loop {
	l := Loop{
		Contents: make(map[int]int),
	}
	for i := 0; i < length; i++ {
		l.Contents[i] = i
	}

	return l
}

func (l *Loop) Reverse(pos, length int) {
	newContents := make(map[int]int)
	// Copy the old Contents
	for k, v := range l.Contents {
		newContents[k] = v
	}
	for i := 0; i < length; i++ {
		oldIndex := pos + i
		for oldIndex > len(l.Contents)-1 {
			oldIndex = oldIndex - len(l.Contents)
		}

		newIndex := pos + length - i - 1
		for newIndex > len(l.Contents)-1 {
			newIndex = newIndex - len(l.Contents)
		}

		newContents[newIndex] = l.Contents[oldIndex]
	}

	l.Contents = newContents
}

func (l *Loop) Print() string {
	out := ""

	for i := 0; i < len(l.Contents); i++ {
		out += fmt.Sprintf("%d,", l.Contents[i])
	}

	return out
}

func (l *Loop) DenseHash() (string, error) {
	if math.Mod(float64(len(l.Contents)), 16) != 0 {
		return "", fmt.Errorf("Cannot compute the dense hash of a Loop which is not a factor of 16 in length")
	}

	var denseHash string

	for i := 0; i*16 < len(l.Contents); i++ {
		index := i * 16
		xored := l.Contents[index] ^ l.Contents[index+1] ^ l.Contents[index+2] ^ l.Contents[index+3] ^ l.Contents[index+4] ^ l.Contents[index+5] ^ l.Contents[index+6] ^
			l.Contents[index+7] ^ l.Contents[index+8] ^ l.Contents[index+9] ^ l.Contents[index+10] ^ l.Contents[index+11] ^ l.Contents[index+12] ^ l.Contents[index+13] ^
			l.Contents[index+14] ^ l.Contents[index+15]

		hex := fmt.Sprintf("%x", xored)
		if len(hex) < 2 {
			hex = "0" + hex
		}

		denseHash += hex
	}

	return denseHash, nil
}
