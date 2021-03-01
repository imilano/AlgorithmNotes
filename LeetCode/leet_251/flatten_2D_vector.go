package leet_251

/*
Implement an iterator to flatten a 2d vector.

For example,
Given 2d vector =

[
  [1,2],
  [3],
  [4,5,6]
]


By calling next repeatedly until hasNext returns false, the order of elements returned by next should be: [1,2,3,4,5,6].

Hint:

How many variables do you need to keep track?
Two variables is all you need. Try with x and y.
Beware of empty rows. It could be the first few rows.
To write correct code, think about the invariant to maintain. What is it?
The invariant is x and y must always point to a valid point in the 2d vector. Should you maintain your invariant ahead of time or right when you need it?
Not sure? Think about how you would implement hasNext(). Which is more complex?
Common logic in two different places should be refactored into a common method.
Follow up:
As an added challenge, try to code it using only iterators in C++ or iterators in Java.
*/

// 方式1，直接使用一个坐标标志当前遍历到的位置，遍历完当前数字之后下标自增；如果最后idx大于数组长度，那么hasNext返回false
type Vector2D struct {
	v   []int
	idx int
}

func New(n [][]int) *Vector2D {
	vec := &Vector2D{
		v:   []int{},
		idx: 0,
	}

	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[i]); j++ {
			vec.v = append(vec.v, n[i][j])
		}
	}
	return vec
}

func (v *Vector2D) next() int {
	res := v.v[v.idx]
	v.idx++
	return res
}

func (v *Vector2D) hasNext() bool {
	return v.idx < len(v.v)
}

// 方式2，使用两个变量x和y；对于hasNext函数，检查当前x是否小于总行数，y是否和当前行的列数相同，如果相同，说明要转到下一行，则x自增1，y设为0，如果此时
// x还是总小于总行数，说明下一个值可以被取出来，那么在next函数就可以直接取出行为x列为y的数字，并将y自增1
type Vector2D2 struct {
	data [][]int
	x, y int
}

func New2(nums [][]int) *Vector2D2 {
	return &Vector2D2{
		data: nums,
		x:    0,
		y:    0,
	}
}

func (v *Vector2D2) next() int {
	v.hasNext()
	res := v.data[v.x][v.y]
	v.y++
	return res
}

func (v *Vector2D2) hasNext() bool {
	for v.x < len(v.data) && v.y == len(v.data[v.x]) {
		v.x++
		v.y = 0
	}

	return v.x < len(v.data)
}
