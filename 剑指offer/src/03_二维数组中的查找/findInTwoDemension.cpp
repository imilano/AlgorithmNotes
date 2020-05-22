// 题目描述
// 在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

// 题解：每次都从数组右上角元素m开始查找。如果target大于m，由于数组自左向右递增，自上向下递增，所以该行不可能有符合条件的元素，删除该行；同理如果target小于m，删除该列；重复上述步骤，直到查找成功即可。


#include <vector>


class Solution {
public:
bool Find(int target, std::vector<std::vector<int> > array) {
    if (array.empty()) return false;
    if (target < array[0][0]) return false;

    bool found =false;
    int row=0;
    int col = array[0].size()-1;

    while(row < array.size() && col >=0) {
        if (array[row][col] == target) {
            found = true;
            break;
        }
        else if (array[row][col] > target) col--;
        else 
            row++;
    }
     
    return found;


//     if (array.empty())return false;
//     //if (target < array[0][0])return false;
//     int _length = array.size();
//     for (int i = 0; i < _length; i++)
//     {
//         if (array[i].empty())continue;
//         else if(target >= array[i][0])
//         {
//             if (target <= array[i][array[i].size() - 1])
//             {
//                 for (int j = array[i].size() - 1; j >= 0; j--)
//                 {
//                     if (target == array[i][j])return 1;
//                     else if (target > array[i][j])break;
//                 }
//             }
//             else {
//                 continue;
//             }
//         }
//         else return false;
//     }
//     return false;
  }
};