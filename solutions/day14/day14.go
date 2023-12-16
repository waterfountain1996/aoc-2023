package day14

import (
	"bufio"
	"crypto/md5"
	"slices"
	"strconv"
	"strings"
)

func replaceAtIndex(s string, ch rune, index int) string {
	out := []rune(s)
	out[index] = ch
	return string(out)
}

type Platform []string

type TiltDirection uint8

const (
	TiltNorth TiltDirection = iota
	TiltWest
	TiltSouth
	TiltEast
)

func parseInput(s *bufio.Scanner) Platform {
	platform := Platform{}
	for s.Scan() {
		platform = append(platform, s.Text())
	}
	return platform
}

func (p Platform) TiltNorth() Platform {
	tilted := make(Platform, len(p))
	copy(tilted, p)

	for i := 1; i < len(p); i++ {
		row := p[i]

		for j, ch := range row {
			if ch != 'O' {
				continue
			}

			newRow := -1

			for slide := i - 1; slide >= 0 && tilted[slide][j] == '.'; slide-- {
				newRow = slide
			}

			if newRow != -1 {
				tilted[newRow] = replaceAtIndex(tilted[newRow], 'O', j)
				tilted[i] = replaceAtIndex(tilted[i], '.', j)
			}
		}
	}

	return tilted
}

func (p Platform) Tilt(dir TiltDirection) Platform {
	tilted := make(Platform, len(p))
	copy(tilted, p)

	s1, s2 := "O.", ".O"

	switch dir {
	case TiltNorth:
		return p.TiltNorth()
	case TiltSouth:
		slices.Reverse(tilted)
		tilted = tilted.TiltNorth()
		slices.Reverse(tilted)
	case TiltWest:
		s2, s1 = s1, s2
		fallthrough
	case TiltEast:
		for i, row := range tilted {
			before := row
			after  := strings.ReplaceAll(row, s1, s2)
			for after != before {
				before = after
				after  = strings.ReplaceAll(after, s1, s2)
			}

			tilted[i] = after
		}
	}

	return tilted
}

func (p Platform) SpinCycle() Platform {
	directions := []TiltDirection{
		TiltNorth,
		TiltWest,
		TiltSouth,
		TiltEast,
	}

	tilted := p
	for _, dir := range directions {
		tilted = tilted.Tilt(dir)
	}
	return tilted
}

func (p Platform) NorthSupportBeamLoad() int {
	load := 0
	for i, row := range p {
		weight := len(p) - i
		load += strings.Count(row, "O") * weight
	}
	return load
}

func (p Platform) Hexsum() [16]byte {
	data := []byte{}
	for _, row := range p {
		data = append(data, []byte(row)...)
	}
	return md5.Sum(data)
}

func PartA(s *bufio.Scanner) string {
	platform := parseInput(s)
	tilted := platform.TiltNorth()
	load := tilted.NorthSupportBeamLoad()
	return strconv.Itoa(load)
}

func PartB(s *bufio.Scanner) string {
	platform := parseInput(s)

	cycles := 1000000000

	seen := make(map[[16]byte]int)
	seen[platform.Hexsum()] = 0

	for i := 0; i < cycles; i++ {
		platform = platform.SpinCycle()

		sum := platform.Hexsum()

		if cycleStart := seen[sum]; cycleStart > 0 {
			cycleLength := i + 1 - cycleStart
			iterations  := (cycles - cycleStart) % cycleLength

			for j := 0; j < iterations; j++ {
				platform = platform.SpinCycle()
			}

			break
		}

		seen[sum] = i + 1
	}

	load := platform.NorthSupportBeamLoad()
	return strconv.Itoa(load)
}
