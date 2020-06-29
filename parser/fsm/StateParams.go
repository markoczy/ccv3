package fsm

import (
	"github.com/markoczy/ccv3/common"
)

type StateParams struct {
	Equation *common.EquationNode
	Negate   bool
}

func NewStateParams(equation *common.EquationNode) *StateParams {
	return &StateParams{
		Equation: equation,
		Negate:   false,
	}
}
