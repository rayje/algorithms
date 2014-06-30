package sort

import (
	"math/rand"
)

// An implementation of a basic quicksort algorithm
func Quicksort(a []string) {
	shuffle(a)

	sort(a, 0, len(a) - 1)
}

func sort(a []string, lo int, hi int) {
	if hi <= lo {
		return
	}

	j := partition(a, lo, hi)
	sort(a, lo, j - 1)
	sort(a, j + 1, hi)
}

func partition(a []string, lo int, hi int) int {
	i := lo + 1
	j := hi

	v := a[lo]

	for {
		// Increment i until a[i] > v
		for ; a[i] < v; i++ {
			if i == hi {
				break
			}
 		}

 		for ; v < a[j]; j-- {
 			if j == lo {
 				break
 			}
 		}

 		if i >= j {
 			break
 		}
 
 		swap(a, i, j)
	} 

	swap(a, lo, j)

	return j
}

func swap(a []string, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func shuffle(a []string) {
	var r int;
	var tmp string;
	l := len(a)

	for i := 0; i < l; i++ {
		r = rand.Intn(l - i)
		tmp = a[i]
		a[i] = a[r]
		a[r] = tmp
	}
 }