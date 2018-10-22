package trie

import "testing"

func testEqStrings(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestTokenize(t *testing.T) {
	text := "one TWO 3three! @FoUr åäö©"
	expected := []string{
		"one",
		"two",
		"3three",
		"four",
		"åäö",
	}

	tokens := Tokenize(text)
	if !testEqStrings(tokens, expected) {
		t.Errorf("incorrect tokens, got: %v, want: %v", tokens, expected)
	}
}
