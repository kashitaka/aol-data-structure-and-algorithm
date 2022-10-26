package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

var cnt = 0

const MAX_NUM int = 10e9

var L, R = make([]int, 500000+2), make([]int, 500000+2)

func solve(in []int) {
	cnt = 0
	mergeSort(in, 0, len(in))
	fmt.Println(strings.TrimRight(fmt.Sprintf("%+v", in)[1:], "]"))
	fmt.Println(cnt)
}

func mergeSort(in []int, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(in, left, mid)
		mergeSort(in, mid, right)
		merge(in, left, mid, right)
	}
}

func merge(in []int, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	for i := 0; i < n1; i++ {
		L[i] = in[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = in[mid+i]
	}
	L[n1] = MAX_NUM
	R[n2] = MAX_NUM
	i, j := 0, 0
	for k := left; k < right; k++ {
		if L[i] <= R[j] {
			in[k] = L[i]
			i++
		} else {
			in[k] = R[j]
			j++
		}
	}
	cnt += right - left
}

// この出力の仕方は最大のボトルネックになる
// 標準出力を送りまくる方式なのでIOが頻繁に入りよくない
// 配列を出力するときは solve でやっているような方法をとる
//func printIntSlice(in []int) {
//	for i, v := range in {
//		if i == len(in)-1 {
//			fmt.Printf("%d\n", v)
//		} else {
//			fmt.Printf("%d ", v)
//		}
//	}
//}
