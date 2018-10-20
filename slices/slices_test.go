package slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if want != got {
		t.Errorf("got '%d', want '%d', given %v", got, want, numbers)
	}
}

func ExampleSum() {
	numbers := [5]int{1, 1, 1, 1, 1}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 5
}

