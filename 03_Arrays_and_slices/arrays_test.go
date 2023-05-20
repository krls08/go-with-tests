package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)

		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum multiple slices", func(t *testing.T) {
		ns := [][]int{}
		ns = append(ns, []int{1, 2}, []int{0, 9}, []int{3, 4, 5})
		//fmt.Println(">>>>>>>> ns", ns)
		got := SumAll([]int{1, 2}, []int{0, 9}, []int{3, 4, 5}) //=>[]{3,9,12}
		want := []int{3, 9, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, ns)
		}

	})

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !(reflect.DeepEqual(got, want)) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

func TestRemoveIndex(t *testing.T) {

	t.Run("delete first element", func(t *testing.T) {
		slice0 := []string{"a", "b", "c"}
		fmt.Println("--------- slice0", slice0)
		got := SliceRemoveIndex(slice0, 0)
		want := []string{"c", "b"}
		assertSlicesStr(t, got, want)
	})

	t.Run("delete last element", func(t *testing.T) {
		slice0 := []string{"a", "b", "c"}
		fmt.Println("--------- slice0", slice0)
		got := SliceRemoveIndex(slice0, 2)
		want := []string{"a", "b"}
		assertSlicesStr(t, got, want)
	})

	t.Run("delete on one size length", func(t *testing.T) {
		slice0 := []string{"a"}
		fmt.Println("--------- slice0", slice0)
		got := SliceRemoveIndex(slice0, 0)
		fmt.Println("--------- got", got)
		want := []string{}
		assertSlicesStr(t, got, want)
	})

	t.Run("panic -> slice of length 0", func(t *testing.T) {
		slice0 := []string{}
		fmt.Println("--------- slice0", slice0)
		wrapperPanic := func() { SliceRemoveIndex(slice0, 0) }
		shouldPanic(t, wrapperPanic)
		//wrapperPanic()
	})
}

func TestGenSliceRemoveIndex(t *testing.T) {

	t.Run("Ok string slice", func(t *testing.T) {
		slice0 := []string{"a", "b"}
		got := GenSliceRemoveIndex(slice0, 1)
		want := []string{"a"}
		assertSlicesStr(t, got, want)
	})
	t.Run("Ok int slice", func(t *testing.T) {
		slice0 := []int{10, 20}
		got := GenSliceRemoveIndex(slice0, 1)
		want := []int{10}
		assertSlicesGen(t, got, want)
	})
}

func shouldPanic(t *testing.T, callback func()) {
	t.Helper()
	defer func() { _ = recover() }()
	callback()
	t.Errorf("should have panicked")
}

func assertSlicesGen[K any](t testing.TB, got, want []K) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertSlicesStr(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCreateStringSlice(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		slice0 := CreateStringSlice()
		fmt.Print("new slice", slice0)
	})
}
