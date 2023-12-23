package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/waterfountain1996/aoc-2023/solutions/day1"
	"github.com/waterfountain1996/aoc-2023/solutions/day10"
	"github.com/waterfountain1996/aoc-2023/solutions/day11"
	"github.com/waterfountain1996/aoc-2023/solutions/day12"
	"github.com/waterfountain1996/aoc-2023/solutions/day13"
	"github.com/waterfountain1996/aoc-2023/solutions/day14"
	"github.com/waterfountain1996/aoc-2023/solutions/day15"
	"github.com/waterfountain1996/aoc-2023/solutions/day16"
	"github.com/waterfountain1996/aoc-2023/solutions/day18"
	"github.com/waterfountain1996/aoc-2023/solutions/day19"
	"github.com/waterfountain1996/aoc-2023/solutions/day2"
	"github.com/waterfountain1996/aoc-2023/solutions/day20"
	"github.com/waterfountain1996/aoc-2023/solutions/day21"
	"github.com/waterfountain1996/aoc-2023/solutions/day3"
	"github.com/waterfountain1996/aoc-2023/solutions/day4"
	"github.com/waterfountain1996/aoc-2023/solutions/day5"
	"github.com/waterfountain1996/aoc-2023/solutions/day6"
	"github.com/waterfountain1996/aoc-2023/solutions/day7"
	"github.com/waterfountain1996/aoc-2023/solutions/day8"
	"github.com/waterfountain1996/aoc-2023/solutions/day9"
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
	solutions["day9:a"] = day9.PartA
	solutions["day9:b"] = day9.PartB
	solutions["day10:a"] = day10.PartA
	solutions["day10:b"] = day10.PartB
	solutions["day11:a"] = day11.PartA
	solutions["day11:b"] = day11.PartB
	solutions["day12:a"] = day12.PartA
	solutions["day12:b"] = day12.PartB
	solutions["day13:a"] = day13.PartA
	solutions["day13:b"] = day13.PartB
	solutions["day14:a"] = day14.PartA
	solutions["day14:b"] = day14.PartB
	solutions["day15:a"] = day15.PartA
	solutions["day15:b"] = day15.PartB
	solutions["day16:a"] = day16.PartA
	solutions["day16:b"] = day16.PartB
	solutions["day18:a"] = day18.PartA
	solutions["day18:b"] = day18.PartB
	solutions["day19:a"] = day19.PartA
	solutions["day19:b"] = day19.PartB
	solutions["day20:a"] = day20.PartA
	solutions["day20:b"] = day20.PartB
	solutions["day21:a"] = day21.PartA
	solutions["day21:b"] = day21.PartB

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
