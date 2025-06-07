package main

import "fmt"

/**
Write a generic singly linked list data type. Each element can hold a comparable
value and has a pointer to the next element in the list. The methods to implement
are as follows:

// adds a new element to the end of the linked list
Add(T)

// adds an element at the specified position in the linked list
Insert(T, int)

// returns the position of the supplied value, -1 if it's not present
Index (T) int
**/

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func (l *List[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	if pos <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	curNode := l.Head
	for i := 1; i < pos; i++ {
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}
	return
}

func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &List[int]{}

	l.Add(5)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println("-----------")

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))
	fmt.Println("-----------")

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))
	fmt.Println("-----------")

}
