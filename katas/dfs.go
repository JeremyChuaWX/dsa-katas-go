package katas

import "dsa-katas/utils"

func dfs(
	curr int,
	needle int,
	graph utils.Graph,
	visitedNodes map[int]bool,
) bool {
	if curr == needle {
		return true
	}
	visited, ok := visitedNodes[curr]
	if ok && visited {
		return false
	}
	visitedNodes[curr] = true
	for _, neighbor := range graph[curr] {
		if dfs(neighbor, needle, graph, visitedNodes) {
			return true
		}
	}
	return false
}

func DFS(start int, needle int, graph utils.Graph) bool {
	visitedNodes := make(map[int]bool)
	return dfs(start, needle, graph, visitedNodes)
}
