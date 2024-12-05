package internal

type trienode struct {
	data     rune
	visited  int
	children map[rune]*trienode
}

type trie struct {
	root *trienode
}

func NewTrie() *trie {
	// visited defaults to 0
	root := trienode{data: 0, children: make(map[rune]*trienode)}
	trie := trie{root: &root}
	return &trie
}

func Insert(t *trie, key []rune) {
	curr := t.root
	for _, c := range key {
		if curr.children[c] == nil {
			curr.children[c] = &trienode{data: c, children: make(map[rune]*trienode)}
		}
		curr.visited++
		curr = curr.children[c]
	}
}

func CountEntries(t *trie, key []rune) int {
	var res int
	curr := t.root
	for _, c := range key {
		if curr.children[c] == nil {
			res = 0
			break
		}
		res = curr.visited
		curr = curr.children[c]
	}
	return res
}
