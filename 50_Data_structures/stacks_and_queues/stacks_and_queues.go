package stacksandqueues

import "fmt"

// Stack represents a stack that hold a slice
type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(a int) {
	s.items = append(s.items, a)
}

// Pop will remove a value at the end
// and RETURNs the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	if l < 0 {
		fmt.Println("Can't pop item. Stack is empty")
		return 0
	}
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(a int) {
	q.items = append(q.items, a)
}

// Dequeue remove a value at the front
// and RETURNs the remove value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
