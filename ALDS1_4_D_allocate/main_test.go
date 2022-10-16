package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(BoxCnt(12, 3, []int{1, 2, 8, 1, 9, 10, 11, 12}))
	fmt.Println(BoxCnt(10, 3, []int{11, 12, 1, 2}))

	solve(5, 3, []int{8, 1, 7, 3, 9})
}
