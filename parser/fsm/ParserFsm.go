package fsm

import (
	"fmt"

	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
	"github.com/markoczy/ccv3/parser/tokenizer"
)

type ParserFsm struct {
	States map[int]ParserState
	State  ParserState
	Stack  *CallStack
}

func (fsm *ParserFsm) End() bool {
	return fsm.State.End()
}

func (fsm *ParserFsm) Step() error {
	if fsm.State.End() {
		return nil
	}

	stack := fsm.Stack
	if stack.Tokens.Len() == 0 {
		return fmt.Errorf("No more tokens and not at end state")
	}

	next, err := fsm.State.Exec(fsm.Stack)
	if err != nil {
		return fsm.formatParserError(err, stack.Tokens.Peek())
	}
	fsm.State = fsm.States[next]
	return nil
}

func (fsm *ParserFsm) Parse(s string, functionMap *common.FunctionMap) (*common.EquationNode, error) {
	tokens, err := tokenizer.CreateTokens(s)
	if err != nil {
		return nil, err
	}

	fsm.Stack = NewCallStack(&tokens, functionMap)

	// &CallStack{Params:NewStateParams(&tokens, common.NewEquationNode())}
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

func NewParserFsm(states map[int]ParserState, start int) *ParserFsm {
	return &ParserFsm{
		States: states,
		State:  states[start],
	}
}
