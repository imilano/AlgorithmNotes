/*  题目描述
 *  请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推。
 *
 * 题解： 经过分析可知，如果当前节点在奇数层，则将当前节点的子节点从左到右入栈，如果当前遍历节点在偶数层，则将当前节点从右到左入栈。
 *       而由于使用单个栈会造成混乱，所以使用两个栈来分别存储奇数层的子节点和偶数层的子节点
 *  
 */  



#include <vector>
#include <stack>
#include <queue>
#include <algorithm>
#include <iostream>

using std::vector;
using std::stack;
using std::cout;

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
  // 队列方法，使用reverse
    vector<vector<int> > Print(TreeNode* pRoot) {
        vector<vector<int>> res;
        if(pRoot == nullptr)
            return res;
        std::queue<TreeNode*> que;
        que.push(pRoot);
        bool even = false;
        while(!que.empty()){
            vector<int> vec;
            const int size = que.size();
            for(int i=0; i<size; ++i){
                TreeNode* tmp = que.front();
                que.pop();
                vec.push_back(tmp->val);
                if(tmp->left != nullptr)
                    que.push(tmp->left);
                if(tmp->right != nullptr)
                    que.push(tmp->right);
            }
            if(even)
                std::reverse(vec.begin(), vec.end());
            res.push_back(vec);
            even = !even;
        }
        return res;
    }

    vector<vector<int> > Print(TreeNode* pRoot) {
      vector<vector<int> > result;
      if (pRoot == nullptr) return result;

      stack<TreeNode*> st[2]; // 一个栈用于保存奇数层节点，另一个栈用于保存偶数层节点
      int current = 0, next = 1;
      st[current].push(pRoot);

      while(!st[current].empty() || !st[next].empty()) {
        TreeNode* node = st[current].top();
        st[current].pop();        
        vector<int> temp;
        temp.push_back(node->val);

        if (current == 0) { // 当前层从左向右遍历，那么将其子节点从左向右压栈
            if (node->left) st[next].push(node->left);
            if (node->right) st[next].push(node->right);
        } else {  // 当前层从右向左遍历，那么将其子节点从右向左压栈
          if (node->right) st[next].push(node->right);
          if (node->left) st[next].push(node->left);
        }

        if (st[current].empty()) {
          result.push_back(temp);
          current = 1 - current;
          next = 1 - next;
      }
    } 
    return result;  
  } 
};



class Solution {
    public:
     vector<int> PrintFromTopToBottom(TreeNode* root) {
        vector<int> result;
        TreeNode* tmp;

        if (!root)
            return result;
        stack<TreeNode*> level[2]; // 一个栈用于保存奇数层的子节点，一个栈用于保存偶数层的字节点

        int current = 0, next = 1;
        level[current].push(root);

        while(!level[0].empty() && !level[1].empty()) {
          tmp = level[current].top();
          result.push_back(tmp->val);
          level[current].pop();

          if (current == 0) {  // 当前层从左到右，则子节点从左向右压栈
            if (tmp->left) level[current].push(tmp->left);
            if (tmp->right) level[current].push(tmp->right);
          } else {  // 当前层从右到左，则字节点从右向左压栈
            if (tmp->right) level[current].push(tmp->right);
            if (tmp->left) level[current].push(tmp->left);
          }

          if (level[current].empty()) {  // 换行并交替打印左右子栈
            cout << "\n";   // 注意换行
            current = 1 - current;
            next = 1 - next;
          }
        }
     }
};
