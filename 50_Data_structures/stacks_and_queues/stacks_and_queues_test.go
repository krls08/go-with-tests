package stacksandqueues

import (
	"fmt"
	"testing"
)

func TestStackMethods(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		s := Stack{}
		assertSize(t, s, 0)
		s.Push(1)
		assertSize(t, s, 1)
	})

	t.Run("Pop", func(t *testing.T) {

	})
}

func TestQueueMethods(t *testing.T) {
	t.Run("Enqueue", func(t *testing.T) {
		q := Queue{}
		q.Enqueue(1)
		q.Enqueue(2)
		assertEqual(t,
			fmt.Sprint([]int{1, 2}), fmt.Sprint(q.items))

	})

	t.Run("Dequeue", func(t *testing.T) {

	})
}

func assertSize(t testing.TB, s Stack, want int) {
	t.Helper()
	//fmt.Println("s.items", len(s.items))
	if len(s.items) != want {
		t.Errorf("Wanted size: %d, got: %d", want, len(s.items))
	}
}

func assertEqual[T comparable](t testing.TB, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("wanted: %v, got: %v", expected, actual)
	}
}
