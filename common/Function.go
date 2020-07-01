package common

const UnspecifiedParamCount = -1

type Function interface {
	Parametrized() bool
	Validate(...Node) error
	AsNode(...Node) (Node, error)
}
