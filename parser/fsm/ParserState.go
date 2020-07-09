package fsm

type ParserState interface {
	End() bool
	Exec(*CallStack) (int, error)
}
