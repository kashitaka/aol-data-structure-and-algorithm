package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	in := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(SelectionSort(in))
}
