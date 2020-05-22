
// 题目描述：输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

// 题解： 如果一棵树只有一个节点，那么他的深度为1；
//     如果根节点只有左子树而没有右子树，那么树的深度应该是其左子树深度加1； 同理
//     ，如果跟节点只有右子树而没有左子树，那么树的深度应该是其右子树的深度加1;
//     如果一个节点既有左子树又有右子树，那么他的深度应该是左右字数的较大值加1；



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
    int TreeDepth(TreeNode* pRoot)
    {
      if (pRoot == nullptr) return 0;
      int nLeft = TreeDepth(pRoot->left);
      int nRight = TreeDepth(pRoot->right);

      return nLeft > nRight ? nLeft + 1 : nRight + 1;
    }
};