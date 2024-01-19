package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	errorMsg := "Expected %t from needle '%d', got %t"
	haystack := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	needle := 3
	expected := true
	got := binarySearch(haystack, needle)
	if expected != got {
		t.Errorf(errorMsg, expected, needle, got)
	}

	needle = 55
	expected = false
	got = binarySearch(haystack, needle)
	if expected != got {
		t.Errorf(errorMsg, expected, needle, got)
	}

	needle = 69420
	expected = true
	got = binarySearch(haystack, needle)
	if expected != got {
		t.Errorf(errorMsg, expected, needle, got)
	}

	needle = 69421
	expected = false
	got = binarySearch(haystack, needle)
	if expected != got {
		t.Errorf(errorMsg, expected, needle, got)
	}

	needle = 1
	expected = true
	got = binarySearch(haystack, needle)
	if expected != got {
		t.Errorf(errorMsg, expected, needle, got)
	}

	needle = 0
	expected = false
	got = binarySearch(haystack, needle)
	if expected != got {
		t.Errorf(errorMsg, expected, needle, got)
	}
}
