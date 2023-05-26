package main

import "testing"

func TestReverseString(t *testing.T) {
	t.Run("Reverse a string func 0", func(t *testing.T) {

		str := "hello"
		got := ReverseString_0(str)
		want := "olleh"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Reverse string func 1", func(t *testing.T) {
		str := "hello"
		got := ReverseString_1(str)
		want := "olleh"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

}
