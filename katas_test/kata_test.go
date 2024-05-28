package katas_test

import (
	"dsa-katas/katas"
	"dsa-katas/utils"
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	if !katas.BinarySearch(3, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Fatal()
	}

	if katas.BinarySearch(3, []int{1, 2, 4, 5, 6, 7, 8, 9}) {
		t.Fatal()
	}
}

func TestSelectionSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	arr := []int{5, 3, 1, 2, 4}
	katas.SelectionSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Fatal("not sorted")
	}
}

func TestBubbleSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}

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
	sorted := []int{1, 2, 3, 4, 5}
	arr := []int{5, 3, 1, 2, 4}
	katas.InsertionSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Fatal("not sorted")
	}
}

func TestMergeSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	arr := []int{5, 3, 1, 2, 4}
	res := katas.MergeSort(arr)

	if !reflect.DeepEqual(res, sorted) {
		t.Fatal("not sorted")
	}
}

func TestQuickSort(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	arr := []int{5, 3, 1, 2, 4}
	katas.QuickSort(arr)

	if !reflect.DeepEqual(arr, sorted) {
		t.Fatal("not sorted")
	}
}

func TestBFS(t *testing.T) {
	if !reflect.DeepEqual(katas.BFS(3, utils.TestDAG), []int{3, 5, 6}) {
		t.Fatal()
	}

	if !reflect.DeepEqual(
		katas.BFS(3, utils.TestGraph),
		[]int{3, 1, 2, 5, 4, 6},
	) {
		t.Fatal()
	}
}

func TestDFS(t *testing.T) {
	if katas.DFS(3, 1, utils.TestDAG) {
		t.Fatal()
	}

	if !katas.DFS(3, 1, utils.TestGraph) {
		t.Fatal()
	}
}
