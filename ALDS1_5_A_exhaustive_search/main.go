package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	inputs := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		inputs[i] = v
	}
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	targets := make([]int, q)
	for i := 0; i < q; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		targets[i] = v
	}
	solve(inputs, targets)
}

func solve(inputs, targets []int) {
	combination := make([]map[int]struct{}, len(inputs))
	combination[0] = map[int]struct{}{
		0:         {},
		inputs[0]: {},
	}
	for i := 1; i < len(inputs); i++ {
		previous := combination[i-1]
		thisTime := make(map[int]struct{}, len(previous)*2)
		v := inputs[i]
		for k := range previous {
			thisTime[k] = struct{}{}
			thisTime[k+v] = struct{}{}
		}
		combination[i] = thisTime
	}
	for _, v := range targets {
		_, ok := combination[len(inputs)-1][v]
		if ok {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}

}
