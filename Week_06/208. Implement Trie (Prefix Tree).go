package Week_06

type Trie struct {
	end bool
	next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for _, single := range word {
		if this.next[single - 'a'] == nil {
			this.next[single - 'a'] = new(Trie)
		}
		this = this.next[single - 'a']
	}
	this.end = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _,single := range word {
		if this.next[single - 'a'] == nil {
			return false
		}
		this = this.next[single - 'a']
	}
	return this.end
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _,single := range prefix {
		if this.next[single - 'a'] == nil {
			return false
		}
		this = this.next[single - 'a']
	}
	return true
}
