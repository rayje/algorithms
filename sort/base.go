package sort

import "fmt"

func swap(a []string, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func show(a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i] + " ")
	}
	fmt.Println()
}

func isSorted(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}