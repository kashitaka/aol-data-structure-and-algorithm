package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	char string
	num  int
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	cards1 := make([]*card, n)
	cards2 := make([]*card, n)
	for i := range cards1 {
		sc.Scan()
		char := sc.Text()
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		cards1[i] = &card{
			char: char,
			num:  num,
		}
		cards2[i] = &card{
			char: char,
			num:  num,
		}
	}
	solve(cards1, cards2)
}

func solve(in1, in2 []*card) {
	quicksort(in1, 0, len(in1)-1)
	mergeSort(in2, 0, len(in2))
	noStable := diff(in1, in2)
	if noStable {
		fmt.Println("Not stable")
	} else {
		fmt.Println("Stable")
	}
	for _, v := range in1 {
		fmt.Println(v.char, v.num)
	}
}

func quicksort(A []*card, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quicksort(A, p, q-1)
		quicksort(A, q+1, r)
	}
}

func partition(A []*card, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j].num <= x.num {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func diff(A, B []*card) bool {
	for i, v := range A {
		if v.char != B[i].char {
			return true
		}
	}
	return false
}

var cnt = 0

const MAX_NUM int = 10e9

var L, R = make([]*card, 500000+2), make([]*card, 500000+2)

func mergeSort(in []*card, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(in, left, mid)
		mergeSort(in, mid, right)
		merge(in, left, mid, right)
	}
}

func merge(in []*card, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	for i := 0; i < n1; i++ {
		L[i] = in[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = in[mid+i]
	}
	init := &card{"", MAX_NUM}
	L[n1] = init
	R[n2] = init
	i, j := 0, 0
	for k := left; k < right; k++ {
		if L[i].num <= R[j].num {
			in[k] = L[i]
			i++
		} else {
			in[k] = R[j]
			j++
		}
	}
	cnt += right - left
}
