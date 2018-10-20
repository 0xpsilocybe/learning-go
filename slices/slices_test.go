package slices

import (
	"fmt"
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

func ExampleSum() {
	numbers := []int{1, 1, 1, 1, 1}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 5
}

