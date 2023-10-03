package katas_test

import (
	"dsa-katas/katas"
	"reflect"
	"testing"
)

var sorted = []int{1, 2, 3, 4, 5}

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 1, 2, 4}
	katas.SelectionSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Fatal("not sorted")
	}
}

func TestBubbleSort(t *testing.T) {
	arr1 := []int{5, 3, 1, 2, 4}
	katas.BubbleSort(arr1)

	if !reflect.DeepEqual(arr1, sorted) {
		t.Fatal("not sorted")
	}

	arr2 := []int{1, 2, 3, 4, 5}
	swap := katas.BubbleSort(arr2)

	if swap {
		t.Fatal("fail to short circuit")
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 3, 1, 2, 4}
	katas.InsertionSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Fatal("not sorted")
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{5, 3, 1, 2, 4}
	res := katas.MergeSort(arr)

	if !reflect.DeepEqual(res, sorted) {
		t.Fatal("not sorted")
	}
}
