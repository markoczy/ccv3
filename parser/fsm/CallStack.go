package fsm

import (
	"fmt"

	"github.com/markoczy/ccv3/common"
)

type CallStack []*StateParams

func (s *CallStack) Len() int {
	return len(*s)
}

func (s *CallStack) Push(params *StateParams) {
	*s = append(*s, params)
}

func (s *CallStack) Cur() *StateParams {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *CallStack) Pop() *StateParams {
	if len(*s) == 0 {
		return nil
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

func (s *CallStack) Enter() error {
	if len(*s) == 0 {
		return fmt.Errorf("Cannot call Enter when stack is empty")
	}
	tokens := s.Cur().Tokens
	child := common.NewEquationNode()
	s.Cur().Cur.AddChild(child)
	*s = append(*s, NewStateParams(tokens, child))
	return nil
}

func (s *CallStack) Exit() error {
	if len(*s) == 0 {
		return fmt.Errorf("Cannot call Exit when stack is empty")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}
