package main

import "testing"

func Test(t *testing.T) {
	solve([]int{0, 0, 0, 1, 1, 2, 2, 3, 3, 3, 3, 5, 6, 6, 6, 8, 8, 9, 9, 9},
		[]int{9, 6, 2, 4, 0, 5, 1, 3, 7, 8})
}
