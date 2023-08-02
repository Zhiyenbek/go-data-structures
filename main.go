package main

import (
	"fmt"

	linkedList "github.com/Zhiyenbek/go-data-structures/pkg/linked_list"
)

func main() {
	ll := linkedList.NewLinkedList()

	ll.Add(1)
	ll.Add(2)
	ll.Add(5)
	ll.Add(6)
	ll.Add(8)
	ll.Add(9)
	ll.Print()
	lst, ok := ll.GetLast()
	fmt.Println(lst)
	fmt.Println(ok)
	ll.RemoveFirst()
	ll.RemoveFirst()
	ll.RemoveFirst()
	ll.RemoveFirst()
	ll.RemoveFirst()
	ll.RemoveFirst()
	ll.RemoveFirst()
	ll.Print()
}
