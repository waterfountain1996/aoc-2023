package day15

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

type Step string

func scanSequence(s *bufio.Scanner) []Step {
	s.Scan()

	steps := []Step{}
	for _, token := range strings.Split(s.Text(), ",") {
		steps = append(steps, Step(token))
	}

	return steps
}

func (s Step) Operation() rune {
	idx := strings.IndexAny(string(s), "-=")
	return rune(s[idx])
}

func (s Step) Label() string {
	idx := strings.IndexAny(string(s), "-=")
	return string(s[:idx])
}

func (s Step) FocalLength() int {
	idx := strings.LastIndex(string(s), "=")
	if idx == -1 {
		return -1
	}

	n, _ := strconv.Atoi(string(s[idx+1:]))
	return n
}

func HASH(s string) int {
	value := 0
	for _, ch := range s {
		value += int(ch)
		value *= 17
		value %= 256
	}
	return value
}

type Lens struct {
	Label  string
	Length int
}

func newLensFromStep(s Step) Lens {
	return Lens{
		Label:  s.Label(),
		Length: s.FocalLength(),
	}
}

type Box []Lens

func PartA(s *bufio.Scanner) string {
	sequence := scanSequence(s)

	sum := 0
	for _, step := range sequence {
		sum += HASH(string(step))
	}

	return strconv.Itoa(sum)
}

func PartB(s *bufio.Scanner) string {
	sequence := scanSequence(s)

	boxes := [256]Box{}

	for _, step := range sequence {
		lens := newLensFromStep(step)

		boxIdx := HASH(lens.Label)
		box := boxes[boxIdx]

		if step.Operation() == '-' {
			for idx, oldLens := range box {
				if oldLens.Label == lens.Label {
					boxes[boxIdx] = slices.Delete(box, idx, idx+1)
					break
				}
			}
		} else {
			replaced := false

			for idx, oldLens := range box {
				if oldLens.Label == lens.Label {
					boxes[boxIdx] = slices.Replace(box, idx, idx+1, lens)
					replaced = true
					break
				}
			}

			if !replaced {
				boxes[boxIdx] = append(box, lens)
			}
		}
	}

	focusingPower := 0

	for boxIdx, box := range boxes {
		for lensIdx, lens := range box {
			focusingPower += (1 + boxIdx) * (1 + lensIdx) * lens.Length
		}
	}

	return strconv.Itoa(focusingPower)
}
