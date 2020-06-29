package common

type NodeStack []*EquationNode

func (s *NodeStack) Len() int {
	return len(*s)
}

func (s *NodeStack) Push(node *EquationNode) {
	*s = append(*s, node)
}

func (s *NodeStack) Peek() *EquationNode {
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *NodeStack) Pop() *EquationNode {
	if len(*s) == 0 {
		return nil
	}
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}
