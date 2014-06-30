package sort


func SelectionSort(a []string) {
	N := len(a)

	for i := 0; i < N; i++ {
		min := i

		for j := i + 1; j < N; j++ {
			// Find the min value
			if a[j] < a[min] {
				min = j
			}
		}

		// Swap the i-th value with min
		swap(a, i, min)
	}
}