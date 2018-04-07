// Mars project main.go
package main

import (
        "bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

var ArrayB [][]int
var Arrayk0 []int
var Arrayk1 []int

func ReadInt32(args ...*int) int {
	scanner.Scan()
	ans, _ := strconv.Atoi(scanner.Text())
	for _, numbers := range args {
		*numbers = ans
	}
	return ans
}

func ReadSymbol(args ...*int) int {
	scanner.Scan()
	pm := scanner.Text()[0]
	var ans int
	if pm == '-' {
		ans = 0 //Совместимы
	} else {
		ans = 1 //Не соместимы
	}
	for _, ch := range args {
		*ch = ans
	}
	return ans
}

func FormingGroup(A [][]int, B []int, k0, k1, k2, pos, n int) {
	if pos == n {
		ArrayB = append(ArrayB, B)
		Arrayk0 = append(Arrayk0, k0)
		Arrayk1 = append(Arrayk1, k1)
		return
	}
	var i int
	for i = 0; i < n; i++ {
		if A[pos][i] == 1 {
			break
		}
	}
	if i == n {
		B[pos] = 0
		FormingGroup(A, B, k0+1, k1, k2, pos+1, n)
	} else {
		B1 := make([]int, n)
		B2 := make([]int, n)
		copy(B1, B)
		copy(B2, B)
		B1[pos] = 1
		B2[pos] = 2
		f1 := true
		f2 := true
		for ; i < n; i++ {
			if !(f1 || f2) {
				break
			}
			if A[pos][i] == 1 {
				if B[i] == 0 {
					B1[i] = 2
					B2[i] = 1
					continue
				}
				if B[i] == 1 {
					f1 = false
					continue
				}
				if B[i] == 2 {
					f2 = false
					continue
				}
			}
		}
		if f1 {
			FormingGroup(A, B1, k0, k1+1, k2, pos+1, n)
		}
		if f2 {
			FormingGroup(A, B2, k0, k1, k2+1, pos+1, n)
		}
	}
}

func main() {
	scanner.Split(bufio.ScanWords)

	var i, j, n int
	n = ReadInt32()
	var A [][]int
	B := make([]int, n)
	for i = 0; i < n; i++ {
		A = append(A, []int{})
		for j = 0; j < n; j++ {
			A[i] = append(A[i], ReadSymbol())
		}
	}
	FormingGroup(A, B, 0, 0, 0, 0, n)
	if len(ArrayB) == 0 {
		fmt.Print("No solution")
		return
	}
	for i = 0; i < len(ArrayB); i++ {
		Arrayk0[i] = n/2 - Arrayk1[i]
	}
	str := make([]string, len(ArrayB))
	for i = 0; i < n; i++ {
		for j = 0; j < len(ArrayB); j++ {
			if ArrayB[j][i] == 0 && Arrayk0[j] > 0 {
				ArrayB[j][i] = 1
				Arrayk0[j]--
			}
			if ArrayB[j][i] == 0 && Arrayk0[j] == 0 {
				ArrayB[j][i] = 2
			}
			str[j] += strconv.Itoa(ArrayB[j][i])
		}
	}
	minstr := str[0]
	for i = 1; i < len(ArrayB); i++ {
		if minstr > str[i] {
			minstr = str[i]
		}
	}
	for i = 0; i < len(minstr); i++ {
		if minstr[i] == '1' {
			fmt.Print(i+1, " ")
		}
	}
}
