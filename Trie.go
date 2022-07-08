package d8astruct

import "strings"

type node struct {
	children  [TRIE_NODE_COUNT]*node
	endOfWord bool
}

type Trie struct {
	root *node
}

func (t *Trie) Insert(input string) {
	inputLen := len(input)
	cur := t.root
	for i := 0; i < inputLen; i++ {
		index := input[i] - 'a'

		if cur.children[index] == nil {
			cur.children[index] = &node{}
		}

		cur = cur.children[index]
	}
	cur.endOfWord = true
}

func (t *Trie) Find(input string) bool {
	return t.search(strings.ToLower(input))
}

func (t *Trie) search(input string) bool {
	inputLen := len(input)
	cur := t.root
	for i := 0; i < inputLen; i++ {
		index := input[i] - 'a'
		if cur.children[index] == nil {
			return false
		}
		cur = cur.children[index]
	}
	if cur.endOfWord {
		return true
	}
	return false
}
