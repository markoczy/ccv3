package fsm

import (
	"fmt"
	"log"

	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
)

type CallStack struct {
	Params []*StateParams
	Tokens *parser.TokenQueue
}

// type CallStack []*StateParams

func (s *CallStack) Len() int {
	return len(s.Params)
}

func (s *CallStack) Push(params *StateParams) {
	s.Params = append(s.Params, params)
}

func (s *CallStack) Cur() *StateParams {
	if len(s.Params) == 0 {
		return nil
	}
	return s.Params[len(s.Params)-1]
}

func (s *CallStack) Pop() *StateParams {
	if len(s.Params) == 0 {
		return nil
	}
	ret := (s.Params)[len(s.Params)-1]
	s.Params = s.Params[:len(s.Params)-1]
	return ret
}

func (s *CallStack) Enter() error {
	if len(s.Params) == 0 {
		return fmt.Errorf("Cannot call Enter when stack is empty")
	}
	var child common.Node
	eq := common.NewEquationNode()
	if s.Cur().Negate {
		s.Cur().Negate = false
		child = &common.NegationNode{
			Val: eq,
		}
	} else {
		child = eq
	}
	s.Cur().Equation.AddChild(child)
	s.Params = append(s.Params, NewStateParams(eq))
	log.Println("Length", len(s.Params))
	return nil
}

func (s *CallStack) Exit() error {
	if len(s.Params) == 0 {
		return fmt.Errorf("Cannot call Exit when stack is empty")
	}
	s.Params = s.Params[:len(s.Params)-1]
	return nil
}

func NewCallStack(tokens *parser.TokenQueue) *CallStack {
	return &CallStack{
		Params: []*StateParams{
			NewStateParams(common.NewEquationNode()),
		},
		Tokens: tokens,
	}
}
