package d8astruct

import (
	"fmt"
)

type llNode struct {
	data any
	next *llNode
}

type LinkedList struct {
	length int
	first  *llNode
	last   *llNode
}

// Length returns the length as an int
func (l *LinkedList) Length() int {
	return l.length
}

// GetFirst returns the first items data from the list
func (l *LinkedList) GetFirst() any {
	if l.first == nil {
		return nil
	}
	return l.first.data
}

// GetLast returns the last items data from the list
func (l *LinkedList) GetLast() any {
	if l.first == nil {
		return nil
	}
	if l.last.next == nil {
		return l.last.data
	} else {
		tList := l
		for tList.first.next != nil {
			tList.first = tList.first.next
		}
		return tList.first.data
	}
}

// Print outputs the data for each item in the linked list, removing each from the list iteration
func (l *LinkedList) Print() {
	tList := l
	for tList.first != nil {
		fmt.Printf("%v", tList.first.data)
		tList.first = tList.first.next
	}
}

// First adds an item to the front of the linked list
func (l *LinkedList) First(T any) {
	newNode := &llNode{data: T}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
		l.length++
		return
	}
	var tNode = l.first
	l.first = newNode
	l.first.next = tNode
	l.length++
}

// Add adds an item to the end of the linked list
func (l *LinkedList) Add(T any) {
	newNode := &llNode{data: T}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
		l.length++
		return
	}

	for l.first.next != nil {
		l.first = l.first.next
	}
	l.first.next = newNode
	l.last = newNode
	l.length++
}
