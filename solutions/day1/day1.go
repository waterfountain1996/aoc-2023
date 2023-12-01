package day1

import (
	"bufio"
	"strconv"
	"strings"
)

const Digits = "123456789"

var LiteralDigits = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func findDigits(s string, includeLiterals bool) int {
	indexes := [2]int{-1, -1}
	nums := [2]int{}

	for i, digit := range Digits {
		idx := strings.Index(s, string(digit))
		if idx < 0 {
			continue
		}

		if indexes[0] < 0 || idx < indexes[0] {
			indexes[0] = idx
			nums[0] = i + 1
		}

		idx = strings.LastIndex(s, string(digit))

		if indexes[1] < 0 || idx > indexes[1] {
			indexes[1] = idx
			nums[1] = i + 1
		}
	}

	if includeLiterals {
		for i, literal := range LiteralDigits {
			idx := strings.Index(s, literal)
			if idx < 0 {
				continue
			}

			if indexes[0] < 0 || idx < indexes[0] {
				indexes[0] = idx
				nums[0] = i + 1
			}

			idx = strings.LastIndex(s, literal)

			if indexes[1] < 0 || idx > indexes[1] {
				indexes[1] = idx
				nums[1] = i + 1
			}
		}
	}

	return nums[0]*10 + nums[1]
}

func PartA(s *bufio.Scanner) string {
	var calibrationValue int

	for s.Scan() {
		line := s.Text()
		calibrationValue += findDigits(line, false)
	}

	return strconv.Itoa(calibrationValue)
}

func PartB(s *bufio.Scanner) string {
	var calibrationValue int

	for s.Scan() {
		line := s.Text()
		calibrationValue += findDigits(line, true)
	}

	return strconv.Itoa(calibrationValue)
}
