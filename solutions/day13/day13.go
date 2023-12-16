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

// This should return at most 1 element for each pattern
// unless you fix a smudge
func findHorizontalReflections(pat Pattern) []int {
	potentialReflections := []int{}

	for i := 0; i < len(pat) - 1; i++ {
		curr, next := pat[i], pat[i+1]
		if curr == next {
			potentialReflections = append(potentialReflections, i)
		}
	}

	reflections := []int{}

	if len(potentialReflections) == 0 {
		return reflections
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
			reflections = append(reflections, start + 1)
		}
	}

	return reflections
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
		if r := findHorizontalReflections(transposePattern(pat)); len(r) > 0 {
			sum += r[0]
		}

		if r := findHorizontalReflections(pat); len(r) > 0 {
			sum += 100 * r[0]
		}
	}

	return strconv.Itoa(sum)
}

func findMismatches(a, b string) []int {
	indexes := []int{}
	for i := range a {
		if a[i] != b[i] {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func findPotentialSmudges(pat Pattern) [][]int {
	smudges := [][]int{}

	for i := range pat {
		for j := i + 1; j < len(pat); j++ {
			a, b := pat[i], pat[j]
			indexes := findMismatches(a, b)

			if len(indexes) == 1 {
				smudges = append(smudges, []int{i, indexes[0]})
			}
		}
	}

	return smudges
}

func swapRune(s string, idx int) string {
	cursor := -1
	return strings.Map(func (c rune) rune {
		cursor++

		if cursor == idx {
			if c == '#' {
				return '.'
			} else {
				return '#'
			}
		}

		return c
	}, s)
}

func PartB(s *bufio.Scanner) string {
	patterns := scanPatterns(s)

	sum := 0

	for _, pat := range patterns {
		smudges := findPotentialSmudges(pat)

		transposed := transposePattern(pat)
		for _, smudge := range findPotentialSmudges(transposed) {
			i := len(pat) - smudge[1] - 1
			j := smudge[0]
			smudges = append(smudges, []int{i, j})
		}

		reflections := []int{}
		if r := findHorizontalReflections(transposed); len(r) > 0 {
			reflections = append(reflections, r[0])
		} else {
			reflections = append(reflections, 0)
		}

		if r := findHorizontalReflections(pat); len(r) > 0 {
			reflections = append(reflections, r[0])
		} else {
			reflections = append(reflections, 0)
		}

		for _, smudge := range smudges {
			i, j := smudge[0], smudge[1]

			newPattern := make(Pattern, len(pat))
			copy(newPattern, pat)
			newPattern[i] = swapRune(newPattern[i], j)

			for i := 0; i < 2; i++ {
				p := newPattern

				if i == 0 {
					p = transposePattern(p)
				}

				for _, r := range findHorizontalReflections(p) {
					if r != reflections[i] && reflections[i] != -1 {
						if i == 0 {
							sum += r
						} else {
							sum += 100 * r
						}

						reflections[i] = -1
					}
				}
			}
		}
	}

	return strconv.Itoa(sum)
}
