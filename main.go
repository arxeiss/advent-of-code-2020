package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arxeiss/advent-of-code-2020/day1"
	"github.com/arxeiss/advent-of-code-2020/day10"
	"github.com/arxeiss/advent-of-code-2020/day11"
	"github.com/arxeiss/advent-of-code-2020/day12"
	"github.com/arxeiss/advent-of-code-2020/day13"
	"github.com/arxeiss/advent-of-code-2020/day14"
	"github.com/arxeiss/advent-of-code-2020/day15"
	"github.com/arxeiss/advent-of-code-2020/day16"
	"github.com/arxeiss/advent-of-code-2020/day17"
	"github.com/arxeiss/advent-of-code-2020/day18"
	"github.com/arxeiss/advent-of-code-2020/day19"
	"github.com/arxeiss/advent-of-code-2020/day2"
	"github.com/arxeiss/advent-of-code-2020/day20"
	"github.com/arxeiss/advent-of-code-2020/day21"
	"github.com/arxeiss/advent-of-code-2020/day22"
	"github.com/arxeiss/advent-of-code-2020/day23"
	"github.com/arxeiss/advent-of-code-2020/day24"
	"github.com/arxeiss/advent-of-code-2020/day25"
	"github.com/arxeiss/advent-of-code-2020/day3"
	"github.com/arxeiss/advent-of-code-2020/day4"
	"github.com/arxeiss/advent-of-code-2020/day5"
	"github.com/arxeiss/advent-of-code-2020/day6"
	"github.com/arxeiss/advent-of-code-2020/day7"
	"github.com/arxeiss/advent-of-code-2020/day8"
	"github.com/arxeiss/advent-of-code-2020/day9"
)

func main() {
	fmt.Println("Welcome to Advent of Code 2020 in Go!")
	day, part := 0, 0
	var err error

	// Handle Day
	if len(os.Args) > 1 {
		day, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Cannot parse '%s' into number, %v\n", os.Args[1], err)
			os.Exit(1)
		}
	} else {
		fmt.Print("Enter which Day to run (1-25): ")
		_, err = fmt.Scanf("%d", &day)
		if err != nil {
			fmt.Printf("Cannot scan number, %v\n", err)
			os.Exit(1)
		}
	}

	// Handle Part
	if len(os.Args) > 2 {
		part, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Cannot parse '%s' into number, %v\n", os.Args[2], err)
			os.Exit(1)
		}
	} else {
		fmt.Print("Enter which Part to run (1 or 2): ")
		_, err = fmt.Scanf("%d", &part)
		if err != nil {
			fmt.Printf("Cannot scan stdin, %v\n", err)
			os.Exit(1)
		}
	}
	if part != 1 && part != 2 {
		fmt.Printf("Part to run can be either 1 or 2, got %d\n", part)
		os.Exit(1)
	}

	start := time.Now()
	switch day {
	case 1:
		err = day1.Day1(part)
	case 2:
		err = day2.Day2(part)
	case 3:
		err = day3.Day3(part)
	case 4:
		err = day4.Day4(part)
	case 5:
		err = day5.Day5(part)
	case 6:
		err = day6.Day6(part)
	case 7:
		err = day7.Day7(part)
	case 8:
		err = day8.Day8(part)
	case 9:
		err = day9.Day9(part)
	case 10:
		err = day10.Day10(part)
	case 11:
		err = day11.Day11(part)
	case 12:
		err = day12.Day12(part)
	case 13:
		err = day13.Day13(part)
	case 14:
		err = day14.Day14(part)
	case 15:
		err = day15.Day15(part)
	case 16:
		err = day16.Day16(part)
	case 17:
		err = day17.Day17(part)
	case 18:
		err = day18.Day18(part)
	case 19:
		err = day19.Day19(part)
	case 20:
		err = day20.Day20(part)
	case 21:
		err = day21.Day21(part)
	case 22:
		err = day22.Day22(part)
	case 23:
		err = day23.Day23(part)
	case 24:
		err = day24.Day24(part)
	case 25:
		err = day25.Day25(part)
	default:
		err = fmt.Errorf("There is no such a day %d", day)
	}

	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Solving puzzle for day %d & part %d took %s\nAll worked! Good bye\n", day, part, time.Since(start))
}
