package day5

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

type Range struct {
	Start, Length int
}

func (r *Range) End() int {
	return r.Start + r.Length
}

func (r *Range) String() string {
	return fmt.Sprintf("[%d; %d)", r.Start, r.End())
}

func (r *Range) Contains(value int) bool {
	return r.Start <= value && value < r.End()
}

type RangeMap struct {
	Src, Dst Range
}

func (rm *RangeMap) Map(value int) (int, bool) {
	if rm.Src.Contains(value) {
		return rm.Dst.Start + (value - rm.Src.Start), true
	}

	return value, false
}

func (rm *RangeMap) RevMap(value int) (int, bool) {
	if rm.Dst.Contains(value) {
		return rm.Src.Start + (value - rm.Dst.Start), true
	}

	return value, false
}

type Section struct {
	Src, Dst string
	Maps     []RangeMap
}

func (s *Section) Map(value int) int {
	for _, rngMap := range s.Maps {
		target, ok := rngMap.Map(value)
		if ok {
			return target
		}
	}

	return value
}

func (s *Section) RevMap(value int) int {
	for _, rngMap := range s.Maps {
		target, ok := rngMap.RevMap(value)
		if ok {
			return target
		}
	}

	return value
}

func newSectionFromScanner(s *bufio.Scanner) *Section {
	s.Scan()
	line := s.Text()

	if line == "" {
		return nil
	}

	categoryRegex := regexp.MustCompile(`(\w+)-to-(\w+)`)
	matches := categoryRegex.FindAllStringSubmatch(line, 2)[0]

	src, dst := matches[1], matches[2]

	maps := []RangeMap{}

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		numbers := extractNumbers(line)
		length := numbers[2]
		maps = append(
			maps,
			RangeMap{
				Src: Range{Start: numbers[1], Length: length},
				Dst: Range{Start: numbers[0], Length: length},
			},
		)
	}

	return &Section{
		Src:  src,
		Dst:  dst,
		Maps: maps,
	}
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

func PartA(s *bufio.Scanner) string {
	s.Scan()
	seeds := extractNumbers(s.Text())

	s.Scan()

	sections := make(map[string]*Section)
	for {
		sec := newSectionFromScanner(s)	
		if sec == nil {
			break
		}
		sections[sec.Src] = sec
	}

	minLocation := -1

	for _, seed := range seeds {
		category := "seed"
		value := seed
		for {
			sec := sections[category]
			if sec == nil {
				break
			}
			value = sec.Map(value)
			category = sec.Dst

			if category == "location"{
				if minLocation == -1 || value < minLocation {
					minLocation = value
				}
			}
		}
	}

	return strconv.Itoa(minLocation) 
}

func PartB(s *bufio.Scanner) string {
	s.Scan()

	seeds := extractNumbers(s.Text())
	seedRanges := []Range{}
	for i, seed := range seeds {
		if i % 2 != 0 {
			continue
		}

		seedRanges = append(seedRanges, Range{
			Start:  seed,
			Length: seeds[i+1],
		})
	}

	s.Scan()

	sections := make(map[string]*Section)
	for {
		sec := newSectionFromScanner(s)	
		if sec == nil {
			break
		}
		sections[sec.Dst] = sec
	}

	minLocation := 0

	for {
		category := "location"
		value := minLocation

		for {
			sec := sections[category]

			category = sec.Src
			value = sec.RevMap(value)

			if category == "seed" {
				break
			}
		}

		// fmt.Printf("Value: %d\n", value)
		for _, rng := range seedRanges {
			if rng.Contains(value) {
				return strconv.Itoa(minLocation) 
			}
		}

		minLocation += 1
	}
}
