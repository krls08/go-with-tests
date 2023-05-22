package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with none string field",
			struct {
				Name string
				Age  int
			}{
				"Chris", 33,
			},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to objects",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			//walk(test.Input, func(input string) {
			//	got = append(got, input)
			//})
			// -- or --
			fn := func(input string) {
				got = append(got, input)
			}
			walk(test.Input, fn)

			fmt.Println("----- got =>", got)

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}

/// test 0 example
//func TestWalk(t *testing.T) {
//	expected := "Chris"
//	var got []string
//
//	x := struct {
//		Name string
//	}{expected}
//
//	walk(x, func(input string) {
//		got = append(got, input)
//	})
//
//	fmt.Println("-------- got", got)
//
//	if len(got) != 1 {
//		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
//	}
//
//}
