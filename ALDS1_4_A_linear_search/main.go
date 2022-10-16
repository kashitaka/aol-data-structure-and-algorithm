package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := inputToIntSlice()
	solve(in)
}

func solve(input [][]int) {
	s := input[1]
	t := input[3]
	target := make(map[int]struct{})
	for _, v := range t {
		target[v] = struct{}{}
	}
	res := 0
	for _, v := range s {
		_, ok := target[v]
		if ok {
			res++
			delete(target, v)
		}
	}
	fmt.Println(res)
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
