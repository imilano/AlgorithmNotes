#include <iostream>
#include <stack>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
    TreeNode(int x) :
            val(x), left(nullptr), right(nullptr) {
    }
};

//  递归版本
void InOrderRec(TreeNode* root) {
  if (root == nullptr) return;
  InOrderRec(root->left);
  std::cout << root->val;
  InOrderRec(root->right);
}

// 非递归版本，使用栈
void InOrderNonRec(TreeNode* root) {
  if (root == nullptr) return;

  std::stack<TreeNode*> nstack;
  TreeNode* node = root;

  while(node != nullptr || !nstack.empty()) {
          // 一直向左边走
          while(node != nullptr) {
            nstack.push(node);
            node = node->left;
          }

          // 此时栈顶与那素应该是当前最左元素
          if(!nstack.empty()) {
            node = nstack.top(); //输出该元素
            std::cout << node->val;
            nstack.pop();
            node = node->right;  // 访问右子树
          }
  }
}