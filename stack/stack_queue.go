package stack

import "errors"

type StackQueue struct {
	stackIn  *Stack
	stackOut *Stack
}

func (q *StackQueue) Init() {
	q.stackIn = GetStack()
	q.stackOut = GetStack()
}

func (q *StackQueue) Add(data int) {
	q.stackIn.push(data)
}

func (q *StackQueue) Poll() (int, error) {
	if q.stackOut.len() == 0 {
		if q.stackIn.len() == 0 {
			return -1, errors.New("empty")
		} else {
			for {
				if val, err := q.stackIn.pop(); err == nil {
					q.stackOut.push(val)
				} else {
					break
				}
			}
		}
	}
	out, err := q.stackOut.pop()
	if err == nil {
		return out.(int), nil
	}
	return -1, err
}

func (q *StackQueue) Peek() (int, error) {
	if q.stackOut.len() == 0 {
		if q.stackIn.len() == 0 {
			return -1, errors.New("empty")
		} else {
			for {
				if val, err := q.stackIn.pop(); err == nil {
					q.stackOut.push(val)
				} else {
					break
				}
			}
		}
	}
	out, err := q.stackOut.top()
	if err == nil {
		return out.(int), nil
	}
	return -1, err
}
