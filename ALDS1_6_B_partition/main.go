package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	inputs := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		inputs[i] = v
	}
	solve(inputs)
}

func solve(in []int) {
	q := partition(in, 0, len(in)-1)
	res := make([]byte, 0)
	for i, v := range in {
		if i == q {
			res = append(res, fmt.Sprintf("[%d] ", v)...)
		} else {
			res = append(res, fmt.Sprintf("%d ", v)...)
		}
	}
	fmt.Println(string(res[0 : len(res)-1]))
}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
