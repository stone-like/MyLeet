/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	key       string
	children  map[rune]*Trie
	endOfWord bool
}

func Constructor() Trie {
	return Trie{
		key:       "",
		children:  make(map[rune]*Trie),
		endOfWord: true,
	}
}

func (this *Trie) Insert(word string) {

	curNode := this
	for i, r := range word {
		if curNode.children[r] == nil {
			curNode.children[r] = &Trie{
				key:      string(r),
				children: make(map[rune]*Trie),
			}
		}

		curNode = curNode.children[r]
		if i == len(word)-1 {
			curNode.endOfWord = true
		}
	}
}

func (this *Trie) SearchWithComplete(word string, isComplete bool) bool {
	curNode := this

	for _, r := range word {
		if curNode.children[r] == nil {
			return false
		}

		curNode = curNode.children[r]

	}

	if !isComplete {
		return true
	}

	return curNode.endOfWord
}

func (this *Trie) Search(word string) bool {
	return this.SearchWithComplete(word, true)
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchWithComplete(prefix, false)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

