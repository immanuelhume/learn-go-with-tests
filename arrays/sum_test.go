package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	AssertEqual := func(t testing.TB, got, want int, given []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, given)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		AssertEqual(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		AssertEqual(t, got, want, numbers)

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("get the tails", func(t *testing.T) {

		got := SumAllTails([]int{2, 3, 4}, []int{0, 3, 1})
		want := []int{7, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely handles empty array", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
