// Kruskal project main.go
package main

import (
        "container/heap"
	"fmt"
	"math"
)

type ElemHeap struct {
	a, b  int
	dist  float64
	index int
}
type Heap []*ElemHeap

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}
func (h *Heap) Push(x interface{}) {
	n := len(*h)
	elem := x.(*ElemHeap)
	elem.index = n
	*h = append(*h, elem)
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	elem := old[n-1]
	elem.index = -1
	*h = old[0 : n-1]
	return elem
}

type ElemTree struct {
	parent *ElemTree
}

func MakeSet() *ElemTree {
	elem := new(ElemTree)
	elem.parent = elem
	return elem
}

func Find(elem *ElemTree) *ElemTree {
	if elem.parent == elem {
		return elem
	}
	elem.parent = Find(elem.parent)
	return elem.parent
}

func Union(elem1, elem2 *ElemTree) {
	root1 := Find(elem1)
	root2 := Find(elem2)
	root1.parent = root2
}

func main() {
	var i, j, ind, n int
        var dis float64
	fmt.Scan(&n)
	Tree := make([]*ElemTree, n)
	HP := make(Heap, n*(n-1)/2)
	X := make([]int, n)
	Y := make([]int, n)
	ind = 0
	for i = 0; i < n; i++ {
		Tree[i] = MakeSet()
		fmt.Scan(&X[i], &Y[i])
		for j = 0; j < i; j++ {
                        dis = float64((X[i]-X[j])*(X[i]-X[j]) + (Y[i]-Y[j])*(Y[i]-Y[j]))
			HP[ind] = &ElemHeap{
				a:     i,
				b:     j,
				dist:  math.Sqrt(dis),
				index: ind,
			}
			ind++
		}
	}
	heap.Init(&HP)
        var nE int
        var DIST float64
	var edge *ElemHeap
	nE = 0
	for nE < n-1 {
		edge = heap.Pop(&HP).(*ElemHeap)
		if Find(Tree[edge.a]) != Find(Tree[edge.b]) {
			Union(Tree[edge.a], Tree[edge.b])
			DIST += edge.dist
			nE++
		}
	}
	fmt.Printf("%.2f", DIST)
}