package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	id        int
	parent    *node
	leftChild *node
	rightBro  *node
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nodes := make([]*node, n)
	for i := range nodes {
		nodes[i] = &node{}
	}
	for i := 0; i < len(nodes); i++ {
		sc.Scan()
		id, _ := strconv.Atoi(sc.Text())
		nodes[id].id = id
		sc.Scan()
		numChild, _ := strconv.Atoi(sc.Text())
		if numChild == 0 {
			continue
		}
		childIds := make([]int, numChild)
		for j := 0; j < numChild; j++ {
			sc.Scan()
			chId, _ := strconv.Atoi(sc.Text())
			childIds[j] = chId
		}
		nodes[id].leftChild = nodes[childIds[0]]
		for k, v := range childIds {
			nodes[v].parent = nodes[id]
			if k != len(childIds)-1 {
				nodes[v].rightBro = nodes[childIds[k+1]]
			}
		}
	}
	printNode(nodes)
}

func printNode(nodes []*node) {
	for _, v := range nodes {
		parentId := -1
		if v.parent != nil {
			parentId = v.parent.id
		}
		depth := getDepth(nodes, v.id)
		tp := getType(nodes, v.id)
		children := getChildIds(nodes, v.id)
		fmt.Printf("node %d: parent = %d, depth = %d, %s, %s\n", v.id, parentId, depth, tp, getStringIds(children))
	}
}

func getDepth(nodes []*node, i int) int {
	res := 0
	id := nodes[i].id
	for true {
		if nodes[id].parent == nil {
			break
		}
		id = nodes[id].parent.id
		res++
	}
	return res
}

func getType(nodes []*node, i int) string {
	if nodes[i].parent == nil {
		return "root"
	}
	if nodes[i].leftChild == nil {
		return "leaf"
	}
	return "internal node"
}

func getChildIds(nodes []*node, i int) []int {
	var res []int
	if nodes[i].leftChild == nil {
		return res
	}
	id := nodes[i].leftChild.id
	for true {
		res = append(res, id)
		if nodes[id].rightBro == nil {
			break
		}
		id = nodes[id].rightBro.id
	}
	return res
}

func getStringIds(in []int) string {
	var res = make([]byte, 0, len(in)*3)
	res = append(res, '[')
	for i, v := range in {
		res = append(res, fmt.Sprint(v)...)
		if i != len(in)-1 {
			res = append(res, ", "...)
		}
	}
	res = append(res, ']')
	return string(res)
}
