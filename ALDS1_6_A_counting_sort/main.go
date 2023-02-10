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
	in := make([]int32, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		in[i] = int32(num)
	}
	solve(in)
}

var k = 10000

func solve(in []int32) {
	out := countingSort(in)
	str := fmt.Sprint(out)
	fmt.Println(str[1 : len(str)-1])
}

func countingSort(a []int32) []int32 {
	// カウンタ行列 c
	c := make([]int32, k)

	// 配列aの数字の出現回数をそれぞれのindexに記録
	for _, v := range a {
		c[v]++
	}

	// i以下の数字の出現回数を数える
	// idx1から順に左のidxの値を加算したものがそれ
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	b := make([]int32, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]]--
	}
	return b

}
