package medium

/*
	208. Implement Trie (Prefix Tree)
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	p := t

	for _, c := range word {
		k := c - 'a'

		if p.children[k] == nil {
			p.children[k] = &Trie{}
		}

		p = p.children[k]
	}

	p.isEnd = true
}

func (t *Trie) Search(word string) bool {
	p := t

	for _, c := range word {
		k := c - 'a'

		if next := p.children[k]; next != nil {
			p = next
		} else {
			return false
		}
	}

	return p.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	p := t

	for _, c := range prefix {
		k := c - 'a'

		if next := p.children[k]; next != nil {
			p = next
		} else {
			return false
		}
	}

	return true
}

func (t *Trie) SearchWithDots(word string) bool {
	p := t

	for i, c := range word {
		k := c - 'a'

		if c == '.' {
			found := false
			for _, v := range p.children {
				if v != nil && v.SearchWithDots(word[i+1:]) {
					found = true
				}
			}

			return found
		} else if next := p.children[k]; next != nil {
			p = next
		} else {
			return false
		}
	}

	return p.isEnd
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
