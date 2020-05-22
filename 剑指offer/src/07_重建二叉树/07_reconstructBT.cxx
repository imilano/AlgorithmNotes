// 题目描述
// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。

// 题解：前序遍历的第一个元素就是根结点，由于不含重复数字，扫描中序序列，就能确定根节点的值的位置，根节点前的部分就是左子树，根节点后的部分就是右子树。重复上述过程，就是一个递归的过程。


#include <iostream>
#include <vector>

using namespace std;

//Definition for binary tree
struct TreeNode
{
        int val;
        TreeNode *left;
        TreeNode *right;
        TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 };

class Solution {
public:
        TreeNode *reConstructBinaryTree(vector<int> pre, vector<int> vin){
                TreeNode *root = reConstructBT(pre, 0, pre.size() - 1, vin, 0, vin.size() - 1);
                return root;
        };

private :
        TreeNode *reConstructBT(vector<int> pre, int startPre, int endPre, vector<int> in, int startIn, int endIn){
                if (startPre > endPre || startIn > endIn)
                        return nullptr;
                TreeNode *root = new TreeNode(pre[startPre]);

                for (int i = startIn; i < endIn; i++) {
                        if (in[i] == pre[startPre]) { // 找到中序序列中根节点的下标
                                root->left = reConstructBT(pre, startPre + 1, i - startIn + startPre, in, startIn, i - 1); // i-startIn表示中序序列中根节点前左子树有几个节点 ，再加上startIn表示的是前序序列中根节点的左子树最后一个元素的位置
                                root->right = reConstructBT(pre, i - startIn + startPre + 1, endPre, in, i + 1, endIn);
                        }
                }
                return root;
        };
};

