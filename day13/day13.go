package day13

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func Day13(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day13/input.txt")
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Part1(string(content))
	} else {
		part = 2
		result, err = Part2(string(content))
		fmt.Printf("Done part 2 with Chinese remainder theorem with result: %d\nNow running brute force:\n", result)
		// Originally the offset was 100000000000000 as this number is written in the task
		// But then it would take more than 5 hours to complete, so I increased the initial offset
		// after I solved the Part 2 with the help of internet
		result, err = Part2BruteForce(string(content), 1118000000000000)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Part1(input string) (result int, err error) {
	timestamp, busses, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	closestBusInMins := timestamp
	closestBusNo := 0
	for _, bus := range busses {
		if bus < 0 {
			continue
		}
		nextRunInMins := int(math.Ceil(float64(timestamp/bus))+1)*bus - timestamp
		if nextRunInMins < closestBusInMins {
			closestBusInMins = nextRunInMins
			closestBusNo = bus
		}
	}

	return closestBusInMins * closestBusNo, err
}

func Part2(input string) (result int, err error) {
	_, timetable, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	// Following Chinese remainder theorem from this video:
	// https://www.youtube.com/watch?v=zIFehsBHB8o

	N := 1
	ni, remainders := []int{}, []int{}
	for i := 0; i < len(timetable); i++ {
		if timetable[i] < 0 {
			continue
		}
		ni = append(ni, timetable[i])
		remainders = append(remainders, timetable[i]-i)
		N *= timetable[i]
	}
	for i, n := range ni {
		Ni := N / n
		result += remainders[i] * Ni * findInverse(Ni, n)
	}

	return result % N, nil
}

func findInverse(Ni, mod int) int {
	xi := Ni % mod
	for i := 0; i < mod; i++ {
		if i*xi%mod == 1 {
			return i
		}
	}

	return 0
}

func Part2BruteForce(input string, startWithTimestamp int) (result int, err error) {
	_, timetable, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	busses, offsets := []int{}, []int{}
	for i := 0; i < len(timetable); i++ {
		if timetable[i] < 0 {
			continue
		}
		busses = append(busses, timetable[i])
		offsets = append(offsets, i)
	}

	maxBusNo, maxBusOffset := 0, 0
	for i := 0; i < len(busses); i++ {
		if maxBusNo < busses[i] {
			maxBusNo = busses[i]
			maxBusOffset = offsets[i]
		}
	}

	// Find the closest time when max bus will arrive
	startWithTimestamp = int(math.Ceil(float64(startWithTimestamp/maxBusNo))) * maxBusNo
outerLoop:
	for timestamp := startWithTimestamp; timestamp < math.MaxInt64-maxBusNo*2; timestamp += maxBusNo {
		timestampOfFirstBus := timestamp - maxBusOffset
		for i := 0; i < len(busses); i++ {
			if (timestampOfFirstBus+offsets[i])%busses[i] != 0 {
				continue outerLoop
			}
		}

		return timestampOfFirstBus, nil
	}

	return 0, fmt.Errorf("No match")
}

func parseInput(input string) (timestamp int, busses []int, err error) {
	inputLines := strings.Split(input, "\n")
	timestamp, err = strconv.Atoi(inputLines[0])
	if err != nil {
		return
	}

	allBusses := strings.Split(inputLines[1], ",")
	for _, busNum := range allBusses {
		if busNum == "" {
			continue
		}
		if busNum == "x" {
			busses = append(busses, -1)
			continue
		}
		nextBus := 0
		nextBus, err = strconv.Atoi(busNum)
		if err != nil {
			return
		}
		busses = append(busses, nextBus)
	}

	return
}
