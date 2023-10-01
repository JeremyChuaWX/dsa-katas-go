package main

import "testing"

func TestBinarySearch(t *testing.T) {
	if !BinarySearch(3, []int{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Fatal()
	}

	if BinarySearch(3, []int{1, 2, 4, 5, 6, 7, 8, 9}) {
		t.Fatal()
	}
}
