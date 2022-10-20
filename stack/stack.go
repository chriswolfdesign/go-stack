package stack

import "fmt"

type Stack struct {
	size int
	head *stackNode
	tail *stackNode
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(element int) {
	newNode := &stackNode{val: element}
	if s.size == 0 {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.next = newNode
		newNode.prev = s.tail
		s.tail = newNode
	}

	s.size++
}

func (s *Stack) Peek() (int, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("Not implemented")
	}

	return s.tail.val, nil
}

type stackNode struct {
	val  int
	next *stackNode
	prev *stackNode
}
