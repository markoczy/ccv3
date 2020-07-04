package common

import "fmt"

type FunctionMap map[string]Function

func (fm *FunctionMap) Parametrized(id string) bool {
	fn := (*fm)[id]
	if fn == nil {
		return false
	}
	return fn.Parametrized()
}

func (fm *FunctionMap) Validate(id string, params ...Node) error {
	fn := (*fm)[id]
	if fn == nil {
		return fmt.Errorf("Undefined function or parameter '%s'", id)
	}
	return fn.Validate()
}

func (fm *FunctionMap) Call(id string, params ...Node) (Node, error) {
	fn := (*fm)[id]
	if fn == nil {
		return nil, fmt.Errorf("Undefined function or parameter '%s'", id)
	}
	return fn.AsNode(params...)
}

func NewFunctionMap(functionMap map[string]Function) *FunctionMap {
	if functionMap == nil {
		functionMap = map[string]Function{}
	}
	ret := FunctionMap(functionMap)
	return &ret
}
