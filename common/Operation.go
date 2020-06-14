package common

type Operation string

const (
	Addition       = Operation("+")
	Subtraction    = Operation("-")
	Multiplication = Operation("*")
	Division       = Operation("/")
)
