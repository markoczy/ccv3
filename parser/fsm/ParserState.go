package fsm

type ParserState struct {
	End        bool
	Func       func(*CallStack) error
	Transition func(*CallStack) (int, error)
}

func NewParserState(end bool, exec func(*CallStack) error, transition func(*CallStack) (int, error)) *ParserState {
	return &ParserState{
		End:        end,
		Func:       exec,
		Transition: transition,
	}
}
