package day12

import (
	"bufio"
	"fmt"
	// "math/bits"
	"strconv"
	"strings"
)

type SpringRow struct {
	Springs string
	Groups  []int
}

func rowFromLine(line string) SpringRow {
	tokens := strings.Split(line, " ")

	springs := tokens[0]
	groups := []int{}

	for _, token := range strings.Split(tokens[1], ",") {
		n, _ := strconv.Atoi(token)
		groups = append(groups, n)
	}

	return SpringRow{
		Springs: springs,
		Groups:  groups,
	}
}

func (row *SpringRow) String() string {
	strGroups := []string{}
	for _, group := range row.Groups {
		strGroups = append(strGroups, strconv.Itoa(group))
	}

	return fmt.Sprintf("%s %s", row.Springs, strings.Join(strGroups, ","))
}

func (row *SpringRow) UniqueArrangements() int {
	unknown := strings.Count(row.Springs, "?")
	counter := 1 << unknown

	result := 0

	for i := 0; i < counter; i++ {
		format := fmt.Sprintf("%%0%db", unknown)
		bitmask := fmt.Sprintf(format, i)

		springs := strings.Split(row.Springs, "")

		bitIdx := 0

		for idx, ch := range row.Springs {
			if rune(ch) == '?' {
				if rune(bitmask[bitIdx]) == '1' {
					springs[idx] = "."
				} else {
					springs[idx] = "#"
				}

				bitIdx++
			}
		}

		newSprings := strings.Join(springs, "")
		works := true
		groupIdx := 0
		for _, token := range strings.Split(newSprings, ".") {
			if len(token) == 0 {
				continue
			}

			if groupIdx+1 > len(row.Groups) {
				works = false
				break
			}

			if len(token) != row.Groups[groupIdx] {
				works = false
				break
			}

			groupIdx++
		}

		if groupIdx != len(row.Groups) {
			works = false
		}

		if works {
			result++
		}
	}

	return result
}

func PartA(s *bufio.Scanner) string {
	sum := 0

	for s.Scan() {
		row := rowFromLine(s.Text())
		sum += row.UniqueArrangements()
	}

	return strconv.Itoa(sum)
}

func PartB(s *bufio.Scanner) string {
	return ""
}
