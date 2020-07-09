package fsm

import "fmt"

type endState struct{}

func (state *endState) End() bool {
	return true
}

func (state *endState) Exec(*CallStack) (int, error) {
	return 0, fmt.Errorf("Called Exec() on End State")
}
