package main

func BinarySearch(needle int, haystack []int) bool {

	var (
		upper     = len(haystack)
		lower int = 0
	)

	for lower < upper {

		middle := (lower + (upper - lower)) / 2
		value := haystack[middle]

		if value == needle {
			return true
		} else if value > needle {
			upper = middle
		} else {
			lower = middle + 1
		}
	}

	return false
}
