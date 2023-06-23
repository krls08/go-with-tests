package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	fmt.Println("Example: 1, 2, 3, 4, 5")
	ex0 := []int32{1, 2, 3, 4, 5}
	miniMaxSum(os.Stdout, ex0)
	fmt.Println("max ->", math.MaxInt32)

}

func miniMaxSum(writer io.Writer, arr []int32) {
	var min, max int32
	var sum int64
	var first bool = true
	for _, val := range arr {
		if first {
			first = false
			min, max = val, val
		} else {
			if val < min {
				min = val
			}
			if val > max {
				max = val

			}
		}
		sum += int64(val)
	}

	writer.Write([]byte(fmt.Sprintf("%d %d", sum-int64(max), sum-int64(min))))

}
