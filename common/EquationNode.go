package common

import "strings"

type EquationNode struct {
	operations []Operation
	children   []Node
}

func (n *EquationNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

func (n *EquationNode) AddOperation(op Operation) {
	n.operations = append(n.operations, op)
}

func (n *EquationNode) Value() float64 {
	return 42 // todo
}

func (n *EquationNode) String() string {
	ret := strings.Builder{}
	for i, v := range n.operations {
		if i == 0 {
			ret.WriteString(n.children[i].String())
		}
		ret.WriteString(string(v))
		ret.WriteString(n.children[i+1].String())
	}
	return ret.String()
}
