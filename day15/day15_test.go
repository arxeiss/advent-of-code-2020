package day15_test

import (
	"os"
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day15"
)

var testInput = `1,3,2`
var testInput2 = `3,1,2`

func TestPart1(t *testing.T) {
	result, err := day15.Solve(testInput, 2020)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1 {
		t.Errorf("Expected result to be 1, got %d", result)
	}
}

func TestPart1_Input2(t *testing.T) {
	result, err := day15.Solve(testInput2, 2020)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1836 {
		t.Errorf("Expected result to be 1836, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// Skip in CI
	if os.Getenv("GITHUB_SHA") != "" {
		t.Skip()
	}
	result, err := day15.Solve(testInput, 30000000)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 2578 {
		t.Errorf("Expected result to be 2578, got %d", result)
	}
}

func TestPart2_Input2(t *testing.T) {
	// Skip in CI
	if os.Getenv("GITHUB_SHA") != "" {
		t.Skip()
	}
	result, err := day15.Solve(testInput2, 30000000)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 362 {
		t.Errorf("Expected result to be 362, got %d", result)
	}
}
