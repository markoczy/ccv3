package parser

type Token struct {
	Type  TokenType
	Value string
	// For debugging:
	Begin int
	End   int
}

var TokenEnd = Token{
	Type: EndToken,
}
