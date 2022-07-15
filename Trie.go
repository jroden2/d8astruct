package d8astruct

import "strings"

type trieNode struct {
	children  [TRIE_NODE_COUNT]*trieNode
	endOfWord bool
}

// Trie is a data structure of particular use for searching words from node tip to tail
type Trie struct {
	root *trieNode
}

// Insert adds a string to the trie, but lowercases to standardize input
func (t *Trie) Insert(input string) {
	t.insert(strings.ToLower(input))
}

// insert adds the input to the trie in its raw
func (t *Trie) insert(input string) {
	inputLen := len(input)
	cur := t.root
	for i := 0; i < inputLen; i++ {
		index := input[i] - 'a'

		if cur.children[index] == nil {
			cur.children[index] = &trieNode{}
		}

		cur = cur.children[index]
	}
	cur.endOfWord = true
}

// Find wraps search but converts input to lowercase
func (t *Trie) Find(input string) bool {
	return t.search(strings.ToLower(input))
}

// search will find any string from a Trie in raw format
// however, this does not support uppercase in current guise.
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
