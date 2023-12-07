package day6

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func extractNumbers(s string) []int {
	var numbers []int

	numberRegex := regexp.MustCompile(`(\d+)`)
	for _, match := range numberRegex.FindAllString(s, -1) {
		n, _ := strconv.Atoi(match)
		numbers = append(numbers, n)
	}

	return numbers
}

func extractJoined(s string) int {
	numberRegex := regexp.MustCompile(`(\d+)`)
	for _, match := range numberRegex.FindAllString(strings.ReplaceAll(s, " ", ""), 1) {
		n, _ := strconv.Atoi(match)
		return n
	}

	return -1
}

func getWaysToWin(raceDuration, targetDistance int) int {
	numWins := 0

	for msHeld := 1; msHeld < raceDuration; msHeld += 1 {
		remaining := raceDuration - msHeld
		distance := remaining * msHeld
		if distance > targetDistance {
			numWins += 1
		}
	}

	return numWins
}

func PartA(s *bufio.Scanner) string {
	s.Scan()
	times := extractNumbers(s.Text())

	s.Scan()
	distances := extractNumbers(s.Text())

	waysToWin := []int{}

	for idx, raceDuration := range times {
		targetDistance := distances[idx]
		waysToWin = append(waysToWin, getWaysToWin(raceDuration, targetDistance))
	}

	r := 1
	for _, n := range waysToWin {
		r *= n
	}

	return strconv.Itoa(r)
}

func PartB(s *bufio.Scanner) string {
	s.Scan()
	raceDuration := extractJoined(s.Text())

	s.Scan()
	targetDistance := extractJoined(s.Text())

	waysToWin := getWaysToWin(raceDuration, targetDistance)
	return strconv.Itoa(waysToWin)
}
