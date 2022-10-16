package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	StartIdx int
	Areas    int
}

type Stack struct {
	Slice []int
	top   int
}

func (s *Stack) Push(x int) {
	s.Slice = append(s.Slice, x)
	s.top += 1
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("stack is empty")
	}
	res := s.Slice[s.top-1]
	s.Slice = s.Slice[:s.top-1]
	s.top -= 1
	return res, nil
}

func (s *Stack) Empty() bool {
	return s.top == 0
}

type ResultStack struct {
	Slice []Item
	top   int
}

func (s *ResultStack) Push(x Item) {
	s.Slice = append(s.Slice, x)
	s.top += 1
}

func (s *ResultStack) Pop() (Item, error) {
	if s.Empty() {
		return Item{}, fmt.Errorf("stack is empty")
	}
	res := s.Slice[s.top-1]
	s.Slice = s.Slice[:s.top-1]
	s.top -= 1
	return res, nil
}

func (s *ResultStack) Empty() bool {
	return s.top == 0
}

func main() {
	input := inputToStrSlice()[0]
	solve(input)
}

func solve(input string) {
	stack1 := Stack{}
	stack2 := ResultStack{}

	for i := 0; i < len(input); i++ {
		if input[i] == 92 {
			stack1.Push(i)
		} else if input[i] == 47 {
			if !stack1.Empty() {
				popped, err := stack1.Pop()
				if err != nil {
					fmt.Errorf("empty stack1")
					return
				}
				// pair成立 popped:i のpair
				if stack2.Empty() {
					stack2.Push(Item{
						StartIdx: popped,
						Areas:    i - popped,
					})
					continue
				}
				areas := 0
				for !stack2.Empty() {
					popped2, _ := stack2.Pop()
					if popped2.StartIdx < popped {
						// 終了
						stack2.Push(popped2)
						break
					} else {
						areas += popped2.Areas
					}
				}
				stack2.Push(Item{
					StartIdx: popped,
					Areas:    areas + i - popped,
				})
			}
		}
	}
	res := 0
	message := fmt.Sprintf("%d", len(stack2.Slice))
	for _, v := range stack2.Slice {
		res += v.Areas
		message += fmt.Sprintf(" %d", v.Areas)
	}
	fmt.Println(res)
	fmt.Println(message)
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
