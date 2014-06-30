package sort

import (
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	a := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}

	SelectionSort(a)

	if !isSorted(a) {
		t.Log("a is not sorted: ", a)
		t.Fail()
	}
}