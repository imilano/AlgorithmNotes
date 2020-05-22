#include <queue>
#include <iostream>

struct TreeNode
{
        int val;
        TreeNode *left;
        TreeNode *right;
        TreeNode(int x) : left(nullptr),right(nullptr) {}
 };


 void BFS(TreeNode* root) {
   if (!root) return;
   std::queue<TreeNode *> q;
   TreeNode *node;
   q.push(root);
   while (!q.empty())
   {
     node = q.front();
     q.pop();
     std::cout << node->val;

     if (!node->left) q.push(node->left);
     if (!node->right) q.push(node->right);
   }
 }
