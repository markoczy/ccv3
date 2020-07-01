package common

import "fmt"

// DefaultFunction implements Function, it is the default implementation
// of a function with an either fixed or infinite length of params.
type DefaultFunction struct {
	identifier string
	paramCount int
	function   func(...float64) float64
}

func (fn *DefaultFunction) Parametrized() bool {
	return fn.paramCount != 0
}

func (fn *DefaultFunction) Validate(nodes ...Node) error {
	if fn.paramCount != UnspecifiedParamCount && len(nodes) != fn.paramCount {
		return fmt.Errorf("Wrong parameter count want %d, have %d", fn.paramCount, len(nodes))
	}
	return nil
}

func (fn *DefaultFunction) AsNode(nodes ...Node) (Node, error) {
	if err := fn.Validate(nodes...); err != nil {
		return nil, err
	}

	return &FunctionNode{
		identifier: fn.identifier,
		function:   fn.function,
		params:     nodes,
	}, nil
}

func NewDefaultFunction(identifier string, paramCount int, function func(...float64) float64) Function {
	return &DefaultFunction{
		identifier: identifier,
		paramCount: paramCount,
		function:   function,
	}
}
