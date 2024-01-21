package bubblesort

func BubbleSort(haystack []int) []int {
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(haystack)-1-i; j++ {
			if haystack[j] > haystack[j+1] {
				haystack[j], haystack[j+1] = haystack[j+1], haystack[j]
			}
		}
	}
	return haystack
}
