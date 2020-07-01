package common

type placeholder struct{}

type EquationFunction struct {
	identifier string
	paramMap   map[string]int
}

func (fn *EquationFunction) Parametrized() bool {
	return len(fn.paramMap) == 0
}

func (fn *EquationFunction) Validate(params ...Node) bool {
	return len(params) == len(fn.paramMap)
}

func (fn *EquationFunction) AsNode(params ...Node) Node {
	return nil
}
