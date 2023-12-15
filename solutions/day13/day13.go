package day13

import (
	"bufio"
	"strconv"
	"strings"
)

type Pattern []string

func scanPatterns(s *bufio.Scanner) []Pattern {
	patterns := []Pattern{}
	current  := Pattern{}

	for s.Scan() {
		line := s.Text()

		if line == "" {
			patterns = append(patterns, current)
			current = Pattern{}
			continue
		}

		current = append(current, line)
	}

	return append(patterns, current)
}

func findHorizontalReflection(pat Pattern) int {
	potentialReflections := []int{}

	for i := 0; i < len(pat) - 1; i++ {
		curr, next := pat[i], pat[i+1]
		if curr == next {
			potentialReflections = append(potentialReflections, i)
		}
	}

	if len(potentialReflections) == 0 {
		return 0
	}

	for _, start := range potentialReflections {
		ok := true
		i, j := start, start + 1

		for i >= 0 && j < len(pat) {

			if pat[i] != pat[j] {
				ok = false
				break
			}

			i--
			j++
		}

		if ok {
			return start + 1
		}
	}

	return 0
}

func transposePattern(pat Pattern) Pattern {
	transposed := Pattern{}
	rowLength  := len(pat[0])

	for i := 0; i < rowLength; i++ {
		var b strings.Builder
		for j := len(pat) - 1; j >= 0; j-- {
			row := pat[j]
			b.WriteByte(row[i])
		}
		transposed = append(transposed, b.String())
	}

	return transposed
}

func PartA(s *bufio.Scanner) string {
	patterns := scanPatterns(s)

	sum := 0
	
	for _, pat := range patterns {
		cols := findHorizontalReflection(transposePattern(pat))
		rows := findHorizontalReflection(pat)

		sum += cols + (100 * rows)
	}

	return strconv.Itoa(sum)
}

func PartB(s *bufio.Scanner) string {
	return ""
}
