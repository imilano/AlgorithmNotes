package leet_284

/*
Design an iterator that supports the peek operation on a list in addition to the hasNext and the next operations.

Implement the PeekingIterator class:

PeekingIterator(int[] nums) Initializes the object with the given integer array nums.
int next() Returns the next element in the array and moves the pointer to the next element.
bool hasNext() Returns true if there are still elements in the array.
int peek() Returns the next element in the array without moving the pointer.


Example 1:

Input
["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 2, 2, 3, false]

Explanation
PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [1,2,3]
peekingIterator.next();    // return 1, the pointer moves to the next element [1,2,3].
peekingIterator.peek();    // return 2, the pointer does not move [1,2,3].
peekingIterator.next();    // return 2, the pointer moves to the next element [1,2,3]
peekingIterator.next();    // return 3, the pointer moves to the next element [1,2,3]
peekingIterator.hasNext(); // return False
*/

// TODO
