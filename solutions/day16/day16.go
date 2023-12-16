package day16

import (
	"bufio"
	"cmp"
	"slices"
	"strconv"
)

type Cave []string

func scanCave(s *bufio.Scanner) Cave {
	cave := Cave{}
	for s.Scan() {
		cave = append(cave, s.Text())
	}
	return cave
}

type Direction uint8

const (
	DirRight Direction = iota
	DirDown
	DirLeft
	DirUp
)

type Beam struct {
	I, J int
	Dir  Direction
}

func newBeam(i, j int, dir Direction) Beam {
	return Beam{
		I:   i,
		J:   j,
		Dir: dir,
	}
}

func (b Beam) Move() Beam {
	var i, j int

	switch b.Dir {
	case DirRight:
		i, j = b.I, b.J+1
	case DirLeft:
		i, j = b.I, b.J-1
	case DirDown:
		i, j = b.I+1, b.J
	case DirUp:
		i, j = b.I-1, b.J
	}

	return newBeam(i, j, b.Dir)
}

func (b *Beam) Inside(cave Cave) bool {
	return b.I >= 0 && b.I < len(cave) && b.J >= 0 && b.J < len(cave[0])
}

func (b Beam) Traverse(cave Cave, visited map[Beam]bool) {
	q := []Beam{b}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if !current.Inside(cave) || visited[current] {
			continue
		}

		tile := cave[current.I][current.J]

		// fmt.Printf("Tile %c at %d %d\n", tile, current.I, current.J)

		switch tile {
		case '.':
			q = append(q, current.Move())
		case '|':
			if current.Dir == DirUp || current.Dir == DirDown {
				q = append(q, current.Move())
			} else {
				q = append(
					q,
					newBeam(current.I, current.J, DirUp).Move(),
					newBeam(current.I, current.J, DirDown).Move(),
				)
			}
		case '-':
			if current.Dir == DirLeft || current.Dir == DirRight {
				q = append(q, current.Move())
			} else {
				q = append(
					q,
					newBeam(current.I, current.J, DirLeft).Move(),
					newBeam(current.I, current.J, DirRight).Move(),
				)
			}
		case '/':
			var newDir Direction

			switch current.Dir {
			case DirUp:
				newDir = DirRight
			case DirRight:
				newDir = DirUp
			case DirDown:
				newDir = DirLeft
			case DirLeft:
				newDir = DirDown
			}

			q = append(q, newBeam(current.I, current.J, newDir).Move())
		case '\\':
			var newDir Direction

			switch current.Dir {
			case DirDown:
				newDir = DirRight
			case DirRight:
				newDir = DirDown
			case DirUp:
				newDir = DirLeft
			case DirLeft:
				newDir = DirUp
			}

			q = append(q, newBeam(current.I, current.J, newDir).Move())
		}

		visited[current] = true
	}
}

func countEnergized(visited map[Beam]bool) int {
	tiles := [][2]int{}

	for beam := range visited {
		tiles = append(tiles, [2]int{beam.I, beam.J})
	}

	slices.SortFunc(tiles, func(a, b [2]int) int {
		if r := cmp.Compare(a[0], b[0]); r != 0 {
			return r
		}

		return cmp.Compare(a[1], b[1])
	})

	return len(slices.Compact(tiles))
}

func TraverseAny(cave Cave) int {

	beams := []Beam{}
	for col := range cave[0] {
		beams = append(beams, newBeam(0, col, DirDown))
	}

	for col := range cave[len(cave)-1] {
		beams = append(beams, newBeam(len(cave)-1, col, DirUp))
	}

	for i, row := range cave {
		beams = append(beams, newBeam(i, 0, DirRight))
		beams = append(beams, newBeam(i, len(row)-1, DirLeft))
	}

	r := -1

	for _, b := range beams {
		visited := make(map[Beam]bool)
		b.Traverse(cave, visited)
		n := countEnergized(visited)
		if r == -1 || n > r {
			r = n
		}
	}

	return r
}

func PartA(s *bufio.Scanner) string {
	cave := scanCave(s)
	beam := newBeam(0, 0, DirRight)

	visited := make(map[Beam]bool)
	beam.Traverse(cave, visited)

	return strconv.Itoa(countEnergized(visited))
}

func PartB(s *bufio.Scanner) string {
	cave := scanCave(s)
	result := TraverseAny(cave)
	return strconv.Itoa(result)
}
