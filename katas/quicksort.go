package katas

func pivot(arr []int, start int, end int) int {
	left := start
	right := end
	pivot := (start + end) / 2

	for left <= right {
		if arr[left] <= arr[pivot] {
			left++
		} else if arr[right] >= arr[pivot] {
			right--
		} else {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
			left++
			right--
		}
	}

	if pivot < right {
		tmp := arr[right]
		arr[right] = arr[pivot]
		arr[pivot] = tmp
		pivot = right
	} else if pivot > left {
		tmp := arr[left]
		arr[left] = arr[pivot]
		arr[pivot] = tmp
		pivot = left
	}

	return pivot
}

func quickSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	pivotIdx := pivot(arr, start, end)
	quickSort(arr, start, pivotIdx)
	quickSort(arr, pivotIdx+1, end)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
