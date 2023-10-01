package main

func binarySearch(needle int, arr []int, start int, end int) bool {
	if start >= end {
		return false
	}

	mid := start + (end-start)/2

	if needle == arr[mid] {
		return true
	} else if needle > arr[mid] {
		return binarySearch(needle, arr, mid+1, end)
	} else {
		return binarySearch(needle, arr, start, mid)
	}
}

func BinarySearch(needle int, arr []int) bool {
	return binarySearch(needle, arr, 0, len(arr))
}
