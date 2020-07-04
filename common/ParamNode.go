package common

import "fmt"

type ParamNode struct {
	identifier string
	node       Node
}

func NewParamNode(identifier string) *ParamNode {
	return &ParamNode{
		identifier: identifier,
	}
}

//* Node interface:

func (n *ParamNode) Value() (float64, error) {
	return 0, fmt.Errorf("Unresolved parameter '%s'", n.identifier)
}

func (n *ParamNode) String() string {
	return n.identifier
}

func (n *ParamNode) Type() NodeType {
	return ParamnNodeType
}

func (n *ParamNode) Clone() Node {
	return &ParamNode{
		identifier: n.identifier,
	}
}

//* Function interface:

func (n *ParamNode) Parametrized() bool {
	return false
}

func (n *ParamNode) Validate(params ...Node) error {
	if len(params) != 0 {
		return fmt.Errorf("Expected no arguments but got: %v", params)
	}
	return nil
}

func (n *ParamNode) AsNode(params ...Node) Node {
	return n
}
