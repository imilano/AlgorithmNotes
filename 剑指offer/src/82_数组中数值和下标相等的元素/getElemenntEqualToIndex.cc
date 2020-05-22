/*  找出数组中数值和下标相等的元素。假设一个单调递增的数组中的每个元素之都是整数并且是唯一的。请编程实现一个函数，找出数组中任意一个数值等于其下标的元素。
 *
 *   题解：使用二分查找法。每次比较mid，如果mid=arr[mid],那么返回mid即可；如果mid < arr[mid],由于数组是递增排序，那么任何下标大于mid的下标k，
 *   其arr[k]也大于k，因此我们只需要在mid左边查找即可；同理，当mid > arr[mid]时，只需要在mid右边查找即可。
 */

int GetNumberSameAsIndex(const int* numbers,int length) {
  if (numbers == nullptr || length <= 0) return -1;

  int left = 0;
  int right = length - 1;
  while(left <= right) {
    int middle = left + ((right - left) >> 1);
    if (numbers[middle] == middle) return middle;

    if (numbers[middle] > middle) right = middle - 1;
    else
      left = middle + 1;
  }
  return -1;
}