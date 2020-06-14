package fsm

import (
	"fmt"

	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser/tokenizer"
)

type ParserFsm struct {
	States map[int]ParserState
	State  ParserState
	Params StateParams
}

func (fsm *ParserFsm) End() bool {
	return fsm.State.End
}

func (fsm *ParserFsm) Step() error {
	if fsm.State.End {
		return nil
	}

	if fsm.Params.Tokens.Len() == 0 {
		return fmt.Errorf("No more tokens and not at end state")
	}

	err := fsm.State.Func(&fsm.Params)
	if err != nil {
		return err
	}

	token := fsm.Params.Tokens.Peek()
	next, found := fsm.State.Transitions[token.Type]
	if !found {
		return fmt.Errorf("Undefined successor state for token: %v", token)
	}

	fsm.State = fsm.States[next]
	return nil
}

func (fsm *ParserFsm) Parse(s string) (*common.EquationNode, error) {
	tokens, err := tokenizer.CreateTokens(s)
	if err != nil {
		return nil, err
	}

	fsm.Params = StateParams{
		Tokens: tokens,
		Cur:    common.EquationNode{},
		Stack:  []common.EquationNode{},
	}

	for !fsm.End() {
		err = fsm.Step()
		if err != nil {
			return nil, err
		}
	}

	return &fsm.Params.Cur, nil
}

func NewParserFsm(states map[int]ParserState, start int) *ParserFsm {
	return &ParserFsm{
		States: states,
		State:  states[start],
		Params: StateParams{},
	}
}
