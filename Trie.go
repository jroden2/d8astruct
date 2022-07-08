package d8astruct

type Word string

func (w *Word) Length() int {
	return len(*w)
}

type node struct {
	children  [TRIE_NODE_COUNT]*node
	endOfWord bool
}

type Trie struct {
	root *node
}

func (t *Trie) Insert(input Word) {
	inputLen := input.Length()
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

func (t *Trie) Find(input Word) bool {
	inputLen := input.Length()
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
