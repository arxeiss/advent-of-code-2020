package day22

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var tileRegex = regexp.MustCompile(`Tile (\d+):`)

func Day22(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day22/input.txt"))
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Part1(string(content))
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

func Part1(input string) (result int, err error) {
	player1, player2, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	player1Card, player2Card := 0, 0
	for len(player1) > 0 && len(player2) > 0 {
		player1Card, player1 = player1[0], player1[1:]
		player2Card, player2 = player2[0], player2[1:]

		if player1Card > player2Card {
			player1 = append(player1, player1Card, player2Card)
		} else if player2Card > player1Card {
			player2 = append(player2, player2Card, player1Card)
		} else {
			return 0, fmt.Errorf("Player 1 & 2 has same card value (%d-%d)", player1Card, player2Card)
		}
	}

	winningCards := player1
	if len(winningCards) == 0 {
		winningCards = player2
	}

	for i, m := 0, len(winningCards); i < len(winningCards); i, m = i+1, m-1 {
		result += winningCards[i] * m
	}

	return result, err
}

func Part2(input string) (result int, err error) {
	player1, player2, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	player1, player2, err = playSubGame(player1, player2)
	if err != nil {
		return 0, err
	}

	winningCards := player1
	if len(winningCards) == 0 {
		winningCards = player2
	}

	for i, m := 0, len(winningCards); i < len(winningCards); i, m = i+1, m-1 {
		result += winningCards[i] * m
	}

	return result, nil
}

func playSubGame(player1Src, player2Src []int) ([]int, []int, error) {
	playedConfiguration := map[string]bool{}
	player1, player2 := []int{}, []int{}
	for _, v := range player1Src {
		player1 = append(player1, v)
	}
	for _, v := range player2Src {
		player2 = append(player2, v)
	}

	player1Card, player2Card := 0, 0
	for len(player1) > 0 && len(player2) > 0 {
		currentConfig := fmt.Sprintf("%v=%v", player1, player2)
		if _, exists := playedConfiguration[currentConfig]; exists {
			return player1, []int{}, nil
		}
		playedConfiguration[currentConfig] = true

		player1Card, player1 = player1[0], player1[1:]
		player2Card, player2 = player2[0], player2[1:]

		if len(player1) >= player1Card && len(player2) >= player2Card {
			player1Result, player2Result, err := playSubGame(player1[0:player1Card], player2[0:player2Card])
			if err != nil {
				return player1Result, player2Result, err
			}
			if len(player1Result) > len(player2Result) {
				player1 = append(player1, player1Card, player2Card)
			} else if len(player2Result) > len(player1Result) {
				player2 = append(player2, player2Card, player1Card)
			} else {
				return []int{}, []int{}, fmt.Errorf(
					"Player 1 & 2 has same subgame result (%v-%v)",
					player1Result,
					player2Result,
				)
			}
		} else if player1Card > player2Card {
			player1 = append(player1, player1Card, player2Card)
		} else if player2Card > player1Card {
			player2 = append(player2, player2Card, player1Card)
		} else {
			return []int{}, []int{}, fmt.Errorf("Player 1 & 2 has same card value (%d-%d)", player1Card, player2Card)
		}
	}

	return player1, player2, nil
}

func parseInput(input string) (player1, player2 []int, err error) {
	lines := strings.Split(input, "\n")
	card := 0

	currPlayer := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Player") {
			currPlayer++
			continue
		}
		card, err = strconv.Atoi(line)
		if currPlayer == 1 {
			player1 = append(player1, card)
		} else {
			player2 = append(player2, card)
		}
	}

	return
}
