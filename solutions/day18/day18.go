package day18

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var digPlanRegex = regexp.MustCompile(`([RLUD]{1})\s(\d+)\ \(([#]{1}[0-9A-Fa-f]{6})\)`)

type DigInstruction struct {
	Dir    rune
	Meters int
	Colour string
}

func NewDigInstruction(dir rune, meters int, colour string) *DigInstruction {
	return &DigInstruction{
		Dir:    dir,
		Meters: meters,
		Colour: colour,
	}
}

func DigInstructionFromString(s string) *DigInstruction {
	tokens := digPlanRegex.FindAllStringSubmatch(s, -1)[0]
	dir := rune(tokens[1][0])
	meters, _ := strconv.Atoi(tokens[2])
	colour := tokens[3]

	return NewDigInstruction(dir, meters, colour)
}

func DigInstructionFromColour(colour string) *DigInstruction {
	meters, _ := strconv.ParseInt(colour[1:6], 16, 32)
	dir, _ := strconv.ParseInt(colour[6:], 16, 32)
	dirs := [4]rune{'R', 'D', 'L', 'U'}
	return NewDigInstruction(dirs[int(dir)], int(meters), colour)
}

func (d *DigInstruction) String() string {
	return fmt.Sprintf("%c %d (%s)", d.Dir, d.Meters, d.Colour)
}

func parseInput(s *bufio.Scanner) []*DigInstruction {
	plan := []*DigInstruction{}
	for s.Scan() {
		plan = append(plan, DigInstructionFromString(s.Text()))
	}
	return plan
}

func shoelace(edges [][2]int) int {
	sum := 0

	for i := range edges {
		a, b := edges[i], edges[(i+1)%len(edges)]
		sum += a[0]*b[1] - a[1]*b[0]
	}

	return int(math.Abs(float64(sum))) / 2
}

func PartA(s *bufio.Scanner) string {
	plan := parseInput(s)

	current := [2]int{0, 0}
	edges := [][2]int{}

	for _, instr := range plan {
		x, y := 0, 0
		switch instr.Dir {
		case 'D':
			y = 1
		case 'U':
			y = -1
		case 'R':
			x = 1
		case 'L':
			x = -1
		}

		for i := 0; i < instr.Meters; i++ {
			current = [2]int{current[0] + x, current[1] + y}
			edges = append(edges, current)
		}
	}

	area := shoelace(edges)
	b := len(edges)
	i := area - b/2 + 1
	return strconv.Itoa(b + i)
}

func PartB(s *bufio.Scanner) string {
	plan := parseInput(s)

	numEdges := 0

	current := [2]int{0, 0}
	edges := [][2]int{}

	for _, instr := range plan {
		instr = DigInstructionFromColour(instr.Colour)

		x, y := 0, 0
		switch instr.Dir {
		case 'D':
			y = 1
		case 'U':
			y = -1
		case 'R':
			x = 1
		case 'L':
			x = -1
		}

		next := [2]int{
			current[0] + instr.Meters*x,
			current[1] + instr.Meters*y,
		}
		edges = append(edges, next)

		numEdges += instr.Meters
		current = next
	}

	area := shoelace(edges)
	b := numEdges
	i := area - b/2 + 1
	return strconv.Itoa(b + i)
}
