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

func (rm *RangeMap) MapRange(r *Range) ([]*Range, bool) {
	ranges := []*Range{}
	ok := true

	src := rm.Src

	if r.Start < src.Start {
		if r.End() <= src.Start {
			fmt.Printf("Range %s is to the left of %s\n", r.String(), src.String())
			ranges = append(ranges, r)
			ok = false
		} else if r.End() <= src.End() {
			fmt.Printf("Range %s (%d) left intersects %s, ", r.String(), r.Length, src.String())

			outer := Range{
				Start:  r.Start,
				Length: src.Start - r.Start,
			}

			inner := Range{
				Start:  rm.Dst.Start,
				Length: r.Length - outer.Length,
			}

			ranges = append(ranges, &outer, &inner)

			fmt.Printf(
				"which remaps to %s (%d) and %s (%d)\n",
				outer.String(), outer.Length,
				inner.String(), inner.Length,
			)

		} else {
			fmt.Printf("Range %s is larger than %s\n", r.String(), src.String())

			leftOuter := Range{
				Start:  r.Start,
				Length: src.Start - r.Start,
			}

			inner := rm.Dst

			rightOuter := Range{
				Start:  src.End(),
				Length: r.Length - src.End(),
			}

			ranges = append(ranges, &leftOuter, &inner, &rightOuter)
		}
	} else if r.Start < src.End() {
		start, _ := rm.Map(r.Start)
		if r.End() > src.End() {
			fmt.Printf("Range %s (%d) right intersects %s, ", r.String(), r.Length, src.String())

			inner := Range{
				Start:  start,
				Length: src.End() - r.Start,
			}

			outer := Range{
				Start:  src.End(),
				Length: r.Length - inner.Length,
			}

			ranges = append(ranges, &outer, &inner)

			fmt.Printf(
				"which remaps to %s (%d) and %s (%d)\n",
				outer.String(), outer.Length,
				inner.String(), inner.Length,
			)

		} else {
			fmt.Printf("Range %s is inside %s\n", r.String(), src.String())
			ranges = append(ranges, &Range{
				Start:  start,
				Length: r.Length,
			})
		}
	} else {
		fmt.Printf("Range %s is to the right of %s\n", r.String(), src.String())
		ranges = append(ranges, r)
		ok = false
	}

	return ranges, ok
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

func (s *Section) MapRange(rng *Range) []*Range {
	allRanges := []*Range{rng}

	for _, rngMap := range s.Maps {
		ranges, ok := rngMap.MapRange(rng)
		if ok {
			allRanges = append(allRanges, ranges...)
		}

		// if ok {
		// 	return ranges
		// }

		// if ok {
		// 	allRanges = append(allRanges, ranges...)
		// } else if !check {
		// 	allRanges = append(allRanges, ranges...)
		// 	check = true
		// }
	}

	return allRanges
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

	ranges := []*Range{}
	for i, seed := range seeds {
		if i % 2 != 0 {
			continue
		}

		ranges = append(ranges, &Range{
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
		sections[sec.Src] = sec
	}

	category := "seed"
	for {
		section := sections[category]
		if section == nil {
			break
		}

		fmt.Printf("Section: %s\n", category)

		newRanges := []*Range{}

		for _, rng := range ranges {
			result := section.MapRange(rng)
			fmt.Printf("  Remap %s to %v\n", rng, result)
			newRanges = append(newRanges, result...)
		}

		category = section.Dst
		ranges = newRanges

		fmt.Println()

	}

	minLocation := -1

	for _, rng := range ranges {
		// fmt.Printf("%d\n", rng.Start)
		if minLocation == -1 || rng.Start < minLocation { //&& rng.Start > 0 {
			minLocation = rng.Start
		}
	}

	return strconv.Itoa(minLocation) 
}
