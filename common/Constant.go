package common

type Constant struct {
	identifier string
	value      float64
}

func (c *Constant) Parametrized() bool {
	return false
}

func (c *Constant) Validate(params ...Node) bool {
	return len(params) == 0
}

func (c *Constant) AsNode(params ...Node) Node {
	return NewValueNode(c.value)
}

func (c *Constant) String() string {
	return c.identifier
}
