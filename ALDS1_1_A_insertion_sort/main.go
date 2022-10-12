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
	insertionSort(input)
}

func insertionSort(input []int) []int {
	printIntSlice(input)
	for i := 1; i < len(input); i++ {
		v := input[i]
		j := i - 1
		for j >= 0 && input[j] > v {
			input[j+1] = input[j]
			j--
		}
		input[j+1] = v
		printIntSlice(input)
	}
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
