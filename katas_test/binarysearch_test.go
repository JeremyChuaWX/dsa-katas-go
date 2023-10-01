package katas_test

import (
	"dsa-katas/katas"
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
