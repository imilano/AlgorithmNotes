package leet_267

/*
Given a string s, return all the palindromic permutations (without duplicates) of it. Return an empty list if no palindromic permutation could be form.

Example 1:

Input: "aabb"
Output: ["abba", "baab"]
Example 2:

Input: "abc"
Output: []
Hint:

If a palindromic permutation exists, we just need to generate the first half of the string.
To generate all distinct permutations of a (half of) string, use a similar approach from: Permutations II or Next Permutation.
*/

//根据上一题，如果一个字符串的排列可以构成回文串，那么该字符串中的字符要么出现偶数次，要么出现奇数次的字符只有一个。那么如何将字符构成字符串呢？
//根据提示，如果所有字符都只出现了偶数次，那么我们只需要字符的前半部分，然后和该字符串的逆序进行拼接即为我们所需的回文串；而如果有字符出现了奇数次，
//那么就需要将字符分为三部分，一个是前半部分，一个中间部分（包含出现奇数次的字符）以及一个后半部分（前半部分的逆序）。我们只需要对前半部分进行
//回文串排列，然后再拼接上中间部分以及前半部分的逆序即可。
//
//那么首先使用map统计每个字符出现的次数。然后对这些字符进行遍历，如果该字符出现了偶数次，那么添其一半长度到front；如果该字符出现了奇数次，那么
//也将该字符的一半添加到front，同时将该字符添加到mid（只添加一次）；如果奇数字符出现次数大于1，直接返回空。现在我们已经有了前半部分字符和中间
//字符，那么我么只需要对前半部分字符进行回文排列，排列结束后，将排列出来的字符拼接上中间字符以及排列字符的逆序，就可以得到最终的回文串，将其添加
//到结果中即可。在进行回文排列的时候，需要注意这几个问题：首先，front中的偶数字符会有多个，如果我们使用交换的方式进行回文排列的话，可以直接跳过
//这部分字符的排列而不影响排列结果；其次回文串可能会再结果字符中多次出现，所以需要使用set进行去重
func generatePalindromes(s string) []string {
	var res []string
	m := make(map[uint8]int)
	set := make(map[string]bool)
	for i := range s {
		m[s[i]]++
	}

	f := func(s string, c int) string {
		var res string
		for i := 0; i < c; i++ {
			res += s
		}

		return res
	}

	var front, mid string
	for k, v := range m {
		if v%2 == 1 {
			mid += string(k)
		}
		front += f(string(k), v/2)
		if len(mid) > 1 {
			return nil
		}
	}

	permute(0, front, mid, &set)
	for k, v := range set {
		if v == true {
			res = append(res, k)
		}
	}

	return res
}

func reverseString(s string) string {
	rs := []rune(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		rs[i], rs[n-i-1] = rs[n-i-1], rs[i]
	}

	return string(rs)
}
func swapString(s string, a, b int) string {
	rs := []rune(s)
	rs[a], rs[b] = rs[b], rs[a]
	return string(rs)
}
func permute(start int, front string, mid string, res *map[string]bool) {
	if start >= len(front) {
		s := front + mid + reverseString(front)
		(*res)[s] = true
	}

	for i := start; i < len(front); i++ {
		if start != i && front[start] == front[i] { // 跳过重复字符
			continue
		}
		front = swapString(front,start,i)
		permute(start+1, front, mid, res)
		front = swapString(front,start,i)
	}
}


//-----------------------
// 同样的思路，只是不是采用交换字符的方法来生成全排列，而是通过添加字符的方式来生成全排列
func generatePalindromes2(s string) []string {
	var res []string
	m := make(map[uint8]int)
	set := make(map[string]bool)

	for i := range s {
		m[s[i]]++
	}

	var front,mid string
	f := func(s string, c int) string {
		var res string
		for i := 0; i < c; i++ {
			res += s
		}

		return res
	}
	for k,v := range m {
		if v %2  == 1 {
			mid += string(k)
		}
		m[k] = v/2
		front += f(string(k),v/2)
		if len(mid) > 1 {
			return nil
		}
	}

	permute2(front,mid,"",m,&set)
	for k,v := range set {
		if v == true {
			res = append(res, k)
		}
	}

	return res
}

func permute2(front,mid string, out string,m map[uint8]int,set *map[string]bool) {
	if len(out) >= len(front) {
		s := out + mid + reverseString(out)
		(*set)[s]=true
	}

	for k,_ := range m {
		if m[k] > 0 {
			m[k]--
			permute2(front, mid, out+string(k), m, set)
			m[k]++
		}
	}
}
