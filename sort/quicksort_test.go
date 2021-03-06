package sort

import (
	"testing"
)

func Test_Quicksort(t *testing.T) {
	a := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}

	Quicksort(a)

	if !isSorted(a) {
		t.Log("a is not sorted: ", a)
		t.Fail()
	}
}