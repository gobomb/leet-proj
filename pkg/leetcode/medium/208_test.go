package medium

import (
	"log"
	"testing"
)

func Test_ConstructorTrie(t *testing.T) {
	trie := ConstructorTrie()
	rs := false

	trie.Insert("apple")

	rs = trie.Search("apple") // return True
	log.Println(rs)

	rs = trie.Search("app") // return False
	log.Println(rs)

	rs = trie.StartsWith("app") // return True
	log.Println(rs)

	trie.Insert("app")
	rs = trie.Search("app") // return True
	log.Println(rs)

	t1 := []string{"Trie", "insert", "search", "search", "startsWith", "startsWith", "insert", "search", "startsWith", "insert", "search", "startsWith"}

	t2 := []string{"", "ab", "abc", "ab", "abc", "ab", "ab", "abc", "abc", "abc", "abc", "abc"}

	null := false
	r1 := []bool{null, null, false, true, false, true, null, false, false, null, true, true}

	for i, e1 := range t1 {
		rs = null

		switch e1 {
		case "Trie":
			trie = ConstructorTrie()
		case "insert":
			trie.Insert(t2[i])
		case "search":
			rs = trie.Search(t2[i])
		case "startsWith":
			rs = trie.StartsWith(t2[i])
		}

		if rs != r1[i] {
			t.Fatalf("failed at %v %v %v %v\n", i, t1[i], t2[i], r1[i])
		}
	}
}
