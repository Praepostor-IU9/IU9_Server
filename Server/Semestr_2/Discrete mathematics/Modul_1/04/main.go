// BridgeNum project main.go
package main

import (
        "fmt"
        "github.com/skorobogatov/input"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var bridge, timer int

func find_bridg(G [][]int) {
	var i int
	timer = 0
	bridge = 0
	tin := make([]int, len(G))
	fup := make([]int, len(G))
	used := make([]bool, len(G))
	for i = 0; i < len(G); i++ {
		dfs(G, fup, tin, i, -1, used)
	}
}
func dfs(G [][]int, fup, tin []int, v, p int, used []bool) {
	var i int
	used[v] = true
	timer++
	tin[v] = timer
	fup[v] = timer
	for i = 0; i < len(G[v]); i++ {
		if G[v][i] == p {
			continue
		}
		if used[G[v][i]] {
			fup[v] = min(fup[v], tin[G[v][i]])
		} else {
			dfs(G, fup, tin, G[v][i], v, used)
			fup[v] = min(fup[v], fup[G[v][i]])
			if fup[G[v][i]] > tin[v] {
				bridge++
			}
		}
	}
}
func main() {
	var i, a, b, n, m int
	input.Scanf("%d%d", &n, &m)
	G := make([][]int, n)
	for i = 0; i < m; i++ {
		input.Scanf("%d%d", &a, &b)
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	find_bridg(G)
	fmt.Printf("%d", bridge)
}