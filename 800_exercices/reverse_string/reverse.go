package main

import (
	"fmt"
)

func ReverseString_0(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func ReverseString_1(s string) string {
	r := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)

}

func main() {
	fmt.Println("------ reverse strings ------")

	fmt.Println(ReverseString_0("123456789"))
	fmt.Println(ReverseString_1("123456789"))
}
