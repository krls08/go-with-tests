// this package contais some benchmarks related to how go assign in arrays / slices
// using the method append vs assign directly to the position
// to laungh the benchmarks run from this directory:
// go test -bench . --benchmem
package benchmark

import "testing"

var result []int

const size = 32

const iterations = 100 * 1000 * 1000

func doAssign() {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	result = data
}

func doAppend() {
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, i)
	}
	result = data
}

func arrAssign() {
	data := [iterations]int{}
	for i := 0; i < size; i++ {
		data[i] = i
	}
}

func BenchmarkArrAssign(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		doAppend()
	}
}
func BenchmarkAssign(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		doAssign()
	}
}

func BenchmarkAppend(b *testing.B) {
	b.N = iterations
	for i := 0; i < b.N; i++ {
		doAppend()
	}
}
