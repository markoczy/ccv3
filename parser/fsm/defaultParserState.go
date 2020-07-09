package fsm

type defaultParserState struct {
	exec func(*CallStack) (int, error)
}

func (state *defaultParserState) End() bool {
	return false
}

func (state *defaultParserState) Exec(stack *CallStack) (int, error) {
	return state.exec(stack)
}
