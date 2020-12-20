package day18

import (
	"fmt"
	"strings"
)

type OperatorPriority uint8

const (
	Normal OperatorPriority = iota
	SameMulAsAdd
	FlipAddAndMul
)

type AST interface {
	Eval() int
}

type BinOperation struct {
	Left      AST
	Operation TokenType
	Right     AST
}

func (op *BinOperation) Eval() int {
	l := op.Left.Eval()
	r := op.Right.Eval()

	if op.Operation == Mul {
		return l * r
	}
	return l + r
}

type Value int

func (n Value) Eval() int {
	return int(n)
}

type Parser struct {
	expr             Expression
	currentPosition  int
	operatorPriority OperatorPriority
}

func NewParser(expr Expression, operatorPriority OperatorPriority) *Parser {
	return &Parser{expr: expr, operatorPriority: operatorPriority}
}

func (p *Parser) pop(expectedTypes ...TokenType) (TokenType, error) {
	anyOf := []string{}
	for _, expType := range expectedTypes {
		if p.current().TType == expType {
			p.currentPosition++
			return expType, nil
		}
		anyOf = append(anyOf, fmt.Sprintf("'%s'", expType))
	}
	if len(anyOf) > 1 {
		return p.current().TType, fmt.Errorf(
			"Expected one of [%s] got '%s' on position %d",
			strings.Join(anyOf, ", "),
			p.current().TType,
			p.currentPosition,
		)
	}
	return p.current().TType, fmt.Errorf(
		"Expected  %s got '%s' on position %d",
		anyOf[0],
		p.current().TType,
		p.currentPosition,
	)
}

func (p *Parser) current() Token {
	return p.expr[p.currentPosition]
}

func (p *Parser) isCurrentType(tokenType TokenType) bool {
	return p.expr[p.currentPosition].TType == tokenType
}

func (p *Parser) shouldTerminateExpression() (_ TokenType, should bool) {
	should = false
	if p.isCurrentType(EOL) || p.isCurrentType(RPar) {
		should = true
	}
	return p.current().TType, should
}

func (p *Parser) getOperatorPriority(tokenType TokenType) int {
	if p.operatorPriority == SameMulAsAdd && tokenType == Mul {
		tokenType = Add
	} else if p.operatorPriority == FlipAddAndMul {
		if tokenType == Mul {
			tokenType = Add
		} else if tokenType == Add {
			tokenType = Mul
		}
	}

	switch tokenType {
	case Number:
		return 4
	case LPar:
		return 3
	case Mul:
		return 2
	case Add:
		return 1
	}
	return 0
}

func (p *Parser) handleTerm() (node AST, err error) {
	if p.isCurrentType(LPar) {
		if _, err = p.pop(LPar); err != nil {
			return
		}
		if node, err = p.handleExpression(p.getOperatorPriority(EOL)); err != nil {
			return
		}
		_, err = p.pop(RPar)
		return
	}
	if p.isCurrentType(Number) {
		node = Value(p.current().Value)
		_, err = p.pop(Number)
		return
	}
	node = nil
	return
}

// N: 1-9
// T: E | (E)
// E: T * T | T + T | N
func (p *Parser) handleExpression(currentPriority int) (node AST, err error) {
	var rNode AST
	var operator TokenType
	if p.getOperatorPriority(p.current().TType) > currentPriority {
		node, err = p.handleExpression(currentPriority + 1)
		if err != nil {
			return
		}
	}
	if termNode, err := p.handleTerm(); err != nil || termNode != nil {
		return termNode, err
	}

	for p.getOperatorPriority(p.current().TType) == currentPriority {
		// Exit early if there is end
		if _, should := p.shouldTerminateExpression(); should {
			return
		}
		if operator, err = p.pop(Mul, Add, EOL); err != nil {
			return
		}
		if operator == EOL {
			return
		}
		rNode, err = p.handleExpression(currentPriority + 1)
		if rNode == nil {
			err = fmt.Errorf("Unexpected token '%s' on position %d", p.current().TType, p.currentPosition)
		}
		if err != nil {
			return
		}

		node = &BinOperation{
			Left:      node,
			Operation: operator,
			Right:     rNode,
		}

	}
	return
}

func (p *Parser) Run() (ast AST, err error) {
	ast, err = p.handleExpression(p.getOperatorPriority(EOL))
	if err == nil && p.currentPosition < len(p.expr)-1 {
		err = fmt.Errorf("Unexpected token '%s' on position %d", p.current().TType, p.currentPosition)
	}
	return
}
