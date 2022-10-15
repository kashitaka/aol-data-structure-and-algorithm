package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLE = "*"
)

type stack struct {
	Slice []int
	top   int
}

func (s *stack) Push(x int) {
	s.Slice = append(s.Slice, x)
	s.top += 1
}

func (s *stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}
	res := s.Slice[s.top-1]
	s.Slice = s.Slice[:s.top-1]
	s.top -= 1
	return res, nil
}

func (s *stack) Empty() bool {
	return s.top == 0
}

func main() {
	input := inputToStrSlice()[0]
	solve(input)
}

func solve(input string) {
	stc := stack{}
	split := strings.Split(input, " ")
	for _, v := range split {
		switch v {
		case MINUS:
			n1, _ := stc.Pop()
			n2, _ := stc.Pop()
			stc.Push(n2 - n1)
		case PLUS:
			n1, _ := stc.Pop()
			n2, _ := stc.Pop()
			stc.Push(n2 + n1)
		case MULTIPLE:
			n1, _ := stc.Pop()
			n2, _ := stc.Pop()
			stc.Push(n2 * n1)
		default:
			n, _ := strconv.Atoi(v)
			stc.Push(n)
		}
	}
	res, _ := stc.Pop()
	fmt.Println(res)
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
