package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	name string
	time int
}

type Queue struct {
	slice []Item
}

func (q *Queue) Enqueue(x Item) {
	q.slice = append(q.slice, x)
}

func (q *Queue) Dequeue() (Item, error) {
	if q.Empty() {
		return Item{}, fmt.Errorf("empty queue")
	}
	res := q.slice[0]
	q.slice = q.slice[1:]
	return res, nil
}

func (q *Queue) Empty() bool {
	return len(q.slice) == 0
}

func main() {
	input := inputToStrSlice()
	solve(input)
}

func solve(input []string) {
	que := Queue{}
	strQ := strings.Split(input[0], " ")[1]
	q, _ := strconv.Atoi(strQ)
	for i := 1; i < len(input); i++ {
		spl := strings.Split(input[i], " ")
		name := spl[0]
		time, _ := strconv.Atoi(spl[1])
		que.Enqueue(Item{
			name: name,
			time: time,
		})
	}

	elapsed := 0
	for !que.Empty() {
		item, _ := que.Dequeue()
		if item.time <= q {
			fmt.Printf("%s %d\n", item.name, elapsed+item.time)
			elapsed += item.time
		} else {
			que.Enqueue(Item{
				name: item.name,
				time: item.time - q,
			})
			elapsed += q
		}
	}
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
