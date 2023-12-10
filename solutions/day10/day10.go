package day10

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Sketch []string

func parseInput(s *bufio.Scanner) Sketch {
	sketch := Sketch{}

	for s.Scan() {
		level := strings.Join(strings.Split("..", ""), s.Text())
		sketch = append(sketch, level)
	}

	groundLevel := strings.Repeat(".", len(sketch[0]))
	sketch = append(sketch, groundLevel)
	sketch = slices.Insert(sketch, 0, groundLevel)

	return sketch
}

type Point struct {
	I, J int
}

func canGoTo(src, dst Point, pipe rune) bool {
	switch pipe {
	case 'F':
		return dst.I < src.I || dst.J < src.J
	case '7':
		return dst.I < src.I || dst.J > src.J
	case 'J':
		return dst.I > src.I || dst.J > src.J
	case 'L':
		return dst.I > src.I || dst.J < src.J
	case '-':
		return dst.I == src.I
	case '|':
		return dst.J == src.J
	default:
		return false
	}
}

func dfs(s Sketch, visited map[Point]bool, i, j int, depth int) ([]Point, bool) {
	directions := [4]Point{
		{I: i - 1, J: j},
		{I: i + 1, J: j},
		{I: i, J: j - 1},
		{I: i, J: j + 1},
	}

	src := Point{I: i, J: j}
	visited[src] = true

	for _, dir := range directions {
		dst := rune(s[dir.I][dir.J])

		if dst == '.' {
			continue
		} else if dst == 'S' && len(visited) > 2 {
			return []Point{src}, true
		}

		if visited[dir] {
			continue
		}

		if canGoTo(Point{I: i, J: j}, dir, dst) {
		   path, ok := dfs(s, visited, dir.I, dir.J, depth + 1)
		   if ok {
			   return append([]Point{src}, path...), true
		   }
		}
	}

	return []Point{}, false
}

func distance(si, sj, ei, ej int) int {
	return int(math.Abs(float64(ei - si)) + math.Abs(float64(ej - sj)))
}

func shoelace(nodes []Point) int {
	sum := 0

	for i := range nodes {
		a, b := nodes[i], nodes[(i + 1) % len(nodes)]
		// fmt.Printf("%d x %d - %d x %d\n", a.I, b.J, a.J, b.I)
		sum += a.I * b.J - a.J * b.I
	}

	return int(math.Abs(float64(sum))) / 2
}

func PartA(s *bufio.Scanner) string {
	sketch := parseInput(s)

	var si, sj int

	for i, level := range sketch {
		for j, pipe := range level {
			if pipe == 'S' {
				si, sj = i, j
				break
			}
		}
	}

	visited := make(map[Point]bool)

	maxDistance := 0

	path, _ := dfs(sketch, visited, si, sj, 0)

	for _, node := range path {
		d := distance(node.I, node.J, si, sj); if d > maxDistance {
			maxDistance = d
		}
	}

	return fmt.Sprintf("%d", len(path) / 2)
}

func PartB(s *bufio.Scanner) string {
	sketch := parseInput(s)

	var si, sj int

	for i, level := range sketch {
		for j, pipe := range level {
			if pipe == 'S' {
				si, sj = i, j
				break
			}
		}
	}

	visited := make(map[Point]bool)

	path, _ := dfs(sketch, visited, si, sj, 0)

	area := shoelace(path)

	boundary := len(path) / 2
	if len(path) % 2 != 0 {
		boundary = len(path) / 2 + 1
	}

	inner := area - boundary + 1
	return fmt.Sprintf("%d", inner)
}
