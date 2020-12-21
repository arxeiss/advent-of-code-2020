package day21_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day21"
)

var testInput = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestSolvePuzzle(t *testing.T) {
	noAllergensCnt, badIngredients, err := day21.SolvePuzzle(testInput)
	t.Log("noAllergensCnt=", noAllergensCnt, "badIngredients=", badIngredients, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if noAllergensCnt != 5 {
		t.Errorf("Expected noAllergensCnt to be 5, got %d", noAllergensCnt)
	}
	if badIngredients != "mxmxvkd,sqjhc,fvjkl" {
		t.Errorf("Expected badIngredients to be 'mxmxvkd,sqjhc,fvjkl', got '%s'", badIngredients)
	}
}
