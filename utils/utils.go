package utils

type Graph map[int][]int

var TestDAG Graph = map[int][]int{
	1: {2, 3},
	2: {3, 4},
	3: {5},
	4: {5, 6},
	5: {6},
	6: {},
}

var TestGraph Graph = map[int][]int{
	1: {2, 3},
	2: {1, 3, 4},
	3: {1, 2, 5},
	4: {2, 5, 6},
	5: {3, 4, 6},
	6: {4, 5},
}

type WeightedGraph map[int][][]int

var TestWeightedGraph = map[int][][]int{
	1: {{2, 4}, {3, 2}},
	2: {{1, 4}, {3, 1}, {4, 7}},
	3: {{1, 2}, {2, 1}, {5, 3}},
	4: {{2, 7}, {5, 2}, {6, 5}},
	5: {{3, 3}, {4, 2}, {6, 1}},
	6: {{4, 5}, {5, 1}},
}
