package stack

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) interface{} {
	s.items = append(s.items, item)
	return item
}

func (s *Stack) Pop() interface{} {
	if s.Len() > 0 {
		n := len(s.items) - 1
		item := s.items[n]
		s.items = s.items[:n]
		return item
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.Len() > 0 {
		return s.items[len(s.items)-1]
	}
	return nil
}

func (s *Stack) Len() int {
	return len(s.items)
}
