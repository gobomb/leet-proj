package medium

/*
	211. Design Add and Search Words Data Structure
*/

type WordDictionary struct {
	t Trie
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{ConstructorTrie()}
}

func (wd *WordDictionary) AddWord(word string) {
	wd.t.Insert(word)
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.t.SearchWithDots(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
