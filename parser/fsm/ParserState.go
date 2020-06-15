package fsm

import "github.com/markoczy/ccv3/parser"

type ParserState struct {
	End         bool
	Func        func(*StateParams) error
	Transitions map[parser.TokenType]int
}

var EndState = ParserState{
	End:         true,
	Func:        NoOp,
	Transitions: map[parser.TokenType]int{},
}

func NewParserState(end bool, fn func(*StateParams) error, transitions map[parser.TokenType]int) *ParserState {
	return &ParserState{
		End:         end,
		Func:        fn,
		Transitions: transitions,
	}
}
