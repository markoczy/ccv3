package tokenizer

import (
	"fmt"
	"strings"

	"github.com/markoczy/ccv3/parser"
)

func CreateToken(q *RuneQueue, tp parser.TokenType, begin int, eval func(rune) bool) parser.Token {
	ret := parser.Token{
		Type:  tp,
		Begin: begin,
	}

	sb := strings.Builder{}
	cnt := 0
	for eval(q.Peek()) {
		sb.WriteRune(q.Dequeue())
		cnt++
	}

	ret.Value = sb.String()
	ret.End = begin + (cnt - 1)
	return ret
}

func CreateSingleRuneToken(q *RuneQueue, tp parser.TokenType, begin int, eval func(rune) bool) parser.Token {
	ret := parser.Token{
		Type:  tp,
		Begin: begin,
	}

	if eval(q.Peek()) {
		ret.Value = string(q.Dequeue())
	}

	ret.End = begin
	return ret
}

func CreateTokens(s string) (parser.TokenQueue, error) {
	queue := RuneQueue(s)
	ret := parser.TokenQueue{}
	initialLength := queue.Len()
	for queue.Len() > 0 {
		begin := initialLength - queue.Len()
		r := queue.Peek()
		var token parser.Token
		switch {
		case parser.IsNumeric(r):
			token = CreateToken(&queue, parser.NumericToken, begin, parser.IsNumeric)
			ret.Enqueue(&token)
		case parser.IsIdentifier(r):
			token = CreateToken(&queue, parser.IdentifierToken, begin, parser.IsIdentifier)
			ret.Enqueue(&token)
		case parser.IsOperator(r):
			token = CreateSingleRuneToken(&queue, parser.OperatorToken, begin, parser.IsOperator)
			ret.Enqueue(&token)
		case parser.IsControl(r):
			token = CreateSingleRuneToken(&queue, parser.ControlToken, begin, parser.IsControl)
			ret.Enqueue(&token)
		default:
			return nil, fmt.Errorf("Unhandled char '%s' at position %d", string(r), begin)
		}
	}
	return ret, nil
}
