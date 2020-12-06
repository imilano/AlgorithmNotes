package leet_170

/*
Design and implement a TwoSum class. It should support the following operations: add and find.

add - Add the number to an internal data structure.
find - Find if there exists any pair of numbers which sum is equal to the value.

Example 1:

add(1); add(3); add(5);
find(4) -> true
find(7) -> false
*/

type twoSum struct {
	m map[int]int
}

func (t *twoSum)New() twoSum{
	return twoSum{
		m : make(map[int]int),
	}
}

func (t *twoSum)add (value int) {
	(*t).m[value]++
}

func (t *twoSum)find(value int) bool {
	for k, v := range (*t).m {
		second := value - k
		// 如果存在second或者second与k相同，并且k出现的次数大于1，则为true
		num,ok := (*t).m[second]
		if ok && (num != k || (num == k && v > 1)) {
			return true
		}
	}

	return false
}
