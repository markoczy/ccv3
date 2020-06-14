package parser

type TokenType int

const (
	NumericToken = TokenType(iota)
	IdentifierToken
	OperatorToken
	ControlToken
	EndToken
)
