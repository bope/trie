package trie

import "testing"

func TestTrie_Find(t *testing.T) {
	tr := New()
	tr.Insert("1", nil, "test1")
	tr.Insert("2", nil, "test2")
	tr.Insert("3", nil, "foobar")
	var n *Node

	if n = tr.Find("test1"); n == nil {
		t.Error("node not found")
	}

	if n = tr.Find("test2"); n == nil {
		t.Error("node not found")
	}

	if n = tr.Find("foobar"); n == nil {
		t.Error("node not found")
	}

	if n = tr.Find("test"); n != nil {
		t.Error("node found, but shouldn't")
	}

	if n = tr.Find("test3"); n != nil {
		t.Error("node found, but shouldn't")
	}

	if n = tr.Find("f"); n != nil {
		t.Error("node found, but shouldn't")
	}
}

func TestTrie_Remove(t *testing.T) {
	trie := New()
	trie.Insert("1", nil, "test1")
	trie.Insert("2", nil, "test2")
	trie.Insert("3", nil, "foobar")

	trie.Remove("2")
	var n *Node

	if n = trie.Find("test2"); n != nil {
		t.Error("node found, but shouldn't")
	}
}
