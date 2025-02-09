package stack

type Stack[T any] struct {
	Data []T
	Top  int
}

func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		Data: make([]T, size),
		Top:  0,
	}
}

func (s *Stack[T]) Push(v T) {
	s.Data = append(s.Data, v)
	s.Top++
}

func (s *Stack[T]) Pop() T {
	if s.Top == 0 {
		return s.Data[s.Top]
	}

	v := s.Data[s.Top-1]
	s.Data = s.Data[:s.Top-1]
	s.Top--
	return v
}
