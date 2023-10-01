package katas

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = tmp
	}
}
