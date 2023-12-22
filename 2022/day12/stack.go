package main

import "fmt"

type Stack struct {
	Item []position
	Size int
}

func (s *Stack) Pop() position {
	res := s.Top()
	s.Item = s.Item[1:]
	s.Size--
	return res
}

func (s *Stack) Push(i position) {
	s.Item = append([]position{i}, s.Item...)
	s.Size++
}

func (s Stack) Top() position {
	return s.Item[0]
}

func (s Stack) String() string {
	str := ""
	for _, i := range s.Item {
		str += fmt.Sprintf("[%d, %d]", i.x, i.y)
	}
	return str
}
