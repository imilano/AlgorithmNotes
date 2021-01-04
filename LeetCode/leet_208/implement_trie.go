package leet_208

/*
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
*/

//应用： 词频统计、前缀匹配、字符串排序。
//特点： 	- 根节点不包含字符，除根节点外每个节点只包含一个字符
//		- 从跟节点到某一个节点，路径上经过的字符连接起来，为该节点对应的字符串
//		- 每个节点包含的子节点包含的字符串各不相同。
//
//如何表示每个节点的子节点有哪些呢？可以在每个节点上开一个数组，用于表示该字符指向的子节点；或者在每个节点上开一个链表，按一定顺序存储与子节点的对应关系。
type Trie struct {
	isWord bool
	next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isWord: false,
		next: [26]*Trie{},
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	p := this
	for i := range word {
		idx := word[i] - 'a'
		if p.next[idx] == nil {
			p.next[idx] = new(Trie)
		}

		p = p.next[idx]
	}

	p.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this
	for i := range word {
		idx := word[i] - 'a'
		if p.next[idx] == nil {
			return  false
		}

		p = p.next[idx]
	}

	return p.isWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for i := range prefix {
		idx := prefix[i] - 'a'
		if p.next[idx] == nil {
			return false
		}

		p = p.next[idx]
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
