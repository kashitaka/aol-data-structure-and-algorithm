package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

}

// inputToStrSlice returns slice of string
// each element is number of lines
// e.g.
// if stdin is:
// "ア イ ウ"
// "カ キ ク"
// returned value is
// []string{"ア イ ウ","カ キ ク"}
func inputToStrSlice() []string {
	var res []string
	var sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 65536) // max 65536. change if input may be longer
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
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
