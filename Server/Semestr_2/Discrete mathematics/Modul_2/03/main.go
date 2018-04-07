// Cmp project main.go
package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "strconv"
)

const MaxInt = 32767

type Vertex struct {
        name  string
	index int
	price int
	color int
	//0 - Белый
	//1 - Красный
	//2 - Синий
	timeIn int
	low    int
	comp   int
	dist   int
	parent []*Vertex
	flag   bool
}

type Tocken struct {
	teg int
	//0 - number
	//1 - name
	//2 - endl
	//3 - endf
	value interface{}
}

type lexems struct {
	array        []Tocken
	mapObjective map[string]int
	V            []*Vertex
	G            [][]*Vertex
	U            [][]int
	flagU        [][]int
	n            int
	sum          int
	max          int
	arrayMax     []*Vertex
	helpArr      []bool
}

func (lex *lexems) lexer(text string) {
	flagNum := false
	flagVal := false
	var c byte
	start := 0
	for i := 0; i < len(text); i++ {
		c = text[i]
		if flagNum {
			if !(c >= '0' && c <= '9') {
				lex.n++
				num, _ := strconv.Atoi(text[start:i])
				lex.sum += num
				lex.array = append(lex.array, Tocken{teg: 0, value: num})
				flagNum = false
			} else {
				continue
			}
		}
		if flagVal {
			if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
				lex.array = append(lex.array, Tocken{teg: 1, value: text[start:i]})
				flagVal = false
			} else {
				continue
			}
		}
		switch c {
		case ';':
			lex.array = append(lex.array, Tocken{teg: 2, value: nil})
		case '(', ')', '<', ' ', '\n', '\t':
			continue
		default:
			if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
				flagVal = true
				start = i
			} else if c >= '0' && c <= '9' {
				flagNum = true
				start = i
			} else {
				//ERROR
				fmt.Println("error")
			}
		}
	}
	if flagNum || flagVal {
		//ERROR
		fmt.Println("error")
	}
	lex.array = append(lex.array, Tocken{teg: 3, value: nil})
	lex.G = make([][]*Vertex, lex.n)
	lex.V = make([]*Vertex, lex.n)
	lex.helpArr = make([]bool, lex.n)
	lex.U = make([][]int, lex.n)
	for i := 0; i < lex.n; i++ {
		lex.U[i] = make([]int, lex.n)
	}
}

func (lex *lexems) graphenizaciya() {
	flagendl := false
	k := 0
	var v, w *Vertex
	for i := 0; i < len(lex.array)-1; i++ {
		if lex.array[i].teg == 1 && lex.array[i+1].teg == 0 {
			v = &Vertex{
				name:   lex.array[i].value.(string),
				index:  k,
				price:  lex.array[i+1].value.(int),
				color:  0,
				timeIn: 0,
				comp:   -1,
				low:    0,
				dist:   0,
				parent: nil,
				flag:   true,
			}
			k++
			lex.V[v.index] = v
			lex.mapObjective[v.name] = v.index
			if flagendl {
				if lex.U[w.index][v.index] == 0 {
					lex.G[w.index] = append(lex.G[w.index], v)
					lex.U[w.index][v.index] = 1
				}
				if v.index == w.index {
					v.color = 2
				}
				w = v
			} else {
				w = v
				flagendl = true
			}
			i++
			continue
		}
		if lex.array[i].teg == 1 && lex.array[i+1].teg != 0 {
			v = lex.V[lex.mapObjective[lex.array[i].value.(string)]]
			if flagendl {
				if lex.U[w.index][v.index] == 0 {
					lex.G[w.index] = append(lex.G[w.index], v)
					lex.U[w.index][v.index] = 1
				}
				if v.index == w.index {
					v.color = 2
				}
				w = v
			} else {
				w = v
				flagendl = true
			}
			continue
		}
		if lex.array[i].teg == 2 {
			flagendl = false
		}
	}
}

func (lex *lexems) PaintChildren(v *Vertex) {
	v.color = 2
	lex.sum -= v.price
	for _, u := range lex.G[v.index] {
		if u.color != 2 {
			lex.PaintChildren(u)
		}
	}
}

func (lex *lexems) Paint() {
	var flag bool
	for i := 0; i < lex.n-1; i++ {
		flag = false
		for j := i + 1; j < lex.n; j++ {
			if lex.V[i].comp == lex.V[j].comp {
				flag = true
				lex.PaintChildren(lex.V[j])
			}
		}
		if flag {
			lex.PaintChildren(lex.V[i])
		}
	}
}

func (lex *lexems) Search(u, v *Vertex) bool {
	for _, x := range v.parent {
		if u == x {
			return true
		}
	}
	return false
}

