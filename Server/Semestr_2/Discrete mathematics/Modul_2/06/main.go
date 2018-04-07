// MapRoute project main.go
package main

import (
        "container/heap"
	"fmt"
        "github.com/skorobogatov/input"
)

const MaxInt = 32767

type Vertex struct {
	indexI int
	indexJ int
	index  int
	value  int
	dist   int
}

type PriorityQueue []*Vertex

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vertex)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(elem *Vertex) {
	heap.Fix(pq, elem.index)
}

func Relax(u, v *Vertex, w int) bool {
	if u.dist+w < v.dist {
		v.dist = u.dist + w
		return true
	}
	return false
}

func Dijcstra(G []*Vertex, root *Vertex, n int) int {
        var v, u *Vertex
	var i, j int
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, root)
	for pq.Len() > 0 {
		v = heap.Pop(&pq).(*Vertex)
		i = v.indexI
		j = v.indexJ
		if j > 0 {
			u = G[i*n+j-1]
			if Relax(v, u, u.value) {
				heap.Push(&pq, u)
			}
		}
		if j < n-1 {
			u = G[i*n+j+1]
			if Relax(v, u, u.value) {
				heap.Push(&pq, u)
			}
		}
		if i > 0 {
			u = G[(i-1)*n+j]
			if Relax(v, u, u.value) {
				heap.Push(&pq, u)
			}
		}
		if i < n-1 {
			u = G[(i+1)*n+j]
			if Relax(v, u, u.value) {
				heap.Push(&pq, u)
			}
		}
	}
	return G[(n-1)*(n+1)].dist
}

func main() {
	var i, j, n, val int
	input.Scanf("%d", &n)
	G := make([]*Vertex, n*n)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			input.Scanf("%d",&val)
			G[i*n+j] = &Vertex{
				indexI: i,
				indexJ: j,
				value:  val,
				dist:   MaxInt,
				index:  i*n + j,
			}
		}
	}
	G[0].dist = G[0].value
	fmt.Println(Dijcstra(G, G[0], n))
}
