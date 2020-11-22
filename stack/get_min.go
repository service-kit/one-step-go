package stack

type GetMinStack struct {
	stack    *Stack
	minStack *Stack
}

func (s *GetMinStack) Init() {
	s.stack = GetStack()
	s.minStack = GetStack()
}

func (s *GetMinStack) Pop() (int, error) {
	data, err := s.stack.pop()
	if err != nil {
		return -1, err
	}
	s.minStack.pop()
	return data.(int), nil
}

func (s *GetMinStack) Push(data int) {
	t, err := s.minStack.top()
	if err != nil {
		s.minStack.push(data)
	} else {
		curMin := t.(int)
		minPutData := data
		if data > curMin {
			minPutData = curMin
		}
		s.minStack.push(minPutData)
	}
	s.stack.push(data)
}

func (s *GetMinStack) GetMin() (int, error) {
	min, err := s.minStack.top()
	if err == nil {
		return min.(int), nil
	}
	return -1, err
}
