package day18

import (
	"fmt"
	"regexp"
	"strconv"
)

type TokenType uint8

const (
	EOL TokenType = iota
	Number
	Add
	Mul
	LPar
	RPar
)

func (tt TokenType) String() string {
	p := []string{"EOL", "Number", "+", "*", "(", ")"}
	return fmt.Sprintf("%s", p[tt])
}

type Token struct {
	TType TokenType
	Value int
}
type Expression []Token

func (t Token) String() string {
	if t.TType == Number {
		return fmt.Sprintf("%d", t.Value)
	}
	p := []string{"", "N", "+", "*", "(", ")"}
	return fmt.Sprintf("%s", p[t.TType])
}

var tokenizeRegex = regexp.MustCompile(`(\d+|\+|\*|\(|\))`)

func Tokenize(input string) (Expression, error) {
	tokens := []Token{}
	// To repeat else part again and save last number
	input += " "
	numberStartsAt := -1
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			if numberStartsAt < 0 {
				numberStartsAt = i
			}
		} else {
			if numberStartsAt >= 0 {
				val, err := strconv.Atoi(input[numberStartsAt:i])
				if err != nil {
					return nil, err
				}
				numberStartsAt = -1
				tokens = append(tokens, Token{TType: Number, Value: val})
			}
			switch input[i] {
			case '+':
				tokens = append(tokens, Token{TType: Add})
			case '*':
				tokens = append(tokens, Token{TType: Mul})
			case '(':
				tokens = append(tokens, Token{TType: LPar})
			case ')':
				tokens = append(tokens, Token{TType: RPar})
			}
		}
	}
	tokens = append(tokens, Token{TType: EOL})
	return tokens, nil
}

func TokenizeWithRegex(input string) (Expression, error) {
	tokens := []Token{}
	matches := tokenizeRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		switch match[0] {
		case "+":
			tokens = append(tokens, Token{TType: Add})
		case "*":
			tokens = append(tokens, Token{TType: Mul})
		case "(":
			tokens = append(tokens, Token{TType: LPar})
		case ")":
			tokens = append(tokens, Token{TType: RPar})
		default:
			val, err := strconv.Atoi(match[0])
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{TType: Number, Value: val})
		}
	}
	tokens = append(tokens, Token{TType: EOL})
	return tokens, nil
}
