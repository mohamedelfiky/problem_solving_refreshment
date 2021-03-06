package trie

type Trie struct {
	children map[byte]*Trie
	isLeaf   bool
}

func (t *Trie) GetChildrenNodes() map[byte]*Trie { return t.children }
func (t *Trie) GetChildren() []string {
	res, i := make([]string, len(t.children)), 0
	for k, _ := range t.children {
		res[i] = string(k)
		i++
	}
	return res
}

func NewTrie() *Trie {
	return &Trie{children: make(map[byte]*Trie)}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	arr := []byte(word)

	if _, ok := t.children[arr[0]]; !ok {
		s := NewTrie()
		if len(word) == 1 {
			s.isLeaf = true
		}
		t.children[arr[0]] = s
	}
	t.children[arr[0]].Insert(word[1:])
}

func (t *Trie) WordExist(word string) bool {
	for i := 0; i < len(word); i++ {
		if tr, ok := t.children[word[i]]; ok {
			t = tr
			if t.isLeaf && len(word) == i+1 {
				return true
			}
		} else {
			break
		}
	}
	return false
}

func (t *Trie) PrefixExist(word string) bool {
	for i := 0; i < len(word); i++ {
		if tr, ok := t.children[word[i]]; ok {
			t = tr
			if string(word[i]) == word[len(word)-1:] {
				return true
			}
		} else {
			break
		}
	}
	return false
}
