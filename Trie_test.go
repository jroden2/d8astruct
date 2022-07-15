package d8astruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	t.Run("Test that a value is added to the trie", func(t *testing.T) {
		trie := &Trie{
			root: &trieNode{},
		}

		words := []string{"hello", "github", "are", "you", "enjoying", "this", "Repo"}
		for _, word := range words {
			trie.Insert(word)
		}

		found := trie.Find("repo")
		if found {
			assert.True(t, true, "Found string in trie")
		}
	})
	t.Run("Test that cannot find string missing", func(t *testing.T) {
		trie := &Trie{
			root: &trieNode{},
		}

		words := []string{"hello", "github", "are", "you", "enjoying", "this", "repo"}
		for _, word := range words {
			trie.Insert(word)
		}

		found := trie.Find("yes")
		if found {
			assert.True(t, false, "String not found")
		}
	})
	t.Run("Test that cannot find a partial string", func(t *testing.T) {
		trie := &Trie{
			root: &trieNode{},
		}

		words := []string{"hello", "github", "are", "you", "enjoying", "this", "repo"}
		for _, word := range words {
			trie.Insert(word)
		}

		found := trie.Find("Rep")
		if found {
			assert.True(t, false, "String not found")
		}
	})
}
