package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(w interface{}) {
	s.v.PushBack(w)
}
func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil

}
func main() {
	stack := NewStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}
	v := stack.Pop()
	for v != nil {
		fmt.Println(v)
		v = stack.Pop()
	}
}
