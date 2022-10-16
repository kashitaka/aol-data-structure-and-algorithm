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
	s := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		s[i] = v
	}
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	t := make([]int, q)
	for i := 0; i < q; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		t[i] = v
	}
	solve(s, t)
}

func solve(s, t []int) {
	left := 0
	right := len(s) - 1
	cnt := 0
	for _, v := range t {
		res := binarySearch(s, v, left, right)
		if res != -1 {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func binarySearch(s []int, t, left, right int) int {
	mid := (left + right) / 2
	if s[mid] == t {
		return mid
	}

	if left > right {
		return -1
	}
	if t < s[mid] {
		right = mid - 1
		return binarySearch(s, t, left, right)
	}
	if s[mid] < t {
		left = mid + 1
		return binarySearch(s, t, left, right)
	}
	panic("impossible")
}
