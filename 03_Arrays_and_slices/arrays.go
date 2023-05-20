package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)

	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func SliceRemoveIndex(aSlice []string, index int) []string {
	aSlice[index] = aSlice[len(aSlice)-1]
	return aSlice[:len(aSlice)-1]
}

func GenSliceRemoveIndex[K interface{}](s []K, i int) []K {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//func remove(s []int, i int) []int {
//    s[i] = s[len(s)-1]
//    return s[:len(s)-1]
//}

func CreateStringSlice() []string {
	s := make([]string, 0, 10)
	fmt.Println("s", s)
	return s
}
