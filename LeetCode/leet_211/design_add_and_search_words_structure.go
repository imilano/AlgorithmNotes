package leet_211

/*
Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.


Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True


Constraints:

1 <= word.length <= 500
word in addWord consists lower-case English letters.
word in search consist of  '.' or lower-case English letters.
At most 50000 calls will be made to addWord and search.
*/

type WordDictionary struct {
	isChild bool
	next    [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		isChild: false,
		next:    [26]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	tmp := this
	for i := range word {
		index := word[i] - 'a'
		if tmp.next[index] == nil {
			tmp.next[index] = new(WordDictionary)
		}
		tmp = tmp.next[index]
	}

	tmp.isChild = true
}

func (this *WordDictionary) Search(word string) bool {
	return search(this, word, 0)
}

func search(dic *WordDictionary, word string, idx int) bool {
	if len(word) == idx {
		return dic.isChild
	}

	if word[idx] == '.' {   // 出现的每个字符都是匹配的
		for i := range dic.next {
			if dic.next[i] != nil && search(dic.next[i], word, idx+1) {
				return true
			}

		}
		return false
	} else {
		index := word[idx] - 'a'
		return dic.next[index] != nil && search(dic.next[index], word, idx+1)
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
