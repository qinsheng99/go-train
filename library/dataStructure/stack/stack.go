package stack

// Stack 栈
type Stack struct {
	size int
	top  int
	data []interface{}
}

func CreateStack(size int) *Stack {
	return &Stack{
		size: size,
		top:  -1,
		data: make([]interface{}, size),
	}
}

func (s *Stack) Push(data interface{}) {
	if s.IsFull() {
		s.dilatation()
	}

	s.top++
	s.data[s.top] = data
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	tmp := s.data[s.top]
	s.top--
	return tmp
}

func (s *Stack) Clear() {
	s.top = -1
	s.data = make([]interface{}, s.size)
}

func (s *Stack) dilatation() {
	size := s.size * 2

	var res = make([]interface{}, size)

	copy(res, s.data)

	s.data = res
	s.size = size
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack) GetData() interface{} {
	return s.data
}

func (s *Stack) GetTop() int {
	return s.top
}
