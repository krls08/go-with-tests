// Nice test-yourself quiz for green belts:
// why isn't this a type error? :)
// https://twitter.com/josh_cheek/status/1629040134988390403

package main

import (
	"fmt"
)

type Objs []Obj
type Obj interface{ GetType() string }
type IntObj struct{ Int int }
type StrObj struct{ Str string }

func (i IntObj) GetType() string { return "int" }
func (s StrObj) GetType() string { return "str" }

func main() {
	i := IntObj{Int: 12}
	s := StrObj{Str: "omg"}
	omg(Objs{i, &i, s, &s})
}

func omg(objs Objs) {
	for _, obj := range objs {
		if intObj, ok := obj.(IntObj); ok {
			fmt.Printf("IntObj: %s, %d\n", intObj.GetType(), intObj.Int)
		} else if strObj, ok := obj.(StrObj); ok {
			fmt.Printf("StrObj: %s, %s\n", strObj.GetType(), strObj.Str)
		} else {
			fmt.Printf("Unknown: %s\n", obj.GetType())
		}
	}
}
