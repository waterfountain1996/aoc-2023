package day1_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/waterfountain1996/aoc-2023/solutions/day1"
)

func TestA(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := day1.PartA(scanner)
	if result != "142" {
		t.Errorf("Expected 142, got %s", result)
	}
}

func TestB(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := day1.PartB(scanner)
	if result != "281" {
		t.Errorf("Expected 281, got %s", result)
	}
}
