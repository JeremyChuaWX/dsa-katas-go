package katas

import (
	"dsa-katas/utils"
)

func BFS(start int, graph utils.Graph) []int {
	idx := 0
	queue := []int{start}
	visitedNodes := make(map[int]bool)
	path := []int{}
	for {
		if idx >= len(queue) {
			break
		}
		node := queue[idx]
		idx++
		visited, ok := visitedNodes[node]
		if ok && visited {
			continue
		}
		visitedNodes[node] = true
		path = append(path, node)
		for _, neighbor := range graph[node] {
			queue = append(queue, neighbor)
		}
	}
	return path
}
