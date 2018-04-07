// EqDist project main.go
package main

import (
        "fmt"
	"sort"
        "github.com/skorobogatov/input"
)

func BFS(G, D [][]int, root int) {
	var v int
	color := make([]bool, len(G))
	q := make([]int, 0, len(G))

	color[root] = true
	q = append(q, root)
	for len(q) > 0 {
		v = q[0]
		q = q[1:]
		for _, u := range G[v] {
			if !color[u] {
				color[u] = true
				q = append(q, u)
				D[root][u] = D[root][v] + 1
			}
		}
	}
}
func main() {
	var i, a, b, n, m, k, j int
	input.Scanf("%d%d", &n, &m)
	G := make([][]int, n)
	D := make([][]int, n)
	for i = 0; i < n; i++ {
		D[i] = make([]int, n)
	}
	for i = 0; i < m; i++ {
		input.Scanf("%d%d", &a, &b)
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	input.Scanf("%d", &k)
	V := make([]int, k)
	for i = 0; i < k; i++ {
		input.Scanf("%d", &V[i])
		BFS(G, D, V[i])
	}
	sort.Ints(V)
	f := true
	k = 0
	for i = 0; i < n; i++ {
		if k < len(V) && i == V[k] {
			k++
			continue
		} else {
			if D[V[0]][i] == 0 {
				continue
			}
			for j = 1; j < len(V); j++ {
				if D[V[j]][i] != D[V[0]][i] {
					break
				}
				if j >= len(V)-1 {
					f = false
					fmt.Printf("%d ", i)
				}
			}
		}
	}
	if f {
		fmt.Printf("-")
	}
}
