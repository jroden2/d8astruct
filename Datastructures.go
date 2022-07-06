package main

type Item any
type Stack []Item

// Length gets the total length of items in the Stack
func (s *Stack) Length() int {
	return len(*s)
}

// IsEmpty returns a boolean value if a Stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds a new item to the Stack
func (s *Stack) Push(newItem Item) {
	*s = append(*s, newItem)
}

// Peek returns the last item from the Stack but does not remove it
func (s *Stack) Peek() (Item, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		i := (len(*s) - 1)
		item := (*s)[i]
		return item, true
	}
}

// Pop returns the last item from the Stack and removes it
func (s *Stack) Pop() (Item, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		i := (len(*s) - 1)
		item := (*s)[i]
		*s = (*s)[:i]
		return item, true
	}
}
