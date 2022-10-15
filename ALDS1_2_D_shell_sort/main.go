package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := inputToIntSlice()
	in := input[1:]
	shellSort(in)
}

func insertionSort(input []int, g int) int {
	//fmt.Printf("g: %d\n", g)
	cnt := 0
	for i := g; i < len(input); i++ {
		v := input[i]
		j := i - g
		for j >= 0 && input[j] > v {
			input[j+g] = input[j]
			cnt++
			j -= g
			//printIntSlice(input)
		}
		input[j+g] = v
	}
	return cnt
}

func shellSort(in []int) {
	var g []int
	totalCnt := 0
	i := 1
	g = append(g, i)
	i *= 2
	for i < len(in) {
		g = append(g, i)
		i = i * 2
	}
	fmt.Println(len(g))
	for j := len(g) - 1; j >= 0; j-- {
		if j != 0 {
			fmt.Printf("%d ", g[j])
		} else {
			fmt.Printf("%d\n", g[j])
		}
	}
	for j := len(g) - 1; j >= 0; j-- {
		totalCnt += insertionSort(in, g[j])
	}
	fmt.Println(totalCnt)
	for _, v := range in {
		fmt.Println(v)
	}
}

func inputToIntSlice() []int {
	var res []int
	var sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 65536) // max 65536. change if input may be longer
	for sc.Scan() {
		line := sc.Text()
		intV, _ := strconv.Atoi(line)
		res = append(res, intV)
	}
	return res
}
