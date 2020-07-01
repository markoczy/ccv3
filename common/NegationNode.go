package common

type NegationNode struct {
	Val Node
}

func (n *NegationNode) Value() (float64, error) {
	ret, err := n.Val.Value()
	return -ret, err
}

func (n *NegationNode) String() string {
	return "-" + n.Val.String()
}

func (n *NegationNode) Type() NodeType {
	return NegationNodeType
}

func (n *NegationNode) Clone() Node {
	return NewNegationNode(n.Val.Clone())
}

func NewNegationNode(val Node) *NegationNode {
	return &NegationNode{
		Val: val,
	}
}
