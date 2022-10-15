package main

import "testing"

func TestLinkedList(t *testing.T) {
	list := DoublyLinkedList{}
	list.Init()
	list.Print()
	list.Insert(&ListNode{
		key:  1,
		prev: nil,
		next: nil,
	})
	list.Print()
	n := &ListNode{
		key:  2,
		prev: nil,
		next: nil,
	}
	list.Insert(n)
	list.Print()
	list.Insert(&ListNode{
		key:  3,
		prev: nil,
		next: nil,
	})
	list.Print()
	list.DeleteFirst()
	list.Print()
	list.Insert(&ListNode{
		key:  4,
		prev: nil,
		next: nil,
	})
	list.Print()
	list.DeleteLast()
	list.Print()
	list.Delete(n)
	list.Print()
	list.DeleteLast()
	list.Print()
}

//func Test(t *testing.T) {
//	solve([]string{"7", "insert 5", "insert 2", "insert 3", "insert 1", "delete 3", "insert 6", "delete 5"})
//}
