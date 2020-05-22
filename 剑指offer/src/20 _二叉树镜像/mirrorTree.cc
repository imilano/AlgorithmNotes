// 题目描述
//     操作给定的二叉树，将其变换为源二叉树的镜像。



#include <stack>

using namespace std;

struct TreeNode
{
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) :
			val(x), left(nullptr), right(nullptr) {
	}
};
class Solution {
public:

// 递归实现，注意交换顺序
    void Mirror(TreeNode *pRoot) {
	    if (pRoot == nullptr)
		    return;
	if (pRoot->left == nullptr && pRoot->right == nullptr)
		return;
	
	TreeNode *temp = pRoot->left;
	pRoot->left = pRoot->right;
	pRoot->right = temp;

	if (pRoot->left != nullptr)
		Mirror(pRoot->left);
	if (pRoot->right != nullptr)
		Mirror(pRoot->right);
    }

    // 非递归实现，使用栈实现
    void Mirror(TreeNode *pRoot) {

        if(pRoot==nullptr)
            return;
        stack<TreeNode*> stackNode;
        stackNode.push(pRoot);
        while(stackNode.size()){
            TreeNode* tree=stackNode.top();
            stackNode.pop();
            if(tree->left!=nullptr || tree->right!=nullptr){
                TreeNode *ptemp=tree->left;
                tree->left=tree->right;
                tree->right=ptemp;
            }
            if(tree->left)
                stackNode.push(tree->left);
            if(tree->right)
                stackNode.push(tree->right);
        }
    }
};