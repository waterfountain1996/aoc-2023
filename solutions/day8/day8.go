package day8

import (
	"bufio"
	"regexp"
	"strconv"
)

var segmentRegex = regexp.MustCompile(`([A-Z0-9]{3})`)

type NetworkPath struct {
	Left, Right string
}

type Network map[string]NetworkPath

func parseInput(s *bufio.Scanner) (string, Network) {
	s.Scan()

	instructions := s.Text()
	network := make(Network)

	s.Scan()
	for s.Scan() {
		elements := segmentRegex.FindAllString(s.Text(), -1)
		network[elements[0]] = NetworkPath{
			Left:  elements[1],
			Right: elements[2],
		}
	}

	return instructions, network
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func solve(startingNode, instructions string ,network Network, atEnd func(string) bool) int {
	node := startingNode

	idx := 0
	steps := 0

	for !atEnd(node) {
		path := network[node]

		direction := rune(instructions[idx]); if direction == 'L' {
			node = path.Left
		} else {
			node = path.Right
		}

		steps++
		idx++
		idx %= len(instructions)
	}

	return steps
}

func PartA(s *bufio.Scanner) string {
	instructions, network := parseInput(s)
	steps := solve("AAA", instructions, network, func (node string) bool {
		return node == "ZZZ"
	})
	return strconv.Itoa(steps)
}

func PartB(s *bufio.Scanner) string {
	instructions, network := parseInput(s)

	result := 1
	for node := range network {
		if rune(node[2]) != 'A' {
			continue
		}

		steps := solve(node, instructions, network, func (node string ) bool {
			return rune(node[2]) == 'Z'
		})
		result = lcm(result, steps)
	}
	return strconv.Itoa(result)
}
