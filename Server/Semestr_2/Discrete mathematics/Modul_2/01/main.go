// GraphBase project main.go
package main

import (
        "fmt"
)

//Stack
type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value *Vertex
		prev  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Pop() *Vertex {
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value *Vertex) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

//Stack end

type Vertex struct {
	name   int
	timeIn int
	low    int
	comp   int
}

var time int
var count int

func Tarjan(G [][]*Vertex, V []*Vertex) {
	var v *Vertex
	time = 1
	count = 0
	S := New()
	for _, v = range V {
		if v.timeIn == 0 {
			VisitVertex_Tarjan(G, v, S)
		}
	}
}
func VisitVertex_Tarjan(G [][]*Vertex, v *Vertex, S *Stack) {
	var u *Vertex
	v.timeIn, v.low = time, time
	time += 1
	S.Push(v)
	for _, u = range G[v.name] {
		if u.timeIn == 0 {
			VisitVertex_Tarjan(G, u, S)
		}
		if u.comp == -1 && v.low > u.low {
			v.low = u.low
		}
	}
	if v.timeIn == v.low {
		for {
			u = S.Pop()
			u.comp = count
			if u.name == v.name {
				break
			}
		}
		count++
	}
}
func main() {
	var m, n, a, b, i, j int
	var v, u *Vertex
	fmt.Scan(&n, &m)
	V := make([]*Vertex, n)
	for i = 0; i < n; i++ {
		V[i] = &Vertex{
			name:   i,
			timeIn: 0,
			comp:   -1,
			low:    0,
		}
	}
	G := make([][]*Vertex, n)
	for i = 0; i < m; i++ {
		fmt.Scan(&a, &b)
		G[a] = append(G[a], V[b])
	}
	Tarjan(G, V)
	Cond := make([][]bool, count)
	for i = 0; i < count; i++ {
		Cond[i] = make([]bool, count)
	}

	for i, v = range V {
		for j, u = range G[i] {
			if v.comp != u.comp {
				Cond[v.comp][u.comp] = true
			}
		}
	}
	for i = 0; i < count; i++ {
		for j = 0; j < count; j++ {
			if Cond[j][i] {
				break
			}
		}
		if j == count {
			for _, v = range V {
				if v.comp == i {
					fmt.Printf("%d ", v.name)
					break
				}
			}
		}
	}
}