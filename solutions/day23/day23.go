package day23

import (
	"bufio"
	// "fmt"
	"maps"
	"slices"
	"strconv"
)

type Point [2]int

func (p Point) Neighbors() []Point {
	return []Point{
		{p[0] - 1, p[1]},
		{p[0] + 1, p[1]},
		{p[0], p[1] - 1},
		{p[0], p[1] + 1},
	}
}

func (p Point) Inside(matrix [][]rune) bool {
	return p[0] >= 0 && p[0] < len(matrix) && p[1] >= 0 && p[1] < len(matrix[0])
}

func findLongestPath(paths [][]rune, start, end Point, visited map[Point]bool) []Point {
	path := []Point{}
	if start == end {
		return path
	}

	path = append(path, start)

	visited[start] = true

	var neighbors []Point

	switch paths[start[0]][start[1]] {
	case '.':
		neighbors = start.Neighbors()
	case '>':
		neighbors = []Point{{start[0], start[1] + 1}}
	case '<':
		neighbors = []Point{{start[0], start[1] - 1}}
	case 'v':
		neighbors = []Point{{start[0] + 1, start[1]}}
	case '^':
		neighbors = []Point{{start[0] - 1, start[1]}}
	}

	var longest []Point

	for _, n := range neighbors {
		if !n.Inside(paths) || paths[n[0]][n[1]] == '#' || visited[n] {
			continue
		}

		subVisited := maps.Clone(visited)

		if tail := findLongestPath(paths, n, end, subVisited); len(tail) > len(longest) {
			longest = tail
		}
	}

	return append(path, longest...)
}

func longestPath(paths [][]rune, start, end Point) int {
	visited := make(map[Point]bool)
	return len(findLongestPath(paths, start, end, visited))
}

func PartA(s *bufio.Scanner) string {
	paths := [][]rune{}

	for s.Scan() {
		paths = append(paths, []rune(s.Text()))
	}

	start := Point{0, slices.Index(paths[0], '.')}
	end := Point{len(paths) - 1, slices.Index(paths[len(paths)-1], '.')}

	r := longestPath(paths, start, end)
	return strconv.Itoa(r)
}

func PartB(s *bufio.Scanner) string { return "" }
