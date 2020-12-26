package day25

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Day25(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day25/input.txt"))
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Part1(string(content))
	} else {
		return fmt.Errorf("There is not part 2 for 25th day")
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Part1(input string) (result int, err error) {
	cardsPublicKey, doorsPublicKey, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	doorsLoopSize := getLoopSize(doorsPublicKey)
	result = getEncryptionKey(cardsPublicKey, doorsLoopSize)

	cardsLoopSize := getLoopSize(cardsPublicKey)
	encryptKey := getEncryptionKey(doorsPublicKey, cardsLoopSize)

	if result != encryptKey {
		err = fmt.Errorf("Bad encryption key, %d != %d", result, encryptKey)
	}

	return result, err
}

func getLoopSize(publicKey int) (loopSize int) {
	subject := 7
	searchedPublicKey := 1
	for loopSize = 1; true; loopSize++ {
		searchedPublicKey = (searchedPublicKey * subject) % 20201227
		if searchedPublicKey == publicKey {
			return
		}
	}
	return -1
}

func getEncryptionKey(subject, loopSize int) (encryptionKey int) {
	encryptionKey = 1
	for i := 0; i < loopSize; i++ {
		encryptionKey = (encryptionKey * subject) % 20201227
	}
	return
}

func parseInput(input string) (cardsPublicKey, doorsPublicKey int, err error) {
	lines := strings.Split(input, "\n")

	if len(lines) < 2 {
		err = fmt.Errorf("Too short input, not enough lines")
		return
	}

	cardsPublicKey, err = strconv.Atoi(lines[0])
	if err != nil {
		return
	}
	doorsPublicKey, err = strconv.Atoi(lines[1])
	if err != nil {
		return
	}

	return
}
