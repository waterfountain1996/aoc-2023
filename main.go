package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/waterfountain1996/aoc-2023/solutions/day1"
	"github.com/waterfountain1996/aoc-2023/solutions/day2"
	"github.com/waterfountain1996/aoc-2023/solutions/day3"
	"github.com/waterfountain1996/aoc-2023/solutions/day4"
	"github.com/waterfountain1996/aoc-2023/solutions/day5"
	"github.com/waterfountain1996/aoc-2023/solutions/day6"
	"github.com/waterfountain1996/aoc-2023/solutions/day7"
	"github.com/waterfountain1996/aoc-2023/solutions/day8"
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
	solutions["day4:a"] = day4.PartA
	solutions["day4:b"] = day4.PartB
	solutions["day5:a"] = day5.PartA
	solutions["day5:b"] = day5.PartB
	solutions["day6:a"] = day6.PartA
	solutions["day6:b"] = day6.PartB
	solutions["day7:a"] = day7.PartA
	solutions["day7:b"] = day7.PartB
	solutions["day8:a"] = day8.PartA
	solutions["day8:b"] = day8.PartB

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
