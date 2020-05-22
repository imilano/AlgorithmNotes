/* 题目描述： 输入一颗二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1,那么它就是一颗平衡二叉树。
 *
 *  题解：采用后序遍历的方式，这样，当便利到一个节点之前已经遍历了它的左子树和右子树，知道了它的高度。只需要在遍历每个节点的时候记录它的深度（某一节点的深度等于它到页节点的路径的长度），
 *        我们就可以一边遍历一边判断每个节点是不是平衡的。
 */ 


#include <algorithm>

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
    
     int getHeight(TreeNode* root) {
       if (root == nullptr) return 0;
       else {
         return std::max(getHeight(root->left), getHeight(root->right)) + 1;
       }
     }

     //方法一，直接求每个节点左右子树的高度，然后判断该节点是否平衡；接下来递归地判断该节点的左右节点
     // 但是，这个方法自顶向下遍历，会导致很多节点重复遍历，效率较低
     bool isBalancedTree_v1(TreeNode* root) {
       if (root == nullptr) return 0;
       int leftHeight = getHeight(root->left);
       int rightHeight = getHeight(root->right);
       if (std::abs(leftHeight - rightHeight) > 1) return false;
       else
         return isBalancedTree_v1(root->left) && isBalancedTree_v1(root->right);
     }

      // 方法二，将方法一的自顶向下改为自底向上
      bool isBalancedTree_v2(TreeNode* root) {
        if (checkHeight(root) == -1) return false;
        else
          return true;
      }

      int checkHeight(TreeNode* root)  {
        // 空树的高度
        if (root == nullptr) return 0;
        //检查左子树是否平衡
        int leftHeight = checkHeight(root->left);
        if (leftHeight == -1) return -1;

        // 检查右子树是否平衡
        int rightHeight = checkHeight(root->right);
        if (rightHeight == -1) return -1;

        // 检查当前节点是否平衡
        int diffHeight = std::abs(leftHeight - rightHeight);
        if (diffHeight > 1) return -1;
        else
          return std::max(leftHeight, rightHeight) + 1;
      }

      // 方法三，同样是自底向上，c采用后序遍历的方式遍历每个节点，这样在遍历到一个节点之前我们已经遍历到了他的左右子树。只需要在遍历每个节点的时候记录他的深度，那么我们就可以一边遍历一边判断每个节点是不是平衡的。
      bool  isBalanced_v3(TreeNode* root, int& height) {
              if (root == nullptr) {
                height = 0;
                return true;
              }

              int leftHeight = 0, rightHeight = 0;
              bool leftRec = isBalanced_v3(root->left, leftHeight);
              bool rightRec = isBalanced_v3(root->right, rightHeight);
              height = std::max(leftHeight, rightHeight) + 1;
              if (std::abs(leftHeight - rightHeight) > 1) return false;
              return leftRec && rightRec;
      }
    bool IsBalanced_Solution(TreeNode* pRoot) {
      int height = 0;
      return isBalanced_v3(pRoot, height);
    }
};