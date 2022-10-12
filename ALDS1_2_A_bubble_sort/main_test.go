package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	in := []int{5, 3, 2, 4, 1}
	fmt.Println(BubbleSort(in))
}
