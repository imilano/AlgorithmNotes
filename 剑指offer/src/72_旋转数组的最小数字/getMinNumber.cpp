// 题目描述：
//         把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
//         例如，数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1.

// 题解：
//         采用二分查找思想。使用三个指针，一个start指向开头，一个end指向结尾，一个mid指向中间。根据题目中旋转的规则，第一个元素应该大于等于最后一个元素。
//         找到中间的元素mid，如果中间元素位于前面的递增子数组,那么最小数字应该在中间元素的后面，于是令start=mid
//                         同理，如果中间元素位于后面的递增子数组，那么最小数字应该在中间元素的前面，于是end=mid;
//                         最终，第一个指针将指向前面子数组的最后一个元素，而第二个指针将指向后面子数组的最后一个元素，也即二者最终会指向两个相邻的元素，而第二个指针指向的刚好是最小的元素。循环至此结束。
//         特殊情况：
//                 1. 数组本身就是有序的，那么最小元素就是第一个元素，于是开始时我们令mid=start。
//                 2. 假设数组中有重复元素，导致mid、start和end都指向相等的元素，这样我们就无法确定mid指向的元素是属于第一个递增数组还是第二个递增数组，例如{1 0 1 1 1}，此时认为最小的数字在mid后面是错误的。在这种情况下，就要使用顺序查找。

#include <vector>

class Solution {
public:
    int minNumberInRotateArray(std::vector<int> rotateArray) {
        if (rotateArray.empty()) return -1;
        int len = rotateArray.size();

        int start = 0;
        int end = len-1;

        int mid = start; // 为了预防数组有序的情况

        while(rotateArray[start] >= rotateArray[end]) {
            if (end-start == 1) {
                mid = end;
                break;
            }

            mid = (start+end)/2;

            // 如果三者相等，则在start和end之间进行顺序查找
            if(rotateArray[start] == rotateArray[end] && rotateArray[end] == rotateArray[mid]) {
                return midInOrder(rotateArray,start,end);
            }

            if (rotateArray[mid] >= rotateArray[start]) start = mid;
            else if(rotateArray[mid] <= rotateArray[end]) end = mid;
        }
        return rotateArray[mid];
    }

    int midInOrder(std::vector<int> arr, int s, int e) {
        int min = arr[s];
        for (int i =s; i <= e; i++) {
            if (min > arr[i]) min = arr[i];
        }
        return min;
    }
};
