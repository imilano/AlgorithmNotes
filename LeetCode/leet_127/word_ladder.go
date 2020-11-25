package leet_127

/*
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
*/

// Using BFS
// Let beginWord, endWord represent startNode and endNode respectively. Every time we can only change a word to transform to next word,
// the transformation between two words are just like links between two node, hence we can form a graph, using startNode as start and endNode
// as endNode. So the problem now transform to find the shortest path from startNode to endNode, that is to say, breadth first search.
// All we need to do is just do as bfs, then comes the answer
func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := make(map[string]int)
	var que []string

	// Initialize set and queue
	for i := range wordList {
		set[wordList[i]] = 1
	}
	que = append(que, beginWord)
	bfsLevel := 1
	for len(que) != 0 {
		curNum := len(que)
		// iterate every node of current level
		// WARNING
		//for i := 0; i <len(que);i++{ // can not use this, cause we will delete node from que, so the node number is always changing,
		// hence not the number of current level node
		for i := 0; i<curNum;i++ {
			// get node of current level
			word := que[0]
			que = que[1:]

			// if find the endWord, then just return the the depth
			if endWord == word {
				return bfsLevel
			}

			// do not iterate this node again
			delete(set,word)
			// check if this node could transform to the node in dict
			for i := 0; i<len(word);i++ {
				for j := 0; j<26;j++ {
					tmp := []byte(word)
					tmp[i] = byte('a' + j)
					if v := set[string(tmp)]; v == 1 {
						que = append(que,string(tmp))
					}
				}
			}
		}
		// increase the depth
		bfsLevel++
	}

	return 0
}

