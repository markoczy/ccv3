package fsm

import (
	"fmt"

	"github.com/markoczy/ccv3/parser"
)

var EndState = &ParserState{
	End:        true,
	Func:       NoOp,
	Transition: NoTransition,
}

func NoOp(*CallStack) error {
	return nil
}

func NoTransition(s *CallStack) (int, error) {
	return 0, fmt.Errorf("Tried to get transition from empty")
}

func TokenMapTransition(transitions map[parser.TokenType]int) func(s *CallStack) (int, error) {
	return func(s *CallStack) (int, error) {
		token := s.Tokens.Peek()
		next, found := transitions[token.Type]
		if !found {
			return 0, fmt.Errorf("Undefined successor state")
		}
		return next, nil
	}
}
