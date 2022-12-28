package kit

import "testing"

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}
