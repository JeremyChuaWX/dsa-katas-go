package katas

import "dsa-katas/utils"

type heap struct{}

func (h *heap) insert(tuple []int) {}

func (h *heap) pop() []int {
	return []int{}
}

func relax(weight int) int {}

func Dijkstra(start int, end int, graph utils.WeightedGraph) []int {
	ptr := 0
	queue := []int{start}
	visitedNodes := make(map[int]bool)
	heap := heap{}
	path := []int{}
	for {
		if ptr >= len(queue) {
			break
		}
		curr := queue[ptr]
		ptr++
		visited, ok := visitedNodes[curr]
		if ok && visited {
			continue
		}
		visitedNodes[curr] = true
		path = append(path, curr)
		for _, neighbor := range graph[curr] {
			node := neighbor[0]
			weight := neighbor[1]
			relaxed := relax(weight)
			heap.insert([]int{node, relaxed})
			minNode := heap.pop()
			queue = append(queue, minNode[0])
		}
	}
	return path
}
