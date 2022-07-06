package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			assert.Fail(t, "Peek incorrectly removed an item from stack")
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
