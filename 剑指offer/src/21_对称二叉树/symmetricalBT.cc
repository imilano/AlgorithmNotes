// 题目描述
//         请实现一个函数，用来判断一颗二叉树是不是对称的。注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的。


// 
//  必须考虑空节点，空节点不能漏，使用特殊值标记
// 
//      前序遍历是跟左右，我们创造出一种对称遍历算法根右左，在遍历的同时把空节点也考虑进去，最后比较两个序列，若两个序列下标对应的值相等，则为对称二叉树，反之则不是。
//  
// 

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(nullptr), right(nullptr) {
    }
};

class Solution {
public:
    bool isSymmetrical(TreeNode* pRoot)
    {
        return isSymmetricalTree(pRoot, pRoot);
    }

    bool isSymmetricalTree(TreeNode* root1,TreeNode*  root2) {
            if (root1 == nullptr && root2 == nullptr)
                    return true;
            if (root1==nullptr || root2 == nullptr)
                    return false;
            if (root1->val != root2->val)
                    return false;
            return isSymmetricalTree(root1->left, root2->right) && isSymmetricalTree(root1->right, root2->left);
    }
};


