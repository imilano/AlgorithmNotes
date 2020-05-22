

/*  题目描述
 *    输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
 *  题解：
 *      判断一个数组是否是二叉树的后序搜索树遍历结果。二叉搜索树左边都比根节点小，右边都比根节点大，而后序遍历的最后一个节点总是根节点，
 *      它的左子树都比根节点小，右子树都比根节点大，它的左右子树也有同样的规律，也就是说，这是一个递归的过程
 */ 


#include <vector>

using  std::vector;

class Solution {
public:

// 循环解法
    bool VerifySquenceOfBST(vector<int> sequence) {
      if (sequence.empty()) return false;
      int len = sequence.size();

      return isSequenceOfBST_rec(sequence, 0,len-1);
    }

  bool isSequenceOfBST_rec(vector<int> s,int left, int right ) {
        if (left >= right) return true;
        int mid = right; // mid 为左右子树的分界线，mid为右子树的第一个节点
        while(mid > left && s[mid-1] > s[right]) mid--;
        for (int j = mid-1; j >=left; j--) // 如果左子树中有比跟节点大的，返回false
          if (s[j] > s[right]) return false;

        return isSequenceOfBST_rec(s,left,mid-1) && isSequenceOfBST_rec(s,mid,right-1); // 分别递归左子树序列和右子树序列。注意右子树是right-1
        //   if (left == right) return true;  // 递归结束条件

        //   int root = s[right];  //  后序遍历，最后一个节点是根节点

        //   // 后序遍历，左子树都比根节点小
        //   int mid = left;
        //   for (; mid < right; ++mid) {
        //     if (s[mid] > root) break;  // 找到左子树的范围
        // }

        // int j = mid;
        // for (; j < right;++j) {  // 右子树有比root小的，则非BST
        //         if (s[j] < root) return false;
        // }

        // if (mid  == left || mid == right) return isSequenceOfBST(s, left,right-1);  
        // else
        //   return isSequenceOfBST(s, left, mid - 1) &&  // 递归遍历左子树
        //          isSequenceOfBST(s, mid, right - 1);  //  递归遍历右子树
    }

        // 递归解法。 左子树的节点都比size小，右子树的节点都比size大。i从做向右找到第一个大于根节点的位置，之后开始向右遍历右子树，一次遍历完之后，如果是搜索树，i必然等于size;如果小雨size，则不是二叉搜索树
bool isSquenceOfBST_iterative(vector<int> sequence) {
        int size = sequence.size();
        if(0==size)return false;
 
        int i = 0;
        while(--size) // BST每次除去最右一个节点之后，都还是一颗BST
        {
            while(sequence[i++]<sequence[size]);
            while(sequence[i++]>sequence[size]);
 
            if(i<size)return false;
            i=0;
        }
        return true;
    }
};
