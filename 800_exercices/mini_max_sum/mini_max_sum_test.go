package main

import (
	"bytes"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {

	assertEqualsString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	cases := []struct {
		description string
		numbers     []int32
		want        string
	}{
		{"1,2,3,4,5", []int32{1, 2, 3, 4, 5}, "10 14"},
		{
			"256741038,623958417,467905213,714532089,938071625",
			[]int32{256741038, 623958417, 467905213, 714532089, 938071625},
			"2063136757 2744467344",
		},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			buffer := bytes.Buffer{}

			miniMaxSum(&buffer, test.numbers)

			assertEqualsString(t, buffer.String(), test.want)

		})
	}

}
