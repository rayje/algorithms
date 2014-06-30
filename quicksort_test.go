package algorithms

import (
	"testing"
	"fmt"
)

func Test_NewRandomNodeID(t *testing.T) {
	a := []string{"S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"}

	show(a)
	Quicksort(a)
	show(a)
}

func show(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i] + " ")
	}
	fmt.Println()
}