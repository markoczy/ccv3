package fsm

import (
	"fmt"

	"github.com/markoczy/ccv3/parser"
)

var EndState ParserState = &endState{}

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

func NewParserState(exec func(*CallStack) (int, error)) ParserState {
	return &defaultParserState{
		exec: exec,
	}
}

func NewParserStateWithExecAndTransitionFunc(exec func(*CallStack) error, transition func(*CallStack) (int, error)) ParserState {
	return &defaultParserState{
		exec: func(stack *CallStack) (ret int, err error) {
			if err = exec(stack); err != nil {
				return
			}
			ret, err = transition(stack)
			return
		},
	}
}

func NewParserStateWithExecAndTransitionMap(exec func(*CallStack) error, transitions map[parser.TokenType]int) ParserState {
	var transition = TokenMapTransition(transitions)
	return &defaultParserState{
		exec: func(stack *CallStack) (ret int, err error) {
			if err = exec(stack); err != nil {
				return
			}
			ret, err = transition(stack)
			return
		}}
}
