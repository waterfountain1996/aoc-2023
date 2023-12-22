package day19

import (
	"bufio"
	"cmp"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	partRegex     = regexp.MustCompile(`([xmas]{1})[=](\d+)`)
	ruleRegex     = regexp.MustCompile(`([xmas]{1})([<>])(\d+)`)
	workflowRegex = regexp.MustCompile(`(\w+)\{(.*)\}`)
)

type Part struct {
	X, M, A, S int
}

func PartFromString(s string) *Part {
	var p Part

	tokens := partRegex.FindAllStringSubmatch(s, -1)
	for _, token := range tokens {
		n, _ := strconv.Atoi(token[2])
		switch token[1] {
		case "x":
			p.X = n
		case "m":
			p.M = n
		case "a":
			p.A = n
		case "s":
			p.S = n
		}
	}

	return &p
}

func (p *Part) String() string {
	return fmt.Sprintf("{x=%d,m=%d,a=%d,s=%d}", p.X, p.M, p.A, p.S)
}

func (p *Part) Rating() int {
	return p.X + p.M + p.A + p.S
}

type Workflow struct {
	Name  string
	Rules []string
}

func WorkflowFromString(s string) *Workflow {
	tokens := workflowRegex.FindStringSubmatch(s)
	return &Workflow{
		Name:  tokens[1],
		Rules: strings.Split(tokens[2], ","),
	}
}

func (w *Workflow) String() string {
	return fmt.Sprintf("%s{%s}", w.Name, strings.Join(w.Rules, ","))
}

func parseRule(rule string) (string, int, int) {
	tokens := ruleRegex.FindStringSubmatch(rule)
	key := tokens[1]
	op := cmp.Compare(tokens[2], "=")
	value, _ := strconv.Atoi(tokens[3])
	return key, op, value
}

func (w *Workflow) Perform(p *Part) string {
	for _, rule := range w.Rules {
		expr, dst, found := strings.Cut(rule, ":")
		if !found {
			return rule
		}

		fmt.Printf("%s ? %s\n", expr, dst)

		key, op, value := parseRule(expr)
		var metric int

		switch key {
		case "x":
			metric = p.X
		case "m":
			metric = p.M
		case "a":
			metric = p.A
		case "s":
			metric = p.S
		}

		if cmp.Compare(metric, value) == op {
			fmt.Printf("  true\n")
			return dst
		} else {
			fmt.Printf("  false\n")
		}
	}

	return ""
}

func parseInput(s *bufio.Scanner) (map[string]*Workflow, []*Part) {
	workflows := make(map[string]*Workflow)
	parts := []*Part{}

	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		w := WorkflowFromString(line)
		workflows[w.Name] = w
	}

	for s.Scan() {
		parts = append(parts, PartFromString(s.Text()))
	}

	return workflows, parts
}

func PartA(s *bufio.Scanner) string {
	workflows, parts := parseInput(s)

	sum := 0

	for _, part := range parts {
		w := workflows["in"]

		for {
			dst := w.Perform(part)
			if dst == "A" {
				sum += part.Rating()
				break
			} else if dst == "R" {
				break
			}

			w = workflows[dst]
		}
	}

	return strconv.Itoa(sum)
}

type Range [2]int

func (r Range) String() string {
	return fmt.Sprintf("[%d; %d]", r[0], r[1])
}

func (r Range) Length() int {
	return r[1] - r[0] + 1
}

func (r Range) Split(at int, boundary int) []Range {
	end := at
	if boundary == -1 {
		end--
	}

	return []Range{
		{r[0], end},
		{end + 1, r[1]},
	}
}

type RangePart struct {
	X, M, A, S Range
}

func NewRangePart(x, m, a, s Range) *RangePart {
	return &RangePart{
		X: x,
		M: m,
		A: a,
		S: s,
	}
}

func (p *RangePart) Length() int {
	return p.X.Length() * p.M.Length() * p.A.Length() * p.S.Length()
}

func findCombinations(workflows map[string]*Workflow, wName string, p *RangePart) int {
	if wName == "A" {
		return p.Length()
	}

	w := workflows[wName]
	if w == nil {
		return 0
	}

	sum := 0

	for _, rule := range w.Rules {
		expr, dst, found := strings.Cut(rule, ":")
		if !found {
			return sum + findCombinations(workflows, rule, p)
		}

		key, op, value := parseRule(expr)

		var newRanges []Range

		switch key {
		case "x":
			newRanges = p.X.Split(value, op)
		case "m":
			newRanges = p.M.Split(value, op)
		case "a":
			newRanges = p.A.Split(value, op)
		case "s":
			newRanges = p.S.Split(value, op)
		}

		var toForward, toKeep Range

		if op < 0 {
			toForward = newRanges[0]
			toKeep = newRanges[1]
		} else {
			toForward = newRanges[1]
			toKeep = newRanges[0]
		}

		newPart := NewRangePart(p.X, p.M, p.A, p.S)

		switch key {
		case "x":
			newPart.X = toForward
			p.X = toKeep
		case "m":
			newPart.M = toForward
			p.M = toKeep
		case "a":
			newPart.A = toForward
			p.A = toKeep
		case "s":
			newPart.S = toForward
			p.S = toKeep
		}

		sum += findCombinations(workflows, dst, newPart)
	}

	return sum
}

func PartB(s *bufio.Scanner) string {
	workflows, _ := parseInput(s)

	result := findCombinations(workflows, "in", NewRangePart(
		Range{1, 4000},
		Range{1, 4000},
		Range{1, 4000},
		Range{1, 4000},
	))

	return strconv.Itoa(result)
}
