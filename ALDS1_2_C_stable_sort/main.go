package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Trump struct {
	Suit  string
	Value int
}

func main() {
	input := inputToStrSlice()[1]
	solve(input)
}

func solve(input string) {
	strTrumps := strings.Split(input, " ")
	trumps1 := convertToRumps(strTrumps)
	trumps2 := make([]Trump, len(trumps1))
	copy(trumps2, trumps1)
	res1 := BubbleSort(trumps1)
	res2 := SelectionSort(trumps2)
	strRes1 := printSlice(res1)
	strRes2 := printSlice(res2)
	fmt.Println(strRes1)
	fmt.Println("Stable")
	fmt.Println(strRes2)
	if strRes1 != strRes2 {
		fmt.Println("Not stable")
	} else {
		fmt.Println("Stable")
	}
}

func convertToRumps(in []string) []Trump {
	trumps := make([]Trump, len(in))
	for i, v := range in {
		intVal, _ := strconv.Atoi(v[1:2])
		trumps[i] = Trump{
			Suit:  v[0:1],
			Value: intVal,
		}
	}
	return trumps
}

func BubbleSort(input []Trump) []Trump {
	flg := true
	for flg {
		flg = false
		for j := len(input) - 1; j > 0; j-- {
			if input[j].Value < input[j-1].Value {
				input[j], input[j-1] = input[j-1], input[j]
				flg = true
			}
		}
	}
	return input
}

func SelectionSort(input []Trump) []Trump {
	for i := range input {
		minIdx := i
		for j := i + 1; j < len(input); j++ {
			if input[j].Value < input[minIdx].Value {
				minIdx = j
			}
		}
		if i != minIdx {
			input[i], input[minIdx] = input[minIdx], input[i]
		}
	}
	return input
}

func printSlice(input []Trump) string {
	res := ""
	for i, v := range input {
		if i != len(input)-1 {
			res += fmt.Sprintf("%s%d ", v.Suit, v.Value)
		} else {
			res += fmt.Sprintf("%s%d", v.Suit, v.Value)
		}
	}
	return res
}

func inputToStrSlice() []string {
	var res []string
	var sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 65536) // max 65536. change if input may be longer
	for sc.Scan() {
		res = append(res, sc.Text())
	}
	return res
}
