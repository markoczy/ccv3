package tokenizer

type RuneQueue []rune

func (q *RuneQueue) Len() int {
	return len(*q)
}

func (q *RuneQueue) Enqueue(r rune) {
	*q = append(*q, r)
}

func (q *RuneQueue) Peek() rune {
	if len(*q) == 0 {
		return 0
	}
	return (*q)[0]
}

func (q *RuneQueue) Dequeue() rune {
	if len(*q) == 0 {
		return 0
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}
