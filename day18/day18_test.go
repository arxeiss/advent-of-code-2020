package day18_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day18"
)

func TestPart1Example0(t *testing.T) {
	result, err := day18.Part1("1 + (2 * 3) + (4 * (5 + 6))")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 51 {
		t.Errorf("Expected result to be 51, got %d", result)
	}
}

func TestPart1Example1(t *testing.T) {
	result, err := day18.Part1("2 * 3 + (4 * 5)")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 26 {
		t.Errorf("Expected result to be 26, got %d", result)
	}
}
func TestPart1Example2(t *testing.T) {
	result, err := day18.Part1("5 + (8 * 3 + 9 + 3 * 4 * 3)")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 437 {
		t.Errorf("Expected result to be 437, got %d", result)
	}
}
func TestPart1Example3(t *testing.T) {
	result, err := day18.Part1("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 12240 {
		t.Errorf("Expected result to be 12240, got %d", result)
	}
}
func TestPart1Example4(t *testing.T) {
	result, err := day18.Part1("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 13632 {
		t.Errorf("Expected result to be 13632, got %d", result)
	}
}

func TestPart2Example0(t *testing.T) {
	result, err := day18.Part2("1 + (2 * 3) + (4 * (5 + 6))")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 51 {
		t.Errorf("Expected result to be 51, got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	result, err := day18.Part2("2 * 3 + (4 * 5)")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 46 {
		t.Errorf("Expected result to be 46, got %d", result)
	}
}
func TestPart2Example2(t *testing.T) {
	result, err := day18.Part2("5 + (8 * 3 + 9 + 3 * 4 * 3)")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1445 {
		t.Errorf("Expected result to be 1445, got %d", result)
	}
}
func TestPart2Example3(t *testing.T) {
	result, err := day18.Part2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 669060 {
		t.Errorf("Expected result to be 669060, got %d", result)
	}
}
func TestPart2Example4(t *testing.T) {
	result, err := day18.Part2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 23340 {
		t.Errorf("Expected result to be 23340, got %d", result)
	}
}
