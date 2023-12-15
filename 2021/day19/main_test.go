package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlaps(t *testing.T) {
	tests := []struct {
		name   string
		a, b   string
		expect bool
	}{
		{
			name: "Ex1",
			a: `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401`,
			b: `--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390`,
			expect: true,
		},
		{
			name: "Ex1 reduced",
			a: `--
-618,-824,-621
-537,-823,-458
-447,-329,318
404,-588,-901
544,-627,-890
528,-643,409
-661,-816,-575
390,-675,-793
423,-701,434
-345,-311,381
459,-707,401
-485,-357,347`,
			b: `--
686,422,578
605,423,415
515,917,-361
-336,658,858
-476,619,847
-460,603,-452
729,430,532
-322,571,750
-355,545,-477
413,935,-424
-391,539,-444
553,889,-390`,
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := parseScanner(tt.a)
			a.oriented = a.readings
			_, actual := parseScanner(tt.b).Overlaps(a)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		name   string
		a, b   xyz
		expect xyz
	}{
		{
			name:   "Ex1",
			a:      xyz{-618, -824, -621},
			b:      xyz{-618, -824, -621}.offset(xyz{68, -1246, -43}),
			expect: xyz{686, 422, 578},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.b.Transformations(), tt.expect)
		})
	}

	// Relative to scanner 0 in scanner 0 orientation
	s0p0 := xyz{-618, -824, -621}
	// Relative to scanner 1 in scanner 1 orientation
	s1p1 := xyz{686, 422, 578}
	// Relative to scanner 1 in scanner 0 orientation
	s0p1 := s0p0.offset(xyz{68, -1246, -43})
	assert.Contains(t, s0p1.Transformations(), s1p1)

	// Offset from s0 axes to s1 axes
	offset := s0p0.offset(s0p1)
	assert.Equal(t, offset, xyz{68, -1246, -43})
	assert.Equal(t, s0p1.translate(offset), s0p0)

	// Translate s1 point to s0 axes
	reverse := s0p1.offset(s0p0)
	assert.Equal(t, s0p0.translate(reverse), s0p1)
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		name   string
		a, b   xyz
		expect xyz
	}{
		{
			name:   "Ex1",
			a:      xyz{-618, -824, -621},
			b:      xyz{686, 422, 578}.offset(xyz{-618, -824, -621}),
			expect: xyz{686, 422, 578},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.a.translate(tt.b))
		})
	}
}
