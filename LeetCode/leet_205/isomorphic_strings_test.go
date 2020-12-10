package leet_205

import (
	"fmt"
	"testing"
)

func TestIsomorphic(t *testing.T) {
	arr := []struct {
		s string
		t string
	}{
		{"egg","add"},
		{"foo","bar"},
		{"paper","title"},
	}

	for _,v := range arr {
		fmt.Println(v.s," ",v.t," ",isIsomorphic(v.s,v.t))
	}
}
