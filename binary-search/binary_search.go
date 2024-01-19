package binarysearch

func binarySearch(haystack []int, needle int) bool {
	lowIndex := 0
	highIndex := len(haystack)

	for lowIndex < highIndex {
		midIndex := lowIndex + ((highIndex - lowIndex) / 2)
		testValue := haystack[midIndex]

		if testValue == needle {
			return true
		}
		// Don't check the middle again
		// Low will be inclusive, high exclusive
		if needle > testValue {
			lowIndex = midIndex + 1
		} else {
			highIndex = midIndex
		}
	}

	return false
}
