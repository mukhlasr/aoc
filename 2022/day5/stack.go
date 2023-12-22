package main

type Stack struct {
	Item []byte
	Size int
}

func (s *Stack) Pop() byte {
	res := s.Top()
	s.Item = s.Item[1:]
	s.Size--
	return res
}

func (s *Stack) Push(i byte) {
	s.Item = append([]byte{i}, s.Item...)
	s.Size++
}

func (s *Stack) Top() byte {
	return s.Item[0]
}

func (s *Stack) PushBack(i byte) {
	s.Item = append(s.Item, i)
	s.Size++
}
