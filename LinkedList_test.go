package d8astruct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Print(t *testing.T) {
	ll := LinkedList{}
	n1 := &llNode{data: "Hello, World"}
	n2 := &llNode{data: 412}

	ll.First(n1)
	ll.First(n2)
	fLength := ll.Length()
	fmt.Println(fLength)
	ll.Print()
	nLength := ll.Length()
	fmt.Println(nLength)
}

func TestLinkedList_AddFirst(t *testing.T) {
	ll := LinkedList{}
	n1 := &llNode{data: "Hello, World"}
	n2 := &llNode{data: 412}

	ll.First(n1)
	ll.First(n2)
	assert.Equal(t, n2, ll.GetFirst())
}

func TestLinkedList_Add(t *testing.T) {
	ll := LinkedList{}
	ll.Add(412)
	ll.Add("Hello, World")
	assert.Equal(t, "Hello, World", ll.GetLast())
}

func TestLinkedList_GetFirst(t *testing.T) {

	t.Run("Assert GetFirst returns nil with empty list", func(t *testing.T) {
		ll := LinkedList{}

		var got = ll.GetFirst()
		assert.Nil(t, got)
	})

	t.Run("Assert GetFirst gets the correct item", func(t *testing.T) {
		ll := LinkedList{}
		n1 := &llNode{data: "Hello, World"}
		n2 := &llNode{data: 412}

		ll.First(n1)
		ll.First(n2)
		assert.Equal(t, n2, ll.GetFirst())
	})
}

func TestLinkedList_GetLast(t *testing.T) {

	t.Run("Assert GetLast returns nil with empty list", func(t *testing.T) {
		ll := LinkedList{}
		var got = ll.GetLast()
		assert.Nil(t, got)
	})

	t.Run("Assert GetLast gets the correct item", func(t *testing.T) {
		ll := LinkedList{}
		n1 := &llNode{data: "Hello, World"}
		n2 := &llNode{data: 412}
		n3 := &llNode{data: 121}

		ll.First(n1)
		ll.First(n2)
		ll.First(n3)

		assert.Equal(t, n1, ll.GetLast())
	})
}
