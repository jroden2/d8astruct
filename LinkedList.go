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

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) GetFirst() any {
	if l.first == nil {
		return nil
	}
	return l.first.data
}

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

func (l *LinkedList) AddFirst(newNode *llNode) {
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
