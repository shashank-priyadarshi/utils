package algo

import "github.com/pkg/errors"

// Trie represents the Trie data structure
type Trie[K comparable, V any] struct {
	root *TrieNode[K, V]
}

// TrieNode represents a node in the Trie
type TrieNode[K comparable, V any] struct {
	children map[K]*TrieNode[K, V]
	isWord   bool
	value    V
}

// NewTrie initializes and returns a new Trie
func NewTrie[K comparable, V any]() *Trie[K, V] {
	return &Trie[K, V]{root: &TrieNode[K, V]{children: make(map[K]*TrieNode[K, V])}}
}

func (t *Trie[K, V]) Insert(key []K, value V) {
	node := t.root

	for _, k := range key {
		child, exists := node.children[k]
		if !exists {
			child = &TrieNode[K, V]{children: make(map[K]*TrieNode[K, V])}
			node.children[k] = child
		}

		node = child
	}

	node.isWord = true
	node.value = value
}

func (t *Trie[K, V]) Search(key []K) (V, bool) {
	node := t.root

	for _, k := range key {
		child, exists := node.children[k]
		if !exists {
			var zeroValue V
			return zeroValue, false
		}

		node = child
	}

	if node.isWord {
		return node.value, true
	}

	var zeroValue V
	return zeroValue, false
}

func (t *Trie[K, V]) Delete(key []K) error {
	return t.deleteRecursive(t.root, key, 0)
}

func (t *Trie[K, V]) deleteRecursive(node *TrieNode[K, V], key []K, depth int) error {
	if depth == len(key) {
		if !node.isWord {
			return errors.New("key not found")
		}

		node.isWord = false
		if len(node.children) == 0 {
			return nil // Node can be deleted by the parent
		}

		return nil
	}

	k := key[depth]
	child, exists := node.children[k]
	if !exists {
		return errors.New("key not found")
	}

	err := t.deleteRecursive(child, key, depth+1)
	if err != nil {
		return err
	}

	// Cleanup empty nodes
	if len(child.children) == 0 && !child.isWord {
		delete(node.children, k)
	}

	return nil
}
