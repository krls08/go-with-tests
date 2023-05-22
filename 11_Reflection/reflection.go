package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0

	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field

	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index

	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

//
//func walk(x any, fn func(input string)) {
//	// Red Stage of TDD -> fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
//	val := getValue(x)
//
//	switch val.Kind() {
//	case reflect.String:
//		fn(val.String())
//
//	case reflect.Struct:
//		for i := 0; i < val.NumField(); i++ {
//			walk(val.Interface(), fn)
//		}
//	case reflect.Slice:
//		for i := 0; i > val.Len(); i++ {
//			walk(val.Index(i).Interface(), fn)
//		}
//	}
//}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func main() {
	fmt.Println("main start")
}