func (lex *lexems) Relax(u, v *Vertex, w int) bool {
	if u.dist == MaxInt {
		v.dist = w
		v.parent = make([]*Vertex, 0)
		v.parent = append(v.parent, u)
		return true
	}
	if u.dist+w == v.dist && !lex.Search(u, v) {
		v.parent = append(v.parent, u)
		return true
	}
	if u.dist+w > v.dist {
		v.dist = u.dist + w
		v.parent = make([]*Vertex, 0)
		v.parent = append(v.parent, u)
		return true
	}
	return false
}
func (lex *lexems) Descent(v *Vertex) {
	if v.color == 2 {
		return
	}
	v.flag = true
	v.parent = make([]*Vertex, 0)
	for _, u := range lex.G[v.index] {
		lex.Descent(u)
	}
}
func (lex *lexems) Clean() {
	for i, _ := range lex.helpArr {
		lex.helpArr[i] = false
	}
}

var countMaxEdge int

func (lex *lexems) Climb(v *Vertex) {
	if v == nil || lex.helpArr[v.index] {
		return
	}
	lex.arrayMax = append(lex.arrayMax, v)
	lex.helpArr[v.index] = true
	for _, u := range v.parent {
		lex.U[u.index][v.index] = countMaxEdge
		lex.Climb(u)
	}
}

func (lex *lexems) BellmanFord(s *Vertex) int {
	if s.color == 2 {
		return -1
	}
	for _, v := range lex.V {
		v.flag = false
		v.dist = 0
		v.parent = nil
	}
	lex.Descent(s)
	var v, u *Vertex
	var j int
	s.dist = MaxInt
	for i := 1; i < len(lex.V); i++ {
		for j, u = range lex.V {
			if u.flag {
				for _, v = range lex.G[j] {
					if v.flag {
						lex.Relax(u, v, v.price)
					}
				}
			}
		}
	}
	ArrMaxVer := make([]*Vertex, 0)
	max := 0
	for _, m := range lex.V {
		if m.dist == MaxInt && max == 0 {
			ArrMaxVer = make([]*Vertex, 0)
			ArrMaxVer = append(ArrMaxVer, m)
		} else if m.dist > max && m.dist != MaxInt {
			max = m.dist
			ArrMaxVer = make([]*Vertex, 0)
			ArrMaxVer = append(ArrMaxVer, m)
		} else if m.dist != 0 && m.dist == max && m.dist != MaxInt {
			ArrMaxVer = append(ArrMaxVer, m)
		}
	}
	max += s.price
	if max > lex.max {
		countMaxEdge++
		lex.Clean()
		lex.arrayMax = make([]*Vertex, 0)
		for _, v = range ArrMaxVer {
			lex.Climb(v)
		}
	} else if max == lex.max {
		lex.Clean()
		for _, v = range ArrMaxVer {
			lex.Climb(v)
		}
	}
	return max
}

func (lex *lexems) Helper() {
	var max int
	for _, v := range lex.V {
		max = lex.BellmanFord(v)
		if max > lex.max {
			lex.max = max
		}
	}
	return
}

func (lex *lexems) Out() {
	fmt.Printf("digraph {\n")
	for _, v := range lex.V {
		fmt.Printf("\t%s [label = \"%s(%d)\"", v.name, v.name, v.price)
		switch v.color {
		case 0:
			fmt.Printf("]\n")
		case 1:
			fmt.Printf(", color = red]\n")
		case 2:
			fmt.Printf(", color = blue]\n")
		}
	}
	for i, v := range lex.V {
		for _, u := range lex.G[i] {
			fmt.Printf("\t%s -> %s", v.name, u.name)
			if v.color == 2 && u.color == 2 {
				fmt.Printf(" [color = blue]\n")
				continue
			}
			if lex.U[v.index][u.index] == countMaxEdge {
				fmt.Printf(" [color = red]\n")
				continue
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("}")
}

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
	for _, u = range G[v.index] {
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
			if u.index == v.index {
				break
			}
		}
		count++
	}
}

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
        text := string(buf)
	lex := lexems{
		array:        make([]Tocken, 0),
		mapObjective: make(map[string]int),
		V:            nil,
		G:            nil,
		n:            0,
		sum:          0,
		max:          0,
		arrayMax:     nil,
	}
        countMaxEdge = 1
	lex.lexer(text + ";")
	lex.graphenizaciya()
	Tarjan(lex.G, lex.V)
	lex.Paint()
	lex.Helper()
	for _, v := range lex.arrayMax {
		v.color = 1
	}
	lex.Out()
}
