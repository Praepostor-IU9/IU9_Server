//MaxComponent project main.go
package main

import (
        "fmt"
        "github.com/skorobogatov/input"
	"sort"
)

//Stack
type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value int
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Pop() int {
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value int) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

//Stack end

func DFS(G [][]int, n int) []int {
	color := make([]int8, n)
	var maxcomponent []int
	var component []int
	var E, maxE, w, v, u int
	maxE = 0
	s := New()
	for w = 0; w < n; w++ {
		if color[w] == 0 {
			s.Push(w)
			E = 1
			for s.length > 0 {
				v = s.Pop()
				if color[v] == 0 {
					component = append(component, v)
					color[v] = 1
					s.Push(v)
					for u = 0; u < len(G[v]); u++ {
						if color[G[v][u]] == 0 {
							E++
							s.Push(G[v][u])
						}
						if color[G[v][u]] == 1 {
							E++
						}
					}
				} else {
					color[v] = 2
				}
			}
			E -= len(component)
			if len(component) > len(maxcomponent) || len(component) == len(maxcomponent) && E > maxE {
				maxcomponent, component = component, maxcomponent
				maxE = E
			}
		}
		component = make([]int, 0)
	}
	return maxcomponent
}

func output(G [][]int, maxcomponent []int, n int) {
	var i, j, k int
	fmt.Printf("graph {\n")
	k = 0
	for i = 0; i < n; i++ {
		if i == maxcomponent[k] {
			fmt.Printf("\t%d [color = red]\n", i)
			k++
			if k == len(maxcomponent) {
				i++
				break
			}
		} else {
			fmt.Printf("\t%d\n", i)
		}
	}
	for ; i < n; i++ {
		fmt.Printf("\t%d\n", i)
	}
	k = 0
	for i = 0; i < n; i++ {
		if i == maxcomponent[k] {
			for j = 0; j < len(G[i]); j++ {
				fmt.Printf("\t%d -- %d [color = red]\n", i, G[i][j])
			}
			k++
			if k == len(maxcomponent) {
				i++
				break
			}
		} else {
			for j = 0; j < len(G[i]); j++ {
				fmt.Printf("\t%d -- %d\n", i, G[i][j])
			}
		}
	}
	for ; i < n; i++ {
		for j = 0; j < len(G[i]); j++ {
			fmt.Printf("\t%d -- %d\n", i, G[i][j])
		}
	}
	fmt.Printf("}")
}

func main() {
	var i, a, b, n, m int
	input.Scanf("%d%d", &n, &m)
	G := make([][]int, n)
	G2 := make([][]int, n)
	for i = 0; i < m; i++ {
		input.Scanf("%d%d", &a, &b)
		G[a] = append(G[a], b)
		G2[a] = append(G2[a], b)
		G2[b] = append(G2[b], a)
	}
	maxcomponent := DFS(G2, n)
	sort.Ints(maxcomponent)
	output(G, maxcomponent, n)
}
