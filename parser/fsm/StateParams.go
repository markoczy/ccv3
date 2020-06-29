package fsm

import (
	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
)

type StateParams struct {
	Tokens   *parser.TokenQueue
	Equation *common.EquationNode
	Negate   bool
}

func NewStateParams(tokens *parser.TokenQueue, equation *common.EquationNode) *StateParams {
	return &StateParams{
		Tokens:   tokens,
		Equation: equation,
		Negate:   false,
	}
}
