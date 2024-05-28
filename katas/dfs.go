package katas

import "dsa-katas/utils"

func dfs(
	curr int,
	graph utils.Graph,
	visitedNodes map[int]bool,
	path []int,
) []int {
	visited, ok := visitedNodes[curr]
	if ok && visited {
		return path
	}
	visitedNodes[curr] = true
	path = append(path, curr)
	for _, neighbor := range graph[curr] {
		path = dfs(neighbor, graph, visitedNodes, path)
	}
	return path
}

func DFS(start int, graph utils.Graph) []int {
	visitedNodes := make(map[int]bool)
	return dfs(start, graph, visitedNodes, []int{})
}
