package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	key        int
	prev, next *ListNode
}

type DoublyLinkedList struct {
	nil  *ListNode
	tail *ListNode
}

func (d *DoublyLinkedList) Init() {
	d.nil = &ListNode{}
	d.tail = &ListNode{}
	d.nil.next = d.tail
	d.tail.prev = d.nil
}

func (d *DoublyLinkedList) Insert(n *ListNode) {
	n.prev = d.nil
	n.next = d.nil.next
	if d.nil.next != nil {
		d.nil.next.prev = n
	}
	d.nil.next = n
}

func (d *DoublyLinkedList) DeleteFirst() {
	if d.nil.next.next != nil {
		d.nil.next = d.nil.next.next
	}
	d.nil.next.prev = d.nil
}

func (d *DoublyLinkedList) DeleteLast() {
	if d.tail.prev.prev != nil {
		d.tail.prev = d.tail.prev.prev
	}
	d.tail.prev.next = d.tail
}

func (d *DoublyLinkedList) Delete(n *ListNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (d *DoublyLinkedList) DeleteByKey(key int) {
	cur := d.nil.next
	for cur.next != nil {
		if cur.key == key {
			d.Delete(cur)
			break
		}
		cur = cur.next
	}
}

func (d *DoublyLinkedList) Print() {
	cur := d.nil.next
	for cur.next != nil {
		if cur.next.next == nil {
			fmt.Printf("%d", cur.key)
			break
		}
		fmt.Printf("%d ", cur.key)
		cur = cur.next
	}
	fmt.Println()
}

const (
	INSERT    = "insert"
	DEL       = "delete"
	DEL_FIRST = "deleteFirst"
	DEL_LAST  = "deleteLast"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 65536) // max 65536. change if input may be longer
	sc.Scan()
	sc.Text()
	list := DoublyLinkedList{}
	list.Init()
	for sc.Scan() {
		in := strings.Split(sc.Text(), " ")
		command := in[0]
		switch command {
		case INSERT:
			key, _ := strconv.Atoi(in[1])
			node := &ListNode{
				key: key,
			}
			list.Insert(node)
		case DEL:
			key, _ := strconv.Atoi(in[1])
			list.DeleteByKey(key)
		case DEL_FIRST:
			list.DeleteFirst()
		case DEL_LAST:
			list.DeleteLast()
		}
	}
	list.Print()
}

//func solve(input []string) {
//	for i := 1; i < len(input); i++ {
//		in := strings.Split(input[i], " ")
//		command := in[0]
//		switch command {
//		case INSERT:
//			key, _ := strconv.Atoi(in[1])
//			node := &ListNode{
//				key: key,
//			}
//			list.Insert(node)
//		case DEL:
//			key, _ := strconv.Atoi(in[1])
//			list.DeleteByKey(key)
//		case DEL_FIRST:
//			list.DeleteFirst()
//		case DEL_LAST:
//			list.DeleteLast()
//		}
//	}
//	list.Print()
//}

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
