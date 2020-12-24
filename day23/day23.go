package day23

import (
	"fmt"
	"io/ioutil"
)

type Element struct {
	Value int
	Next  *Element
}

type List struct {
	FirstElement *Element
	LastElement  *Element
	hashMap      map[int]*Element
}

func (list *List) ToArray() (arr []int) {
	el := list.FirstElement
	for el != nil {
		arr = append(arr, el.Value)
		el = el.Next
	}
	return arr
}

func (list *List) Append(newEl *Element) {
	if list.FirstElement == nil {
		list.FirstElement = newEl
		list.LastElement = newEl
	} else {
		if list.LastElement == nil {
			list.RefreshLast()
		}
		list.LastElement.Next = newEl
		list.LastElement = newEl
	}
}

func (list *List) Max() (maxVal int) {
	el := list.FirstElement
	for el != nil {
		if el.Value > maxVal {
			maxVal = el.Value
		}
		el = el.Next
	}
	return
}

func (list *List) FindVal(val int) *Element {
	if list.hashMap != nil {
		return list.hashMap[val]
	}
	el := list.FirstElement
	for el != nil {
		if el.Value == val {
			return el
		}
		el = el.Next
	}
	return nil
}

func (list *List) GetLast() *Element {
	return list.LastElement
}

func (list *List) RefreshLast() {
	el := list.FirstElement
	for el.Next != nil {
		el = el.Next
	}
	list.LastElement = el
}

func (list *List) BuildMap() {
	list.hashMap = make(map[int]*Element)
	el := list.FirstElement
	for el != nil {
		list.hashMap[el.Value] = el
		el = el.Next
	}
}

func Day23(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day23/input.txt"))
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Part1(string(content), 100)
	} else {
		part = 2
		result, err = Part2(string(content))
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Part1(input string, rounds int) (result int, err error) {
	cupsList, err := parseInput(input)
	_ = cupsList.ToArray() // Force compiler not to remove method
	if err != nil {
		return 0, err
	}
	cupsList.BuildMap()

	cupsList = playGame(cupsList, rounds)

	elem1 := cupsList.FindVal(1)
	currElem := elem1.Next
	if currElem == nil {
		currElem = cupsList.FirstElement
	}

	for currElem.Value != 1 {
		result = result*10 + currElem.Value

		currElem = currElem.Next
		if currElem == nil {
			currElem = cupsList.FirstElement
		}
	}

	return result, err
}

func Part2(input string) (result int, err error) {
	cupsList, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	for i := cupsList.Max() + 1; i <= 1_000_000; i++ {
		cupsList.Append(&Element{Value: i})
	}

	cupsList.BuildMap()

	cupsList = playGame(cupsList, 10_000_000)

	elem1 := cupsList.FindVal(1)
	f1 := elem1.Next
	if f1 == nil {
		f1 = cupsList.FirstElement
	}
	f2 := f1.Next
	if f2 == nil {
		f2 = cupsList.FirstElement
	}

	result = f1.Value * f2.Value

	return result, err
}

func playGame(cupsList *List, rounds int) *List {
	maxCup := cupsList.Max()
	for i := 0; i < rounds; i++ {
		current := cupsList.FirstElement
		removed := current.Next
		current.Next = nil

		rest := removed.Next.Next.Next
		removed.Next.Next.Next = nil // Cut removed list to create sublist
		removedList := &List{FirstElement: removed}
		restList := &List{FirstElement: rest, LastElement: cupsList.LastElement, hashMap: cupsList.hashMap}

		destination := current.Value - 1
		for removedList.FindVal(destination) != nil || destination < 1 {
			if destination < 1 {
				destination = maxCup
			} else {
				destination--
			}
		}

		destElement := restList.FindVal(destination)
		afterDest := destElement.Next
		destElement.Next = removed
		removed.Next.Next.Next = afterDest
		if afterDest == nil {
			restList.LastElement = removed.Next.Next
		}
		restList.Append(current)

		cupsList = restList
	}

	return cupsList
}

func parseInput(input string) (list *List, err error) {
	list = &List{}
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			break
		}
		list.Append(&Element{Value: int(input[i] - '0')})
	}
	return
}
