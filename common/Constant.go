package common

import "fmt"

type Constant struct {
	identifier string
	value      float64
}

func NewConstant(identifier string, value float64) *Constant {
	return &Constant{
		identifier: identifier,
		value:      value,
	}
}

func (c *Constant) Parametrized() bool {
	return false
}

func (c *Constant) Validate(params ...Node) error {
	if len(params) != 0 {
		return fmt.Errorf("Expected no arguments but got: %v", params)
	}
	return nil
}

func (c *Constant) AsNode(params ...Node) Node {
	return NewValueNode(c.value)
}

func (c *Constant) String() string {
	return c.identifier
}
