package linkedlist

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Node struct {
	Data any
	Next *Node
}

type List struct {
	head *Node
	size int
}

func (list *List) append(element any) {
	newNode := &Node{Data: element, Next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (list *List) display() {
	current := list.head
	if list.head == nil {
		logrus.Warn("Empty List")
		return
	}
	for current != nil {
		// Add commas after each element enclosed by brackets
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (list *List) addToHead(element any) {
	if list.head == nil {
		logrus.Warn("List is Empty, Adding to head")
	}
	newNode := &Node{Data: element, Next: nil}
	newNode.Next = list.head
	list.head = newNode
}

func (list *List) deleteFromHead() {
	if list.head == nil {
		logrus.Error("Cannot Delete!, Empty List")
		return
	}
	list.head = list.head.Next
}

func (list *List) deleteFromTail() {
	current := list.head
	if list.head == nil {
		logrus.Error("Cannot Delete!, Empty List")
		return
	}
	if list.head.Next == nil {
		list.head = nil
		return
	}
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}
