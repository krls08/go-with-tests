package main

import (
	"fmt"
	"time"
)

func main() {
	var td time.Duration
	fmt.Println("Default time Duration:", td)
	td = time.Minute*10 + time.Hour*300
	fmt.Println("td + 10 Minutes + 300 Hours:", td)
	fmt.Println("")

	t, _ := time.Parse(time.RFC3339, "2020-01-29T10:00:00Z")
	fmt.Println("t =>", t)

	nt := t.Add(td)
	fmt.Println("t =>", nt)

}
