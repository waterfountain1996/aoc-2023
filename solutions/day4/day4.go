package day4

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	id int
	winningNumbers, ownNumbers []int
}

func (c *card) NumWinning() int {
	numWinning := 0

	for _, n := range c.ownNumbers {
		for _, winning := range c.winningNumbers {
			if n == winning {
				numWinning += 1
			}
		}
	}

	return numWinning
}

func (c *card) Points() int {
	numWinning := c.NumWinning()

	if numWinning == 0 {
		return 0
	} else if numWinning > 1 {
		numWinning = (1 << (numWinning - 1))
	}

	return numWinning
}

func extractNumbers(s string) []int {
	var numbers []int

	numberRegex := regexp.MustCompile(`(\d+)`)
	for _, match := range numberRegex.FindAllString(s, -1) {
		n, _ := strconv.Atoi(match)
		numbers = append(numbers, n)
	}

	return numbers
}

func cardFromString(s string) *card {
	tokens := strings.SplitN(s, ": ", 2)
	subtokens := strings.SplitN(tokens[1], "|", 2)

	id := extractNumbers(tokens[0])

	winningNumbers := extractNumbers(subtokens[0])
	ownNumbers := extractNumbers(subtokens[1])

	return &card{
		id:             id[0],
		winningNumbers: winningNumbers,
		ownNumbers:     ownNumbers,
	}
}

func PartA(s *bufio.Scanner) string {
	result := 0

	for s.Scan() {
		line := s.Text()
		c := cardFromString(line)
		result += c.Points()
	}
	return strconv.Itoa(result)
}

func PartB(s *bufio.Scanner) string {
	scratchcards := make(map[int]int)
	idx := 1

	for s.Scan() {
		line := s.Text()
		c := cardFromString(line)

		scratchcards[idx] += 1

		for copies := 0; copies < scratchcards[idx]; copies++ {
			for offset := 1; offset <= c.NumWinning(); offset += 1 {
				scratchcards[idx + offset] += 1
			}
		}

		idx += 1
	}

	result := 0
	for _, v := range scratchcards {
		result += v
	}

	return strconv.Itoa(result)
}
