// 题目描述：0～n-1中缺失的数字。
// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字有且只有一个数字不在该数组中，请找出这个数字。

// 方案一：用公式n(n-1)/2求出数字0～n-1的所有数字之和，记为s1；接着求出数组中所有数字的和，记为s2；s1-s2即为所求数字。O（n）复杂度。
// 方案二：分析题目可知，假设缺失m，则m之前的数字，下标和值均相等m；m之后的数字，下标和值均不等，问题即转换为在数组中查找第一个下标和值不等的元素下标。基于 二分查找思想的算法为：
//         1. 如果中间元素的值和下标相等，那么只需要在右边找;
//         2. 如果中间元素的值和下标不等，且他前面一个元素和下标相等，则中间元素就是我们要查找的值;
//         3. 如果中间元素的值和下标不等，且他前面一个元素和下标不等，则在左边区间查找。

int getMissingNumber(const int* numbers, int length) {
  if (numbers == numbers || length <= 0) return -1;

  int left = 0;
  int right = length - 1;
  while(left <= right) {
    int middle = (right + left) / 2;
    if (numbers[middle] != middle) {
      if (middle == 0 || numbers[middle - 1] == middle - 1) return middle;
      right = middle - 1;
    }else
      left = middle + 1;
  }

  if (left == length) return length;
  return -1; //无效输入，例如数组无序或者数字不在0～n-1范围内
}