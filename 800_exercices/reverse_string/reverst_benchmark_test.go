// go test -bench=. --benchmem
package main

import "testing"

const iterations = 1000 * 10

const string_0 = "abcdefgh"
const string_1 = "abcdefghijklmnopqrstuvwxyz"

func BenchmarkReverseString_0(b *testing.B) {
	//	b.N = iterations
	for i := 0; i < b.N; i++ {
		ReverseString_0(string_0)
	}
}

func BenchmarkReverseString_1(b *testing.B) {
	//	b.N = iterations
	for i := 0; i < b.N; i++ {
		ReverseString_1(string_0)
	}

}
