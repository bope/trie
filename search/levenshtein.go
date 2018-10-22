package search

import "github.com/bope/trie"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func anylesseq(l []int, b int) bool {
	for _, a := range l {
		if a <= b {
			return true
		}
	}
	return false
}

func Levenshtein(root *trie.Trie, text string, cost int) *NodeResult {
	var row []int
	var size int
	var mask int

	tokens := trie.Tokenize(text)
	res := &NodeResult{make(NodeHits, 0)}

	for _, token := range tokens {
		runes := trie.Runify(token)
		size = len(runes)
		row = make([]int, size+1)
		mask = 0

		for i := uint(size - cost); i <= (uint(size + cost)); i++ {
			mask |= 1 << i
		}

		for i := 0; i <= size; i++ {
			row[i] = i
		}

		for _, r := range runes {
			if child, ok := root.Root.Children[r]; ok {
				levenshtein(child, r, runes, row, cost, mask, res)
			}
		}

	}

	return res
}

func levenshtein(node *trie.Node, char rune, runes []rune, lastRow []int, cost, mask int, res *NodeResult) {
	var inDelCost, repCost int

	if mask&node.Mask == 0 {
		return
	}

	lastSize := len(lastRow)
	row := make([]int, lastSize)
	row[0] = lastRow[0] + 1

	for i := 1; i < lastSize; i++ {
		inDelCost = min(row[i-1]+1, lastRow[i]+1)
		if runes[i-1] == char {
			repCost = lastRow[i-1]
		} else {
			repCost = lastRow[i-1] + 1
		}
		row[i] = min(inDelCost, repCost)
	}

	if len(node.Buckets) != 0 && row[lastSize-1] <= cost {
		score := float64(lastSize-row[lastSize-1]) / float64(lastSize)
		res.Hits = append(res.Hits, &NodeHit{node, score})
	}

	if anylesseq(row, cost) {
		for nextChar, child := range node.Children {
			levenshtein(child, nextChar, runes, row, cost, mask, res)
		}
	}
}
