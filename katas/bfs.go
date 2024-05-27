package katas

import (
	"dsa-katas/utils"
)

func BFS(start int, needle int, graph utils.Graph) bool {
	idx := 0
	queue := []int{start}
	visitedNodes := make(map[int]bool)
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
		if node == needle {
			return true
		}
		visitedNodes[node] = true
		for neighbor := range graph[node] {
			queue = append(queue, neighbor)
		}
	}
	return false
}
