package search

import (
	"sort"
	"testing"
)

func TestNodeHits_Sort(t *testing.T) {
	hits := NodeHits{&NodeHit{nil, 0.5}, &NodeHit{nil, 0.7}}
	sort.Sort(hits)

	if hits[0].Score != 0.7 {
		t.Errorf("invalid sort")
	}
}


func TestBucketHits_Search(t *testing.T) {
	hits := BucketHits{&BucketHit{"", nil, 0.5}, &BucketHit{"", nil, 0.7}}
	sort.Sort(hits)

	if hits[0].Score != 0.7 {
		t.Errorf("invalid sort")
	}
}
