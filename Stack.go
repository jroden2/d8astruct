package d8astruct

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

// EmptyStack removes all items from the stack
func (s *Stack) EmptyStack() {
	*s = make([]Item, 0)
}

// Push adds a new item to the Stack
func (s *Stack) Push(newItem Item) {
	*s = append(*s, newItem)
}

// Peek returns the last item from the Stack but does not remove it
func (s *Stack) Peek() Item {
	if s.IsEmpty() {
		return nil
	} else {
		i := (len(*s) - 1)
		item := (*s)[i]
		return item
	}
}

// Pop returns the last item from the Stack and removes it
func (s *Stack) Pop() Item {
	if s.IsEmpty() {
		return nil
	} else {
		i := (len(*s) - 1)
		item := (*s)[i]
		*s = (*s)[:i]
		return item
	}
}

// First returns the first item from the Stack, removing it from the Stack
func (s *Stack) First() Item {
	if s.IsEmpty() {
		return nil
	} else {
		item := (*s)[0]
		*s = (*s)[1:]
		return item
	}
}
