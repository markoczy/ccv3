package fsm

import (
	"fmt"

	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
	"github.com/markoczy/ccv3/parser/tokenizer"
)

type ParserFsm struct {
	States map[int]*ParserState
	State  *ParserState
	Stack  *CallStack
}

func (fsm *ParserFsm) End() bool {
	return fsm.State.End
}

func (fsm *ParserFsm) Step() error {
	if fsm.State.End {
		return nil
	}

	state := fsm.Stack.Cur()
	if state.Tokens.Len() == 0 {
		return fmt.Errorf("No more tokens and not at end state")
	}

	token := state.Tokens.Peek()
	err := fsm.State.Func(fsm.Stack)
	if err != nil {
		return fsm.formatParserError(err, token)
	}

	next, err := fsm.State.Transition(fsm.Stack)
	if err != nil {
		return fsm.formatParserError(err, state.Tokens.Peek())
	}
	fsm.State = fsm.States[next]
	return nil
}

func (fsm *ParserFsm) Parse(s string) (*common.EquationNode, error) {
	tokens, err := tokenizer.CreateTokens(s)
	if err != nil {
		return nil, err
	}

	fsm.Stack = &CallStack{NewStateParams(&tokens, common.NewEquationNode())}
	for !fsm.End() {
		err = fsm.Step()
		if err != nil {
			return nil, err
		}
	}

	return fsm.Stack.Cur().Equation, nil
}

func (fsm *ParserFsm) formatParserError(err error, token *parser.Token) error {
	if token.Begin == token.End {
		return fmt.Errorf("Parser Error: %w (token '%s', position %d)", err, token.Value, token.Begin)

	}
	return fmt.Errorf("Parser Error: %w (token '%s', position %d-%d)", err, token.Value, token.Begin, token.End)
}

func NewParserFsm(states map[int]*ParserState, start int) *ParserFsm {
	return &ParserFsm{
		States: states,
		State:  states[start],
		Stack:  &CallStack{&StateParams{}},
	}
}
