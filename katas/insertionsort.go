package katas

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
			} else {
				break
			}
		}
	}
}
