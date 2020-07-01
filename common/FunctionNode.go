package common

import "strings"

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

func (node *FunctionNode) Clone() Node {
	params := make([]Node, len(node.params))
	for i := range node.params {
		params[i] = node.params[i].Clone()
	}
	return NewFunctionNode(node.identifier, node.function, params)

}

func NewFunctionNode(identifier string, function func(...float64) float64,
	params []Node) *FunctionNode {
	return &FunctionNode{
		identifier: identifier,
		function:   function,
		params:     params,
	}

}
