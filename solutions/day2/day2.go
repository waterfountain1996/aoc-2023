package day2

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var reGameId = regexp.MustCompile(`Game (\d+): `)
var reCubeCount = regexp.MustCompile(`(\d+) (\w+)`)

type cubeColour string

const (
	cubeColourRed   cubeColour = "red"
	cubeColourGreen cubeColour = "green"
	cubeColourBlue  cubeColour = "blue"
)

type cubeSet map[cubeColour]uint

type game struct {
	id   uint
	sets []cubeSet
}

func gameFromString(s string) *game {
	reGameId := regexp.MustCompile(`Game (\d+): `)
	gameId, _ := strconv.Atoi(reGameId.FindStringSubmatch(s)[1])

	s = s[len(reGameId.FindString(s)):]

	sets := []cubeSet{}
	for _, set := range strings.Split(s, "; ") {
		colours := make(map[cubeColour]uint)
		for _, submatch := range reCubeCount.FindAllStringSubmatch(set, -1) {
			count, colour := submatch[1], submatch[2]
			n, _ := strconv.Atoi(count)
			colours[cubeColour(colour)] = uint(n)
		}
		sets = append(sets, colours)
	}

	return &game{
		id:   uint(gameId),
		sets: sets,
	}
}

func PartA(s *bufio.Scanner) string {
	isValid := func(g *game) bool {
		constraints := cubeSet{
			"red":   12,
			"green": 13,
			"blue":  14,
		}

		for _, set := range g.sets {
			for k, v := range constraints {
				if set[k] > v {
					return false
				}
			}
		}

		return true
	}

	var result uint = 0

	for s.Scan() {
		line := s.Text()
		game := gameFromString(line)
		if isValid(game) {
			result += game.id
		}
	}

	return strconv.Itoa(int(result))
}

func PartB(s *bufio.Scanner) string {
	fewestPossibleSet := func(g *game) cubeSet {
		result := cubeSet{}
		for _, set := range g.sets {
			for k, v := range set {
				if v > result[k] {
					result[k] = v
				}
			}
		}
		return result
	}

	setProduct := func(s cubeSet) int {
		r := 1
		for _, v := range s {
			if v > 0 {
				r *= int(v)
			}
		}
		if r == 1 {
			r = 0
		}
		return r
	}

	result := 0

	for s.Scan() {
		line := s.Text()
		game := gameFromString(line)
		set := fewestPossibleSet(game)
		result += setProduct(set)
	}

	return strconv.Itoa(result)
}
