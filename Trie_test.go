package d8astruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestWord_Length tests that a Word (string mask) is correct length when checked
func TestWord_Length(t *testing.T) {
	tests := []struct {
		name string
		w    Word
		want int
	}{
		{
			"Assert length is correct on word",
			"Test",
			4,
		},
		{
			"Assert length is correct on string with break",
			"Hello, World",
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.w.Length(), "Length()")
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	t.Run("Test that a value is added to the trie", func(t *testing.T) {
		trie := &Trie{
			root: &node{},
		}

		words := []Word{"hello", "github", "are", "you", "enjoying", "this", "repo"}
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
			root: &node{},
		}

		words := []Word{"hello", "github", "are", "you", "enjoying", "this", "repo"}
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
			root: &node{},
		}

		words := []Word{"hello", "github", "are", "you", "enjoying", "this", "repo"}
		for _, word := range words {
			trie.Insert(word)
		}

		found := trie.Find("rep")
		if found {
			assert.True(t, false, "String not found")
		}
	})
}
