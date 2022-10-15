package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := inputToIntSlice()[1]
	SelectionSort(input)
}

func SelectionSort(input []int) []int {
	swapCnt := 0
	for i := range input {
		minIdx := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIdx] {
				minIdx = j
			}
		}
		if i != minIdx {
			input[i], input[minIdx] = input[minIdx], input[i]
			swapCnt++
		}
	}
	printIntSlice(input)
	fmt.Println(swapCnt)
	return input
}

func printIntSlice(input []int) {
	s := fmt.Sprint(input)
	fmt.Println(s[1 : len(s)-1])
}

func inputToIntSlice() [][]int {
	var res [][]int
	var sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 65536) // max 65536. change if input may be longer
	for sc.Scan() {
		line := sc.Text()
		ss := strings.Split(line, " ")
		intSlice := make([]int, len(ss))
		for i, v := range ss {
			intV, _ := strconv.Atoi(v)
			intSlice[i] = intV
		}
		res = append(res, intSlice)
	}
	return res
}
