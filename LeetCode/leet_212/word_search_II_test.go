package leet_212

import (
	"fmt"
	"testing"
)

func TestFindWords(t *testing.T) {
	words :=  []string{"aa"}
	board := [][]byte{
		{'a','a'},
	}

	fmt.Println(findWords(board,words))

	board2 := [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words2 := []string{"oath","pea","eat","rain"}
	fmt.Println(findWords(board2,words2))
}
