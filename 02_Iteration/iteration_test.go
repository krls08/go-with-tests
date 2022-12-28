package iteration

import (
	"fmt"
	"testing"

	"github.com/krls08/go-with-tests/kit"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	kit.AssertCorrectMessage(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("n", 3)
	fmt.Println(repeated)
	// Output: nnn
}
