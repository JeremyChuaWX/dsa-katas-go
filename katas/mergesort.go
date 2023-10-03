package katas

func partition(arr []int) int {
	return len(arr) / 2
}

func merge(left []int, right []int) []int {
	var res []int
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	for i < len(left) {
		res = append(res, left[i])
		i++
	}
	for j < len(right) {
		res = append(res, right[j])
		j++
	}
	return res
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := partition(arr)
	arr1 := MergeSort(arr[:mid])
	arr2 := MergeSort(arr[mid:])
	return merge(arr1, arr2)
}
