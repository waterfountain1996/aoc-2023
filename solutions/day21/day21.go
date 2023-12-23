package day21

import (
	"bufio"
	"cmp"
	"fmt"
	// "math"
	"slices"
	"strconv"
	"strings"
)

type Point [2]int

func (p Point) Inside(matrix [][]rune) bool {
	return p[0] >= 0 && p[0] < len(matrix) && p[1] >= 0 && p[1] < len(matrix[0])
}

func mod(v, m int) int {
	return (v%m + m) % m
}

func (p Point) Neighbors() []Point {
	return []Point{
		{p[0] - 1, p[1]},
		{p[0] + 1, p[1]},
		{p[0], p[1] - 1},
		{p[0], p[1] + 1},
	}
}

func reachGardenPlots(garden [][]rune, start Point, steps int, expandGarden bool) int {
	q := []Point{start}

	for i := 0; i < steps; i++ {
		newQ := []Point{}

		for _, node := range q {
			for _, n := range node.Neighbors() {
				// fmt.Println(node, n)

				tile := garden[mod(n[0], len(garden))][mod(n[1], len(garden[0]))]
				if (expandGarden || n.Inside(garden)) && tile != '#' {
					newQ = append(newQ, n)
				}
			}
		}

		slices.SortFunc(newQ, func(a, b Point) int {
			if r := cmp.Compare(a[0], b[0]); r != 0 {
				return r
			} else {
				return cmp.Compare(a[1], b[1])
			}
		})
		q = slices.Compact(newQ)
	}

	return len(q)
}

func PartA(s *bufio.Scanner) string {
	matrix := [][]rune{}
	var si, sj int

	for s.Scan() {
		line := s.Text()

		if idx := strings.IndexRune(line, 'S'); idx != -1 {
			si, sj = len(matrix), idx
		}

		matrix = append(matrix, []rune(line))
	}

	result := reachGardenPlots(matrix, Point{si, sj}, 64, false)
	return strconv.Itoa(result)
}

func lagrangeBasis(x float64, xi []float64, i int) float64 {
	result := 1.0
	for j := 0; j < len(xi); j++ {
		if j != i {
			result *= (x - xi[j]) / (xi[i] - xi[j])
		}
	}
	return result
}

func lagrangeInterpolation(x, xi, yi []float64) func(float64) float64 {
	return func(val float64) float64 {
		result := 0.0
		for i := 0; i < len(xi); i++ {
			result += yi[i] * lagrangeBasis(val, xi, i)
		}
		return result
	}
}

func PartB(s *bufio.Scanner) string {
	matrix := [][]rune{}
	var si, sj int

	for s.Scan() {
		line := s.Text()

		if idx := strings.IndexRune(line, 'S'); idx != -1 {
			si, sj = len(matrix), idx
		}

		matrix = append(matrix, []rune(line))
	}

	size := len(matrix)

	xi := []float64{}
	yi := []float64{}

	for i := 0; i < 3; i++ {
		x := 65 + size*i
		r := reachGardenPlots(matrix, Point{si, sj}, x, true)
		xi = append(xi, float64(x))
		yi = append(yi, float64(r))
	}

	polynomial := lagrangeInterpolation(xi, xi, yi)
	return fmt.Sprintf("%.0f", polynomial(26501365))
}
