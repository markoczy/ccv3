package fsm

import (
	"github.com/markoczy/ccv3/common"
	"github.com/markoczy/ccv3/parser"
)

type StateParams struct {
	Tokens parser.TokenQueue
	Cur    common.EquationNode
	Stack  []common.EquationNode
}
