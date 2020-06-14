package common

import "strconv"

type ValueNode float64

func (n *ValueNode) Value() float64 {
	return float64(*n)
}

func (n *ValueNode) String() string {
	return strconv.FormatFloat(float64(*n), 'f', -1, 64)
}
