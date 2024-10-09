package main

import (
	"fmt"
)

func main() {
	var values = []int{1, 2, 4, 6, 8, 20}
	fmt.Println(findNumber(values, 5))
	fmt.Println(proAdder(2, 5, 8, 7))

	// Example usage of Kosaraju's algorithm
	n, m := 5, 5
	fmt.Println(m)
	adj := make([][]int, n)
	edges := [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
	}
	fmt.Println("Number of strongly connected components:", kosaraju(n, adj))
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func findNumber(value []int, target int) bool {
	left := 0
	right := len(value) - 1

	for left <= right {
		mid := left + (right-left)/2

		if value[mid] == target {
			return true
		} else if value[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func dfs(node int, vis []bool, adj [][]int, st *[]int) {
	vis[node] = true
	for _, child := range adj[node] {
		if !vis[child] {
			dfs(child, vis, adj, st)
		}
	}
	*st = append(*st, node)
}

func dfs2(node int, vis []bool, adj [][]int) {
	vis[node] = true
	for _, child := range adj[node] {
		if !vis[child] {
			dfs2(child, vis, adj)
		}
	}
}

func kosaraju(V int, adj [][]int) int {
	vis := make([]bool, V)
	st := make([]int, 0)

	for i := 0; i < V; i++ {
		if !vis[i] {
			dfs(i, vis, adj, &st)
		}
	}

	adjT := make([][]int, V)
	for u := 0; u < V; u++ {
		vis[u] = false
		for _, v := range adj[u] {
			adjT[v] = append(adjT[v], u)
		}
	}

	scc := 0
	for len(st) > 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]

		if !vis[node] {
			dfs2(node, vis, adjT)
			scc++
		}
	}

	return scc
}
