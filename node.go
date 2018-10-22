package trie

type Node struct {
	// Child nodes mapped by next char in token
	Children map[rune]*Node

	// Stored data in node (also signifies that node is en of a token)
	Buckets map[string]interface{}

	// Bitmask of word length in this node and its children
	Mask int

	// Node depth in trie
	Depth uint
}

func newNode() *Node {
	return &Node{
		make(map[rune]*Node),
		make(map[string]interface{}),
		0,
		0,
	}
}

func (node *Node) newChild(r rune) *Node {
	n := newNode()
	n.Depth = node.Depth + 1
	node.Children[r] = n
	return n
}

func (node *Node) insert(id string, doc interface{}, runes []rune) int {
	size := uint(len(runes))
	node.Mask |= 1 << size
	count := 0

	current := node
	for _, r := range runes {
		if found, ok := current.Children[r]; ok {
			current = found
		} else {
			count += 1
			current = current.newChild(r)
		}
		current.Mask |= 1 << size
	}
	current.Buckets[id] = doc
	return count
}

func (node *Node) remove(id string) int {
	count := 0
	mask := 0
	for r, child := range node.Children {
		child.remove(id)
		if len(child.Children) == 0 && len(child.Buckets) == 0 {
			delete(node.Children, r)
			count += 1
		} else {
			mask |= child.Mask
		}
	}

	if _, ok := node.Buckets[id]; ok {
		delete(node.Buckets, id)
	}

	if len(node.Buckets) != 0 {
		mask |= 1 << node.Depth
	}

	return count
}

func (node *Node) find(runes []rune) *Node {
	var ok bool
	current := node
	for _, r := range runes {
		current, ok = current.Children[r]
		if !ok {
			return nil
		}
	}

	if len(current.Buckets) == 0 {
		return nil
	}

	return current
}
