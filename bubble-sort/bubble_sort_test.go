package bubblesort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	errorMsg := "Not sorted, got %v"
	haystack := BubbleSort([]int{1, 69, 71, 4, 99, 90, 22, 5235, 1337, 5})

	sorted := true
	for i := 1; i < len(haystack); i++ {
		if haystack[i-1] > haystack[i] {
			sorted = false
		}
	}
	if !sorted {
		t.Errorf(errorMsg, haystack)
	}
}
