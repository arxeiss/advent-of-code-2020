package day16_test

import (
	"io/ioutil"
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day16"
)

var testInput = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestPart1(t *testing.T) {
	result, err := day16.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 71 {
		t.Errorf("Expected result to be 71, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// This one was not really used during the development, because there were no real test data
	// So after I solved this, I added this to test just to handle possible changes in the code in the future
	content, _ := ioutil.ReadFile("input.txt")
	result, err := day16.Part2(string(content))
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1305243193339 {
		t.Errorf("Expected result to be 1305243193339, got %d", result)
	}
}
