package trie

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	text = strings.ToLower(text)
	text = strings.Map(func(r rune) rune {
		switch {
		case unicode.IsSpace(r):
			return ' '
		case unicode.IsPunct(r):
			return ' '
		case unicode.IsSymbol(r):
			return -1
		}
		return r
	}, text)
	parts := strings.Split(text, " ")
	ret := make([]string, 0)
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		ret = append(ret, p)
	}
	return ret
}

func Runify(word string) []rune {
	return []rune("$" + word)
}
