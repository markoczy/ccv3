package common

import (
	"fmt"
	"strings"
)

const UnspecifiedParamCount = -1

type Function interface {
	Validate(...Node) error
	AsNode(...Node) (Node, error)
}

type FunctionNode struct {
	identifier string
	function   func(...float64) float64
	params     []Node
}

func (node *FunctionNode) Value() (float64, error) {
	params := []float64{}
	for _, v := range node.params {
		val, err := v.Value()
		if err != nil {
			return 0, err
		}
		params = append(params, val)
	}
	return node.function(params...), nil
}

func (node *FunctionNode) String() string {
	if len(node.params) == 0 {
		return node.identifier
	}
	params := []string{}
	for _, v := range node.params {
		params = append(params, v.String())
	}
	return node.identifier + "(" + strings.Join(params, ", ") + ")"
}

func (node *FunctionNode) Type() NodeType {
	return FunctionNodeType
}

type DefaultFunction struct {
	identifier string
	paramCount int
	function   func(...float64) float64
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
