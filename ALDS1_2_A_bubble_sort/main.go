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
	BubbleSort(input)
}

func BubbleSort(input []int) []int {
	flg := true
	swapCnt := 0
	for flg {
		flg = false
		for j := len(input) - 1; j > 0; j-- {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
				swapCnt++
				flg = true
			}
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
