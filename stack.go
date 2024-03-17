package main

type stack[T any] struct {
	data []T
}

func (s *stack[T]) Push(x T) {
	s.data = append(s.data, x)
}
func (s *stack[T]) Pop() T {
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}
func (s *stack[T]) Size() int {
	return len(s.data)
}
func (s *stack[T]) Top() T {
	return s.data[len(s.data)-1]
}
