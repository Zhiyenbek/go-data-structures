package LinkedList

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkedList struct {
	Head *node
}

type node struct {
	Value int64
	Next  *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(value int64) {
	node := &node{
		Value: value,
	}
	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}

}

func (l *LinkedList) GetLast() (int64, bool) {
	if l.Head == nil {
		return 0, false
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Value, true
}

func (l *LinkedList) RemoveFirst() *node {
	if l.Head == nil {
		return nil
	}
	l.Head = l.Head.Next
	return l.Head
}

func (l *LinkedList) Print() {
	output := &strings.Builder{}
	_, err := output.WriteRune('{')
	if err != nil {
		panic(err)
	}
	current := l.Head

	for {
		val := strconv.Itoa(int(current.Value))
		output.WriteString(val)
		if current.Next != nil {
			output.WriteString(", ")
			current = current.Next
		} else {
			break
		}
	}
	output.WriteRune('}')
	fmt.Println(output.String())
}
