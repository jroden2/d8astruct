package d8astruct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_EmptyStack(t *testing.T) {
	t.Run("Testing that a stack size does change with EmptyStack", func(t *testing.T) {
		var stack Stack

		stack.Push("Hello")
		stack.Push(',')
		stack.Push(" World")
		initStackSize := stack.Length()

		stack.EmptyStack()
		newStackSize := stack.Length()

		if newStackSize == initStackSize {
			assert.Fail(t, "EmptyStack did not remove items from stack")
			return
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("Testing that a stack size does change with push", func(t *testing.T) {
		var stack Stack

		initStackSize := stack.Length()
		stack.Push("Hello")
		newStackSize := stack.Length()

		if newStackSize == initStackSize {
			assert.Fail(t, "Push did not add an item to stack")
			return
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("Testing that a peek does not run if stack is empty", func(t *testing.T) {
		var stack Stack

		initStackSize := stack.Length()
		fmt.Println(stack.Peek())
		newStackSize := stack.Length()

		if newStackSize != initStackSize {
			assert.Fail(t, "Peek incorrectly removed an item from stack")
			return
		}
	})
	t.Run("Testing that a stack size does not change with peek", func(t *testing.T) {
		var stack Stack

		stack.Push("Hello")
		stack.Push(',')
		stack.Push(" World")

		initStackSize := stack.Length()
		fmt.Println(stack.Peek())
		newStackSize := stack.Length()

		if newStackSize != initStackSize {
			assert.Fail(t, "Peek incorrectly removed an item from stack")
			return
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("Testing that a pop does not run if stack is empty", func(t *testing.T) {
		var stack Stack

		initStackSize := stack.Length()
		fmt.Println(stack.Pop())
		newStackSize := stack.Length()

		if newStackSize != initStackSize {
			assert.Fail(t, "pop incorrectly removed an item from stack")
			return
		}
	})
	t.Run("Testing that a stack size does change with pop", func(t *testing.T) {
		var stack Stack

		stack.Push("Hello")
		stack.Push(',')
		stack.Push(" World")

		initStackSize := stack.Length()
		fmt.Println(stack.Pop())
		newStackSize := stack.Length()

		if newStackSize == initStackSize {
			assert.Fail(t, "Pop did not remove an item from stack")
			return
		}
	})
}

func TestStack_First(t *testing.T) {
	t.Run("Testing that a First (Pop from front) does not run if stack is empty", func(t *testing.T) {
		var stack Stack

		initStackSize := stack.Length()
		fmt.Println(stack.First())
		newStackSize := stack.Length()

		if newStackSize != initStackSize {
			assert.Fail(t, "First incorrectly removed an item from stack")
			return
		}
	})
	t.Run("Testing that a stack size does change with pop", func(t *testing.T) {
		var stack Stack

		stack.Push("Hello")
		stack.Push(',')
		stack.Push(" World")

		initStackSize := stack.Length()
		got := stack.First()
		newStackSize := stack.Length()

		if newStackSize == initStackSize {
			assert.Fail(t, "Pop did not remove an item from stack")
			return
		}

		assert.Equal(t, "Hello", got)
	})
}
