package main

import (
        "fmt"
)

func SearchDel(n uint64) []uint64 {
	H := [1000]uint64{0}
	P := [100]uint64{0}
	var i, j, a uint64
	for i = 2; i*i <= n; i++ {
		if n%i == 0 {
			a = n / i
			H[0]++
			H[H[0]] = i
			for j = 1; j <= P[0]; j++ {
				if P[j]%a == 0 {
					break
				}
			}
			if j > P[0] {
				P[0]++
				P[P[0]] = a
			}
		}
	}
	for i = H[0]; i > 0; i-- {
		for j = 1; j <= P[0]; j++ {
			if P[j]%H[i] == 0 {
				break
			}
		}
		if j > P[0] {
			P[0]++
			P[P[0]] = H[i]
		}
	}
	if P[0] == 0 {
		P[0]++
		P[P[0]] = 1
	}
	return P[1 : P[0]+1]
}

func Graph(n uint64, A *[2000]uint64) {
	var i, j uint64
	P := SearchDel(n)
	for i = 0; i < uint64(len(P)); i++ {
		fmt.Printf("\t%v -- %v\n", n, P[i])
	}
	for i = 0; i < uint64(len(P)); i++ {
		for j = 1; j <= A[0]; j++ {
			if A[j] == P[i] {
				P[i] = 0
				break
			}
		}
		if j > A[0] {
			A[0]++
			A[A[0]] = P[i]
		}
	}
	for i = 0; i < uint64(len(P)); i++ {
		if P[i] != 0 {
			fmt.Printf("\t%v\n", P[i])
			if P[i] != 1 {
				Graph(P[i], A)
			}
		}
	}
}
func main() {
	var n uint64
	fmt.Scanf("%v", &n)
	if n == 1 {
		fmt.Printf("graph {\n\t1\n}")
	} else {
		fmt.Printf("graph {\n\t%v\n", n)
		A := [2000]uint64{0}
		Graph(n, &A)
		fmt.Printf("}")
	}
}