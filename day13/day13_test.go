package day13_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day13"
)

var testInput = `939
7,13,x,x,59,x,31,19`

var test2SecondInput = `111
1789,37,47,1889`

func TestPart1(t *testing.T) {
	result, err := day13.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 295 {
		t.Errorf("Expected result to be 295, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day13.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1068781 {
		t.Errorf("Expected result to be 1068781, got %d", result)
	}
}

func TestPart2_Second(t *testing.T) {
	result, err := day13.Part2(test2SecondInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1202161486 {
		t.Errorf("Expected result to be 1202161486, got %d", result)
	}
}

func TestPart2BruteForce(t *testing.T) {
	result, err := day13.Part2BruteForce(testInput, 0)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1068781 {
		t.Errorf("Expected result to be 1068781, got %d", result)
	}
}

func TestPart2BruteForce_Second(t *testing.T) {
	result, err := day13.Part2BruteForce(test2SecondInput, 0)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1202161486 {
		t.Errorf("Expected result to be 1202161486, got %d", result)
	}
}
