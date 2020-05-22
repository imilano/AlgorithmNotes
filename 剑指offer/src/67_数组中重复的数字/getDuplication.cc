// 题目描述
// 在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任意一个重复的数字。 例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是第一个重复的数字2。

//题解：
//    数字范围都在0~n-1，如果无重复，那么排序后数字i出现在下标为i的位置，由于数组中有重复的数字，有些位置可能存在多个数字，有些位置可能没有数字。
//    现在重排数字，从头到尾扫描每个数字.当扫描到下标为i的数字时，首先比较该数字（假设为m）是不是等于i，如果是，接着扫描下一个;如果不是，拿它和第m个数字比较，如果二者相等，则出现了第一个重复的数字;如果不等则把第i个数字和第m个数字进行交换，把m放回他的位置。接下来重复这个比较过程即可。
//

class Solution {
public:
    // Parameters:
    //        numbers:     an array of integers
    //        length:      the length of array numbers
    //        duplication: (Output) the duplicated number in the array number
    // Return value:       true if the input is valid, and there are some duplications in the array number
    //                     otherwise false
    bool duplicate(int numbers[], int length, int* duplication) {
        if (numbers == nullptr || length <= 0) return false;

        // 数字必须位于[0,n-1]
        for (int i=0; i< length;i++) 
        if (numbers[i] <0 || numbers[i] > length-1) return false;

        for(int i=0; i<length; i++) {
            while(numbers[i] !=i) {
                if (numbers[i] == numbers[numbers[i]]) {
                    *duplication = numbers[i];
                    return true;
                }
                int temp = numbers[i];
                numbers[i] = numbers[temp];
                numbers[temp] = temp;
            }
        }
        return false;
    }
};