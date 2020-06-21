package parser

type TokenType string

const (
	NumericToken    = TokenType("Number")
	IdentifierToken = TokenType("Identifier")
	OperatorToken   = TokenType("Operator")
	ControlToken    = TokenType("Control")
	EndToken        = TokenType("End")
)
