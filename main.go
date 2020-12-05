package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arxeiss/advent-of-code-2020/day1"
	"github.com/arxeiss/advent-of-code-2020/day2"
	"github.com/arxeiss/advent-of-code-2020/day3"
	"github.com/arxeiss/advent-of-code-2020/day4"
	"github.com/arxeiss/advent-of-code-2020/day5"
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
		fmt.Print("Enter which Day to run (1-5): ")
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
	default:
		err = fmt.Errorf("There is no such a day %d", day)
	}

	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("All worked! Good bye\n")
}
