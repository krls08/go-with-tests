package main

import (
	"fmt"
	"runtime"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {
	// Below is an example of using our PrintMemUsage() function
	// Print our starting memory usage (should be around 0mb)
	PrintMemUsage()

	mM := make(map[string]person)
	PrintMemUsage()

	mM["zero"] = person{"krls", 123}
	PrintMemUsage()
	mM["onee"] = person{"krls", 123}

	// Allocate memory using make() and append to overall (so it doesn't get
	// garbage collected). This is to create an ever increasing memory usage
	// which we can track. We're just using []int as an example.

	// Print our memory usage at each interval
	PrintMemUsage()
	time.Sleep(time.Second)
	//mM = nil
	delete(mM, "zero")
	delete(mM, "onee")

	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	//overall = nil

	PrintMemUsage()
	//time.Sleep(10 * time.Second)

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	PrintMemUsage()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v B", m.Alloc)
	fmt.Printf("\tFrees = %v B", m.Frees)
	fmt.Printf("\tHeapSys= %v B", m.HeapSys)
	fmt.Printf("\tTotalAlloc = %v B", m.TotalAlloc)
	fmt.Printf("\tSys = %v B", m.Sys)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
