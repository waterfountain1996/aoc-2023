package day3

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(r rune) bool {
	return !(unicode.IsDigit(r) || r == '.')
}

func hasSymbols(s string) bool {
	for _, r := range s {
		if isSymbol(r) {
			return true
		}
	}
	return false
}

func findPartNumbers(prevLine, line, nextLine string) []int {
	if len(prevLine) == 0 {
		prevLine = strings.Repeat(".", len(line))
	}
	if len(nextLine) == 0 {
		nextLine = strings.Repeat(".", len(line))
	}

	var partNumbers []int

	numString := ""
	for idx, ch := range line {
		if unicode.IsDigit(ch) {
			numString += string(ch)
		} else {
			if len(numString) > 0 {
				start := idx - len(numString) - 1
				end := idx + 1

				if start < 0 {
					start = 0
				}
				if end > len(line) {
					end = len(line)
				}

				if (
					hasSymbols(prevLine[start:end]) ||
					hasSymbols(nextLine[start:end]) ||
					hasSymbols(line[start:start+1]) ||
					hasSymbols(line[end-1:end])) {
					num, _ := strconv.Atoi(numString)
					partNumbers = append(partNumbers, num)
				}

				numString = ""
			}
		}
	}

	if len(numString) > 0 {
		start := len(line) - len(numString) - 1

		if start < 0 {
			start = 0
		}
		if (
			hasSymbols(prevLine[start:]) ||
			hasSymbols(nextLine[start:]) ||
			hasSymbols(line[start:start+1]) ||
			hasSymbols(line[len(line)-1:])) {
			num, _ := strconv.Atoi(numString)
			partNumbers = append(partNumbers, num)
		}
	}

	return partNumbers
}

func findGearRatios(prevLine, line, nextLine string) []int {
	if len(prevLine) == 0 {
		prevLine = strings.Repeat(".", len(line))
	}
	if len(nextLine) == 0 {
		nextLine = strings.Repeat(".", len(line))
	}

	var gearRatios []int

	for idx, ch := range line {
		if ch == '*' {
			garbageRegex := regexp.MustCompile(`[^0-9.]`)

			tokens := []string{
				strings.Repeat(".", idx),
				string(ch),
				strings.Repeat(".", len(line) - idx - 1),
			}
			singleGearLine := strings.Join(tokens, "")

			var partNumbers []int
			partNumbers = append(partNumbers, findPartNumbers(
				"",
				garbageRegex.ReplaceAllString(prevLine, "."),
				singleGearLine,
			)...)
			partNumbers = append(partNumbers, findPartNumbers(
				singleGearLine,
				garbageRegex.ReplaceAllString(nextLine, "."),
				"",
			)...)
			partNumbers = append(partNumbers, findPartNumbers(
				"",
				garbageRegex.ReplaceAllString(line, "."),
				singleGearLine,
			)...)

			if len(partNumbers) != 2 {
				continue
			}

			gearRatios = append(gearRatios, partNumbers[0] * partNumbers[1])
		}
	}

	return gearRatios
}

func PartA(s *bufio.Scanner) string {
	s.Scan()
	prevLine := s.Text()

	s.Scan()
	line := s.Text()

	var nextLine string

	result := 0
	for _, partNum := range findPartNumbers("", prevLine, line) {
		result += partNum
	}

	for s.Scan() {
		nextLine = s.Text()

		for _, partNum := range findPartNumbers(prevLine, line, nextLine) {
			result += partNum
		}

		prevLine = line
		line = nextLine
	}

	for _, partNum := range findPartNumbers(prevLine, line, "") {
		result += partNum
	}

	return strconv.Itoa(result)
}

func PartB(s *bufio.Scanner) string {
	s.Scan()
	prevLine := s.Text()

	s.Scan()
	line := s.Text()

	var nextLine string

	result := 0
	for _, gearRatio := range findGearRatios(prevLine, line, "") {
		result += gearRatio
	}

	for s.Scan() {
		nextLine = s.Text()

		for _, gearRatio := range findGearRatios(prevLine, line, nextLine) {
			result += gearRatio
		}

		prevLine = line
		line = nextLine
	}

	for _, gearRatio := range findGearRatios(prevLine, line, "") {
		result += gearRatio
	}

	return strconv.Itoa(result)
}
