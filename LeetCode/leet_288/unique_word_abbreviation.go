package leet_288

import (
	"sort"
	"strconv"
)

/*
An abbreviation of a word follows the form <first letter><number><last letter>. Below are some examples of word abbreviations:

a) it                      --> it    (no abbreviation)

     1
b) d|o|g                   --> d1g

              1    1  1
     1---5----0----5--8
c) i|nternationalizatio|n  --> i18n

              1
     1---5----0
d) l|ocalizatio|n          --> l10n
Assume you have a dictionary and given a word, find whether its abbreviation is unique in the dictionary. A word's abbreviation is unique if no other word from the dictionary has the same abbreviation.

Example:

Given dictionary = [ "deer", "door", "cake", "card" ]

isUnique("dear") -> false
isUnique("cart") -> true
isUnique("cane") -> false
isUnique("make") -> true
*/

type wordAbbr struct {
	m map[string]int
}

func (w *wordAbbr)validWordAbbr(dic []string) {
	sort.Strings(dic)
	for i := range dic {
		if i > 0 && dic[i] == dic[i-1] {
			continue   // 跳过重复字符
		}
		d := buildAbbr(dic[i])
		w.m[d]++
	}
}

func (w *wordAbbr) isUnique(word string)bool {
	d := buildAbbr(word)
	if v,ok := w.m[d]; ok {
		if v > 1 {
			return false
		} else if v== 1{
			return true
		}
	}

	// don't exist in dictionary
	return false
}

func buildAbbr(s string) string {
	if len(s) <= 2 {
		return s
	}

	n := len(s)
	return s[:1] + strconv.Itoa(len(s[1:n-1]))+s[n-1:n]
}
