package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input := inputToIntSlice()
	in := input[1:]
	fmt.Println(solve(in))
}

func solve(input []int) int {
	max := math.MinInt32
	min := input[0]
	for i := 1; i < len(input); i++ {
		if max < input[i]-min {
			max = input[i] - min
		}
		if min > input[i] {
			min = input[i]
		}
	}
	return max
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
