#include <stack>
#include <iostream>

struct TreeNode
{
        int val;
        TreeNode *left;
        TreeNode *right;
        TreeNode(int x) : left(nullptr),right(nullptr) {}
 };

 void DFS(TreeNode* root) {
   if (root == nullptr) return;
   std::stack<TreeNode*> s;
   s.push(root);
   TreeNode* node;

   while ( !s.empty()) {
     node = s.top();
     s.pop();
     std::cout << node->val;

     if (!node->right) s.push(node->right); // 先把右节点压入栈中
     if (!node->left) s.push(node->left); // 再把左节点压入栈中
   }
 }