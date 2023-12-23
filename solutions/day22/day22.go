package day22

import (
	"bufio"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Axis rune

const (
	XAxis Axis = 'x'
	YAxis Axis = 'y'
)

type Range [2]int

func (r Range) Start() int {
	return r[0]
}

func (r Range) End() int {
	return r[1]
}

func (r Range) Length() int {
	return r.End() - r.Start() + 1
}

func (r Range) Contains(n int) bool {
	return n >= r.Start() && n <= r.End()
}

func (r Range) Intersects(other Range) bool {
	var a, b Range
	if r.Length() >= other.Length() {
		a, b = r, other
	} else {
		a, b = other, r
	}

	return a.Contains(b.Start()) || a.Contains(b.End())
}

type Brick [6]int

func (b Brick) String() string {
	return fmt.Sprintf("%d,%d,%d~%d,%d,%d", b[0], b[1], b[2], b[3], b[4], b[5])
}

func (b Brick) Bottom() int {
	return b[2]
}

func (b Brick) Top() int {
	return b[5]
}

func (b Brick) Cubes(axis Axis) Range {
	idx := axis - 'x'
	return Range{b[idx], b[idx+3]}
}

func (b Brick) Supports(other Brick) bool {
	if b.Top()+1 != other.Bottom() {
		return false
	}

	for _, axis := range []Axis{XAxis, YAxis} {
		if !b.Cubes(axis).Intersects(other.Cubes(axis)) {
			return false
		}
	}
	return true
}

func scanBricks(s *bufio.Scanner) []Brick {
	bricks := []Brick{}
	for s.Scan() {
		line := s.Text()
		brick := Brick{}
		for i, end := range strings.Split(line, "~") {
			for j, coordinate := range strings.Split(end, ",") {
				num, _ := strconv.Atoi(coordinate)
				brick[i*3+j] = num
			}
		}
		bricks = append(bricks, brick)
	}
	return bricks
}

func PartA(s *bufio.Scanner) string {
	bricks := scanBricks(s)
	slices.SortFunc(bricks, func(a, b Brick) int {
		return cmp.Compare(a.Bottom(), b.Bottom())
	})

	// Let the bricks fall
	for i := 1; i < len(bricks); i++ {
		current := bricks[i]

	OUTER:
		for current.Bottom() > 1 {
			for j := i - 1; j >= 0; j-- {
				b := bricks[j]
				// fmt.Printf("does %s support %s? %v\n", b, current, b.Supports(current))
				if b.Supports(current) {
					break OUTER
				}
			}
			current[2]--
			current[5]--
		}
		bricks[i] = current
	}

	supports := make(map[Brick][]Brick)
	supportedBy := make(map[Brick][]Brick)

	for i, b := range bricks {
		supports[b] = []Brick{}
		for j := i + 1; j < len(bricks); j++ {
			above := bricks[j]
			if b.Supports(above) {
				supports[b] = append(supports[b], above)
				supportedBy[above] = append(supportedBy[above], b)
			}
		}
	}

	n := 0

	for _, others := range supports {
		canDisintegrate := true

		for _, o := range others {
			if len(supportedBy[o]) < 2 {
				canDisintegrate = false
				break
			}
		}

		if canDisintegrate {
			n++
		}
	}

	return strconv.Itoa(n)
}

func PartB(s *bufio.Scanner) string {
	bricks := scanBricks(s)
	slices.SortFunc(bricks, func(a, b Brick) int {
		return cmp.Compare(a.Bottom(), b.Bottom())
	})

	// Let the bricks fall
	for i := 1; i < len(bricks); i++ {
		current := bricks[i]

	OUTER:
		for current.Bottom() > 1 {
			for j := i - 1; j >= 0; j-- {
				b := bricks[j]
				// fmt.Printf("does %s support %s? %v\n", b, current, b.Supports(current))
				if b.Supports(current) {
					break OUTER
				}
			}
			current[2]--
			current[5]--
		}
		bricks[i] = current
	}

	supports := make(map[Brick][]Brick)
	supportedBy := make(map[Brick][]Brick)

	for i, b := range bricks {
		supports[b] = []Brick{}
		for j := i + 1; j < len(bricks); j++ {
			above := bricks[j]
			if b.Supports(above) {
				supports[b] = append(supports[b], above)
				supportedBy[above] = append(supportedBy[above], b)
			}
		}
	}

	n := 0

	for _, b := range bricks {
		fallen := make(map[Brick]bool)
		q := []Brick{b}
		for len(q) > 0 {
			current := q[0]
			q = q[1:]
			for _, o := range supports[current] {
				willFall := true
				for _, s := range supportedBy[o] {
					if !fallen[s] && s != b {
						willFall = false
						break
					}
				}
				if willFall {
					fallen[o] = true
					q = append(q, o)
				}
			}
		}

		n += len(fallen)
	}

	return strconv.Itoa(n)
}
