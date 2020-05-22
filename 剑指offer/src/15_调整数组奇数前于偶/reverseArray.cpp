// 题目描述
//         输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

// 题解： 
//         1. 非稳定算法：双指针
//         2. 稳定算法：插入排序，插排是稳定排序



#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
// 不稳定
    void reOrderArray(vector<int> &array) {
            if (array.empty() || array.size() == 1)
                    return;
            int left = 0, right = array.size() - 1;
            while(left<right) {
                    while (left < right && isOdd(array[left]))
                            left++;
                    while (left < right && !isOdd(array[right]))
                            right--;
                        if (left < right)
                                swap(array, left, right);
            }
    }
        // 稳定
    void reOrderArray(vector<int> &array) {
            if (array.empty() || array.size() == 1)
                    return;
            int len = array.size();
            int k = 0;  // k个奇数
            for (int i = 0; i < len;i++) {  // 插入排序的思想
                    if (isOdd(array[i]))
                    {
                            int j = i;
                            while (j > k)
                            {
                                    int tmp = array[j];
                                    array[j] = array[j - 1];
                                    array[j - 1] = tmp;
                                    j--;
                            }
                            k++;
                        }
                }
        }

    bool isOdd(int n) {
        if(n &1) 
                return true;
        else
                return false;
        }
    void swap(vector<int> &arr,int left, int right) {
            int temp = arr[left];
            arr[left] = arr[right];
            arr[right] = temp;
    }
};
 

 int main() {
         vector<int> arr = vector<int>{1, 2, 3, 4, 5, 6, 7};
         Solution s;
         s.reOrderArray(arr);
         return 0;
 }