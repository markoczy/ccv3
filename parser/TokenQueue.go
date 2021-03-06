package parser

import (
	"fmt"
	"strings"
)

type TokenQueue []*Token

func (q *TokenQueue) Len() int {
	return len(*q)
}

func (q *TokenQueue) Enqueue(t *Token) {
	*q = append(*q, t)
}

func (q *TokenQueue) Peek() *Token {
	if len(*q) == 0 {
		return &TokenEnd
	}
	return (*q)[0]
}

func (q *TokenQueue) Dequeue() *Token {
	if len(*q) == 0 {
		return &TokenEnd
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *TokenQueue) String() string {
	strs := make([]string, len(*q))
	for i, v := range *q {
		strs[i] = fmt.Sprintf("%v", *v)
	}
	return strings.Join(strs, ", ")
}
