package common

import (
	"log"
	"math"
	"strings"
)

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

func (n *EquationNode) Value() (float64, error) {
	n.solveAll(Exponentiation)
	n.solveAll(Division)
	n.solveAll(Multiplication)
	n.solveAll(Subtraction)
	n.solveAll(Addition)

	ret, err := n.children[0].Value()
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (n *EquationNode) solveAll(operation Operation) error {
	hasMore := true
	for hasMore {
		hasMore = false
		for i, op := range n.operations {
			if op == operation {
				err := n.solveOperation(i)
				if err != nil {
					return err
				}
				hasMore = true
				break
			}
		}
	}
	return nil
}

func (n *EquationNode) solveOperation(idx int) error {
	left, err := n.children[idx].Value()
	if err != nil {
		return err
	}
	right, err := n.children[idx+1].Value()
	if err != nil {
		return err
	}
	op := n.operations[idx]

	log.Printf("Solving: %f %s %f\n", left, op, right)

	var res ValueNode
	switch op {
	case Addition:
		res = ValueNode(left + right)
	case Subtraction:
		res = ValueNode(left - right)
	case Multiplication:
		res = ValueNode(left * right)
	case Division:
		res = ValueNode(left / right)
	case Exponentiation:
		res = ValueNode(math.Pow(left, right))
	}

	log.Printf("Result: %f %s %f = %f\n", left, op, right, float64(res))

	log.Println("Operations before:", len(n.operations))
	n.operations = append(n.operations[:idx], n.operations[idx+1:]...)
	log.Println("Operations after:", len(n.operations))
	log.Println("Children before:", len(n.children))
	rest := n.children[idx+2:]
	n.children = append(n.children[:idx], &res)
	n.children = append(n.children, rest...)
	log.Println("Children after:", len(n.children))
	return nil
}

func (n *EquationNode) String() string {
	ret := strings.Builder{}
	for i, v := range n.operations {
		if i == 0 {
			writeNode(n.children[i], &ret)
		}
		ret.WriteString(string(v))
		writeNode(n.children[i+1], &ret)
	}
	return ret.String()
}

func (n *EquationNode) Type() NodeType {
	return EquationNodeType
}

func NewEquationNode() *EquationNode {
	return &EquationNode{}
}

func writeNode(node Node, sb *strings.Builder) {
	if node.Type() == EquationNodeType {
		sb.WriteRune('(')
	}
	sb.WriteString(node.String())
	if node.Type() == EquationNodeType {
		sb.WriteRune(')')
	}
}
