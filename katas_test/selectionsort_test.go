package katas_test

import (
	"dsa-katas/katas"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	arr := []int{5, 3, 1, 2, 4}

	katas.SelectionSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Fatal()
	}
}
