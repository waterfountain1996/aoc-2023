package day11

import (
	"bufio"
	"fmt"
	"math"
)

type Point [2]int

func (p *Point) DistanceTo(other Point) int {
	rows := math.Abs(float64(other[0] - (*p)[0]))
	cols := math.Abs(float64(other[1] - (*p)[1]))
	return int(rows) + int(cols)
}

var UndefinedPoint = Point{-1, -1}

func parseInputOptimized(s *bufio.Scanner, fill int) []Point {
	rowIdx := 0

	emptyRows, emptyCols := []bool{}, []bool{}

	galaxies := []Point{}

	for s.Scan() {
		row := s.Text()

		if len(emptyCols) == 0 {
			emptyCols = make([]bool, len(row))
		}

		hasGalaxies := false

		for colIdx, col := range row {
			if col == '#' {
				emptyCols[colIdx] = true
				hasGalaxies = true

				galaxies = append(galaxies, Point{rowIdx, colIdx})
			}
		}

		emptyRows = append(emptyRows, hasGalaxies)

		rowIdx++
	}

	fill -= 1

	rowIncrement := 0

	for rowIdx, rowHasGalaxies := range emptyRows {
		if rowHasGalaxies {
			continue
		}

		for idx, galaxy := range galaxies {
			if galaxy[0] > rowIdx+rowIncrement {
				galaxy[0] += fill
				galaxies[idx] = galaxy
			}
		}

		rowIncrement += fill
	}

	colIncrement := 0

	for colIdx, colHasGalaxies := range emptyCols {
		if colHasGalaxies {
			continue
		}

		for idx, galaxy := range galaxies {
			if galaxy[1] > colIdx+colIncrement {
				galaxy[1] += fill
				galaxies[idx] = galaxy
			}
		}

		colIncrement += fill
	}

	return galaxies
}

func solve(s *bufio.Scanner, fill int) string {
	galaxies := parseInputOptimized(s, fill)

	sum := 0

	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			length := galaxies[i].DistanceTo(galaxies[j])
			sum += length
		}
	}

	return fmt.Sprintf("%d", sum)
}

func PartA(s *bufio.Scanner) string {
	return solve(s, 2)
}

func PartB(s *bufio.Scanner) string {
	return solve(s, 1_000_000)
}
