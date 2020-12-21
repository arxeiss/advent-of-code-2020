package day21

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var ingredientsAndAllergensRegex = regexp.MustCompile(`(.*) \(contains (.*)\)`)

type Meal struct {
	Ingredients []string
	Allergens   []string
}

func Day21(part int) (err error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("day21/input.txt"))
	if err != nil {
		return err
	}

	noAllergensCnt, badIngredients, err := SolvePuzzle(string(content))
	if err != nil {
		return err
	}
	stringResult := ""
	if part == 1 {
		stringResult = fmt.Sprintf("%d", noAllergensCnt)
	} else {
		part = 2
		stringResult = badIngredients
	}
	fmt.Printf("Done, result of part %d is %s \n", part, stringResult)

	return nil
}

func SolvePuzzle(input string) (noAllergensCnt int, badIngredients string, err error) {
	meals, err := parseInput(input)
	if err != nil {
		return
	}

	allergens := map[string]map[string]bool{}
	for _, meal := range meals {
		for _, allergen := range meal.Allergens {
			if possibleIngredients, exists := allergens[allergen]; exists {
				toRemove := []string{}
				for posIng := range possibleIngredients {
					found := false
					for _, ingredient := range meal.Ingredients {
						if ingredient == posIng {
							found = true
							break
						}
					}
					if !found {
						toRemove = append(toRemove, posIng)
					}
				}
				for _, removeKey := range toRemove {
					delete(possibleIngredients, removeKey)
				}
			} else {
				allergens[allergen] = map[string]bool{}
				for _, ingredient := range meal.Ingredients {
					allergens[allergen][ingredient] = true
				}
			}
		}
	}
	allAllergens := []string{}
	for allergen := range allergens {
		allAllergens = append(allAllergens, allergen)
	}

	mapAllergenToIngredient := map[string]string{}
	for len(allergens) > 0 {
		toRemove := ""
		for allergen, possibleIngredients := range allergens {
			if len(possibleIngredients) == 0 {
				delete(possibleIngredients, allergen)
				break
			}
			if len(possibleIngredients) == 1 {
				// Just to get first item without knowing the key
				ingredient := ""
				for ingredient = range possibleIngredients {
				}
				mapAllergenToIngredient[allergen] = ingredient
				delete(allergens, allergen)
				toRemove = ingredient
				break
			}
		}
		if toRemove != "" {
			for _, possibleIngredients := range allergens {
				delete(possibleIngredients, toRemove)
			}
		}
	}

	// Part 1
	for _, meal := range meals {
		for _, ingredient := range meal.Ingredients {
			found := false
			for _, v := range mapAllergenToIngredient {
				if v == ingredient {
					found = true
					break
				}
			}
			if !found {
				noAllergensCnt++
			}
		}
	}

	// Part 2
	sort.Strings(allAllergens)
	badIngredients = ""
	for _, allergen := range allAllergens {
		badIngredients += "," + mapAllergenToIngredient[allergen]
	}
	badIngredients = badIngredients[1:]

	return
}

func parseInput(input string) (meals []*Meal, err error) {
	mealLines := strings.Split(input, "\n")

	for _, m := range mealLines {
		matches := ingredientsAndAllergensRegex.FindStringSubmatch(m)
		if matches != nil {
			meals = append(meals, &Meal{
				Ingredients: strings.Split(matches[1], " "),
				Allergens:   strings.Split(matches[2], ", "),
			})
		}
	}

	return
}
