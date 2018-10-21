package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertCollectionSum := func(got int, want int, slice []int, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d', want '%d', given %v", got, want, slice)
		}
	}

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCollectionSum(got, want, numbers, t)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{4, 5}, []int{1, 9})
	want := []int{9, 10}
	if !reflect.DeepEqual(got, want) { 
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("tails are being summed correctly", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 5, 5})
		want := []int{2, 10}
		checkSum(t, got, want)
	})

	t.Run("safely sum empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSum(t, got, want)
	})

}

func ExampleSum() {
	numbers := []int{1, 1, 1, 1, 1}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 5
}

func ExampleSumAll() {
	sliceOne := []int{1, 2}
	sliceTwo := []int{0, 9}
	sum := SumAll(sliceOne, sliceTwo)
	fmt.Println(sum)
	// Output: [3 9]
}

func ExampleSumAllTails() {
	emptySlice := []int{}
	slice := []int{1, 4, 11}
	sum := SumAllTails(slice, emptySlice)
	fmt.Println(sum)
	// Output: [15 0]
}

