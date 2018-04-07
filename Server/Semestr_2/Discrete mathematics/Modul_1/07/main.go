// Prim project main.go
package main

import (
        "container/heap"
	"fmt"
)

type ElemHeap struct {
	val   int
	dist  int
	index int
}
type Heap []*Vertex

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].key < h[j].key
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}
func (h *Heap) Push(x interface{}) {
	n := len(*h)
	//elem := x.(*Vertex)
	x.(*Vertex).index = n
	*h = append(*h, x.(*Vertex))
}
func (pq *Heap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -2
	*pq = old[0 : n-1]
	return item
}
func (h *Heap) update(elem *Vertex, value, key int) {
	elem.value = value
	elem.key = key
	heap.Fix(h, elem.index)
}

type Edge struct {
	tail *Vertex
	dist int
}
type Vertex struct {
	name, index, key, value int
}

func MSTPrim(G [][]Edge, V []Vertex) int {
	var ed Edge
	var u, v *Vertex
	var a, dist int
	pq := make(Heap, 0)
	v = &V[0]
	for {
		v.index = -2
		for _, ed = range G[v.name] {
			u = ed.tail
			a = ed.dist
			if u.index == -1 {
				u.key = a
				u.value = v.name
				heap.Push(&pq, u)
			} else {
				if u.index != -2 && a < u.key {
					pq.update(u, v.name, a)
				}
			}
		}
		if pq.Len() == 0 {
			break
		}
		v = heap.Pop(&pq).(*Vertex)
		dist += v.key
	}
	return dist
}
func main() {
	var i, n, m, a, b, d int
	fmt.Scan(&n, &m)
	G := make([][]Edge, n)
	V := make([]Vertex, n)
	for i = 0; i < n; i++ {
		V[i].index = -1
		V[i].name = i
	}
	for i = 0; i < m; i++ {
		fmt.Scan(&a, &b, &d)
		G[a] = append(G[a], Edge{
			tail: &V[b],
			dist: d,
		})
		G[b] = append(G[b], Edge{
			tail: &V[a],
			dist: d,
		})
	}
	fmt.Printf("%d\n", MSTPrim(G, V))
}