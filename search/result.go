package search

import "github.com/bope/trie"

// NodeHit stores a node search hit with its score
type NodeHit struct {
	// Node
	Node *trie.Node

	// Search score
	Score float64
}

// NodeHits implements sort interface
type NodeHits []*NodeHit

func (h NodeHits) Len() int {
	return len(h)
}

func (h NodeHits) Less(i, j int) bool {
	// todo: secondary order
	return h[i].Score > h[j].Score
}

func (h NodeHits) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// NodeResult stores NodeHits from a search
type NodeResult struct {
	Hits NodeHits
}

// BucketHit stores a Bucket from a search result
type BucketHit struct {
	// Bucket id
	ID string

	// Bucket data
	Bucket interface{}

	// Score from search
	Score float64
}

// BucketHits implements sort interface
type BucketHits []*BucketHit

func (h BucketHits) Len() int {
	return len(h)
}

func (h BucketHits) Less(i, j int) bool {
	// todo: secondary order
	return h[i].Score > h[j].Score
}

func (h BucketHits) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// BucketReult returns NodeResult converted to a BucketResult
func (nr *NodeResult) BucketResult() *BucketResult {
	bmap := make(map[string]*BucketHit)

	for _, nh := range nr.Hits {
		for id, bucket := range nh.Node.Buckets {
			if _, ok := bmap[id]; !ok {
				bmap[id] = &BucketHit{id, bucket, 0}
			}
			bmap[id].Score += nh.Score
		}
	}

	ret := &BucketResult{make(BucketHits, 0, len(bmap))}
	for _, bh := range bmap {
		ret.Hits = append(ret.Hits, bh)
	}

	return ret
}

// BucketResult stores BucketHits from a search result
type BucketResult struct {
	Hits BucketHits
}
