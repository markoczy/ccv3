package fsm

import (
	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
)

type StateParams struct {
	Tokens *parser.TokenQueue
	Cur    *common.EquationNode
}

func NewStateParams(tokens *parser.TokenQueue, eq *common.EquationNode) *StateParams {
	return &StateParams{
		Tokens: tokens,
		Cur:    eq,
	}
}
