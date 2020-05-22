// 题目描述：给定一棵二叉搜索树，请找出其中的第k小的结点。例如， （5，3，7，2，4，6，8)中，按结点数值大小顺序第三小结点的值为4。
// 题解：一棵BST的中序遍历序列就是一个有序数组。其中第k个就是第K大节点。实际上就是考察中序遍历。

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
 int count = 0;
 TreeNode* KthNode(TreeNode* pRoot, int k) {
   if (pRoot == nullptr || k < 0) return nullptr;

   return getKthNode_v2(pRoot, k);
 }

    TreeNode* getKthNode_v2(TreeNode* root, int k) {
      if (root != nullptr) {
        TreeNode* node = getKthNode_v2(root->left, k);
        if (node != nullptr) return node;
        if (++count == k) return root;
        node = getKthNode_v2(root->right, k);
        if (node) return node;
      }
      return nullptr;
    }

    TreeNode* getKthNode_v1(TreeNode* root, int k) { 
        TreeNode* target = nullptr;
        if (root->left != nullptr) target = getKthNode_v1(root->left, k);

        if (target == nullptr) {
          if (k == 1) target = root;
          k--;
        }

        if (target == nullptr && root->right != nullptr)
          target = getKthNode_v1(root->right, k);
        return target;
    }
};