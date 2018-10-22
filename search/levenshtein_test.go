package search

import (
	"github.com/bope/trie"
	"sort"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	tr := trie.New()
	tr.Insert("1", nil, "test")
	tr.Insert("2", nil, "atest")
	tr.Insert("3", nil, "testz")
	tr.Insert("4", nil, "foobar")
	tr.Insert("5", nil, "foobaar")

	res := Levenshtein(tr, "test", 1)
	sort.Sort(res.Hits)

	if c := len(res.Hits); c != 3 {
		t.Errorf("result len != 3: %v", res)
		return
	}

	if _, ok := res.Hits[0].Node.Buckets["1"]; !ok {
		t.Errorf("res[0].Document != 1: %v", res)
		return
	}

	bres := res.BucketResult()
	sort.Sort(bres.Hits)

	if c := len(bres.Hits); c != 3 {
		t.Errorf("bresult len != 3: %v", res)
		return
	}

	if bres.Hits[0].ID != "1" {
		t.Errorf("bres[0].Document != 1: %v", res)
		return
	}

}

func TestLevenshtein2(t *testing.T) {
	tr := trie.New()
	tr.Insert("1", nil, "test meh")
	tr.Insert("2", nil, "foobar foobar")

	res := Levenshtein(tr, "test meh", 1)
	sort.Sort(res.Hits)

	if c := len(res.Hits); c != 2 {
		t.Errorf("result len != 2: %v", res)
		return
	}

	if res.Hits[0].Score != 1.0 {
		t.Errorf("res[0].Score != 2.0: %v", res)
		return
	}

	bres := res.BucketResult()
	sort.Sort(bres.Hits)

	if c := len(bres.Hits); c != 1 {
		t.Errorf("bresult len != 3: %v", res)
		return
	}

	if bres.Hits[0].ID != "1" {
		t.Errorf("bres[0].Document != 1: %v", res)
		return
	}

	if bres.Hits[0].Score != 2.0 {
		t.Errorf("bres[0].Score != 2.0: %v", res)
		return
	}
}

func BenchmarkLevenshtein1(b *testing.B) {
	tr := trie.New()
	tr.Insert("1", nil, "some text here")
	tr.Insert("2", nil, "another text here, and something else")
	tr.Insert("3", nil, "add some longerwordsherejusttotest")
	tr.Insert("4", nil, `Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the â€œgo testâ€ command, which automates execution of any function of the form`)
	tr.Insert("5", nil, `To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the â€œgo testâ€ command is run. For more detail, run â€œgo help testâ€ and â€œgo help testflagâ€.`)
	tr.Insert("6", nil, `The package also runs and verifies example code. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run. (The comparison ignores leading and trailing space.) These are examples of an example:`)

	for n := 0; n < b.N; n++ {
		Levenshtein(tr, "testing", 1)
	}

}

func BenchmarkLevenshtein2(b *testing.B) {
	t := trie.New()
	t.Insert("1", nil, "some text here")
	t.Insert("2", nil, "another text here, and something else")
	t.Insert("3", nil, "add some longerwordsherejusttotest")
	t.Insert("4", nil, `Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the â€œgo testâ€ command, which automates execution of any function of the form`)
	t.Insert("5", nil, `To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the â€œgo testâ€ command is run. For more detail, run â€œgo help testâ€ and â€œgo help testflagâ€.`)
	t.Insert("6", nil, `The package also runs and verifies example code. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run. (The comparison ignores leading and trailing space.) These are examples of an example:`)

	for n := 0; n < b.N; n++ {
		Levenshtein(t, "testing", 2)
	}

}
