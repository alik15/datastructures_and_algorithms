package main

import (
	"datastructs/linkedlist"
)

func main() {
	ll := &linkedlist.List{}
	ll.Append(9)
	ll.Append(10)
	ll.DeleteFromTail()
	ll.Display()
}
