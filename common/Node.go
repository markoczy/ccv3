package common

type Node interface {
	Value() (float64, error)
	String() string
}
