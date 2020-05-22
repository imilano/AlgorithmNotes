// 题目描述：一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。

// 题解：使用抑或运算。一个数字跟自己本身异或，其值为0.于是我们把数组中的所有数字都进行异或，最后的结果就是两个只出现一次的数字相异或的结果。在该结果中找到第一位值为1的数位，将数组中其他数字以该位是否为1为标准 ，将原数字分为两组。
// 因为两个相同的数字各位都相同，所以 相同的数字都分到同一组，而1是两个不同数字异或的结果，于是，两个不同的数字 肯定会被分到不同组。再把这两组的数字分别异或，就能分别找出这两个数字。

#include<vector>

class Solution {
public:
// 从右向左找到第一个1所在的位置（自左向右和自右向左的结果都一样）
 int findFirst1(unsigned int data) { 
  int num = 0;
  while((data &1)== 0 && num < 8*sizeof(data)) {
    data = data >> 1;
    num++;
  }
  return num;
 }

 // 判断data的index位是否为1
 bool isBit1(int data, unsigned int index) {
   data = data >> index;
   return data & 1;
 }
 void FindNumsAppearOnce(std::vector<int> data, int* num1, int* num2) {
   if (data.empty()) return;
   int len = data.size();

   unsigned int result = 0; // 0和任何一个数亦或的结果还是原来的数
   *num1 = 0, *num2 = 0;
   for (int i = 0; i < len; i++) result = result ^ data[i];

   int indexOfFirst1 = findFirst1(result);

   
   for (int i = 0; i < len; i++) {
     if (isBit1(data[i], indexOfFirst1))
       *num1 ^= data[i];
     else
       *num2 ^= data[i];
   }
    }
};