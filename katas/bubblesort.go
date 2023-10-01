package katas

// return swapped status
func BubbleSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		swap := false

		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap = true
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}

		if !swap {
			return true
		}
	}

	return false
}
