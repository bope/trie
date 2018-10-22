package trie

import "sync"

type Trie struct {
	*sync.RWMutex
	Root  *Node
	Count int
}

// New creates a new Trie
func New() *Trie {
	return &Trie{&sync.RWMutex{}, newNode(), 0}
}

// Insert a document with associated id and text
func (trie *Trie) Insert(id string, doc interface{}, text string) {
	tokens := Tokenize(text)
	trie.Lock()
	defer trie.Unlock()
	for _, token := range tokens {
		trie.Count += trie.Root.insert(id, doc, Runify(token))
	}
}

// Remove all documents by id
func (trie *Trie) Remove(id string) {
	trie.Lock()
	defer trie.Unlock()
	trie.Count -= trie.Root.remove(id)
}

// Find node by word
func (trie *Trie) Find(word string) *Node {
	trie.RLock()
	defer trie.RUnlock()
	return trie.Root.find(Runify(word))
}
