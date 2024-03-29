package main

import (
	"fmt"
)

type Flight struct {
	xName    *string
	xAvModel *string
}

func main() {

	name1 := "carlos"
	name2 := ""
	var name3 string
	var name4 *string
	exPtr(&name1)
	exPtr(&name2)
	exPtr(&name3)
	exPtr(name4)

	fmt.Printf("\n--------------------\n\n")
	var name *string
	aFlight := Flight{}
	aFlight.xName = name
	fmt.Println("flight avModel", aFlight.xAvModel)
	fmt.Println("flight name", aFlight.xName)

	fmt.Printf("\n----------- Lists -----------\n\n")
	l := []string{"UNKNOWN", "COMMAND_0"}
	listLen(l)
	listLen(nil)
}

func exPtr(name *string) {
	if name != nil {
		fmt.Println("name is not nil:", *name)
	} else {
		fmt.Println("name is nil!. Default name is pepito")
	}
}
func listLen(list []string) {
	fmt.Println("list ->", list)
	fmt.Println("length:", len(list))
}
