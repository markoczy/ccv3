package common

import "strconv"

type ValueNode float64

func (n *ValueNode) Value() (float64, error) {
	return float64(*n), nil
}

func (n *ValueNode) String() string {
	return strconv.FormatFloat(float64(*n), 'f', -1, 64)
}

func (n *ValueNode) Type() NodeType {
	return ValueNodeType
}

func (n *ValueNode) Clone() Node {
	val := float64(*n)
	return NewValueNode(val)
}

func NewValueNode(val float64) *ValueNode {
	ret := ValueNode(val)
	return &ret
}
