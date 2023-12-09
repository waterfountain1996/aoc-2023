package day9

import (
	"bufio"
	"strconv"
	"strings"
)

func extractNumbers(s string) []int {
	var numbers []int

	for _, token := range strings.Split(s, " ") {
		n, _ := strconv.Atoi(token)
		numbers = append(numbers, n)
	}

	return numbers
}

func parseInput(s *bufio.Scanner) [][]int {
	histories := [][]int{}

	for s.Scan() {
		histories = append(histories, extractNumbers(s.Text()))
	}

	return histories
}

func predict(history []int, backwards bool) int {
	prev := make([]int, len(history))
	copy(prev, history)

	var initial int
	if backwards {
		initial = prev[0]
	} else {
		initial = prev[len(prev) - 1]
	}

	values := []int{initial}

	for {
		current := []int{}
		allZeros := true

		for i := 1; i < len(prev); i++ {
			diff := prev[i] - prev[i - 1]
			if diff != 0 {
				allZeros = false
			}
			current = append(current, diff)
		}

		if backwards {
			values = append(values, current[0])
		} else {
			values = append(values, current[len(current) - 1])
		}

		prev = current

		if allZeros {
			break
		}
	}

	r := 0
	for i := range values {
		i = len(values) - 1 - i

		if backwards {
			r = values[i] - r
		} else {
			r = values[i] + r
		}
	}

	return r
}

func solve(s *bufio.Scanner, backwards bool) int {
	histories := parseInput(s)

	sum := 0

	for _, history := range histories {
		sum += predict(history, backwards)
	}

	return sum
}

func PartA(s *bufio.Scanner) string {
	sum := solve(s, false)
	return strconv.Itoa(sum)
}

func PartB(s *bufio.Scanner) string {
	sum := solve(s, true)
	return strconv.Itoa(sum)
}
