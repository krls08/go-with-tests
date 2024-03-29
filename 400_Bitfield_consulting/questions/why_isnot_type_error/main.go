// Nice test-yourself quiz for green belts:
// why isn't this a type error? :)
// https://twitter.com/josh_cheek/status/1629040134988390403

// Answer
// The spec says (in the Expressions / Calls) section:
// A method call x.m() is valid if the method set of (the type of) x contains m and the argument list can
// be assigned to the parameter list of m. If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m().
// In less grandiloquent words, if Foo has a method Bar, then *Foo has that method.
//


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
		//intObj0 := obj.(*IntObj)
		//fmt.Println("intObj0", intObj0)
		if intObj, ok := obj.(IntObj); ok {
			fmt.Printf("IntObj: %s, %d\n", intObj.GetType(), intObj.Int)
		} else if strObj, ok := obj.(StrObj); ok {
			fmt.Printf("StrObj: %s, %s\n", strObj.GetType(), strObj.Str)
		} else if intObjPnt, ok := obj.(*IntObj); ok {
			fmt.Printf("IntObjPnt: %s, %d\n", intObjPnt.GetType(), intObjPnt.Int)
		} else {
			fmt.Printf("Unknown: %s\n", obj.GetType())
			//i := (*IntObj)(obj)
			//if i := obj.(IntObj); ok {
			//	fmt.Println("---> ok, i", ok, i)
			//	fmt.Printf("type of: %T", i)
			//}
		}
	}
}
