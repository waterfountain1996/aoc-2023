package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/waterfountain1996/aoc-2023/solutions/day1"
	"github.com/waterfountain1996/aoc-2023/solutions/day2"
	"github.com/waterfountain1996/aoc-2023/solutions/day3"
)

type Solution func(*bufio.Scanner) string

func main() {
	solutions := make(map[string]Solution)

	solutions["day1:a"] = day1.PartA
	solutions["day1:b"] = day1.PartB
	solutions["day2:a"] = day2.PartA
	solutions["day2:b"] = day2.PartB
	solutions["day3:a"] = day3.PartA
	solutions["day3:b"] = day3.PartB

	var day int
	var part string

	flag.IntVar(&day, "day", 1, "1-25")
	flag.StringVar(&part, "part", "a", "a or b")

	flag.Parse()

	s := bufio.NewScanner(os.Stdin)

	key := fmt.Sprintf("day%v:%v", day, part)
	sol := solutions[key]
	if sol == nil {
		fmt.Fprintf(os.Stderr, "No solution for day %d part %s\n", day, part)
		os.Exit(1)
	}

	result := sol(s)
	fmt.Printf("Day %d part %s:\n%s\n", day, part, result)
}
