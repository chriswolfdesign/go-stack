package stack

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

type stackNode struct {
	val  int
	next *stackNode
	prev *stackNode
}
