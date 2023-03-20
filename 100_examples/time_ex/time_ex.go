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

	fmt.Println("-----------------")
	tnow := time.Now()
	fmt.Println("time now:", tnow)
	nowPlus8 := time.Now().Add(8 * time.Minute)
	nowMinus8 := time.Now().Add((-8) * time.Minute)
	fmt.Println("time until +8:", time.Until(nowPlus8))
	fmt.Println("time until -8:", time.Until(nowMinus8))
	if time.Until(nowPlus8) < 9*time.Minute {
		fmt.Println(tnow, "is less than 9 min below", nowPlus8)
	}
	if time.Until(nowMinus8) < 0 {
		fmt.Println(nowMinus8, "is less than 0 min until", tnow)
	}

}
