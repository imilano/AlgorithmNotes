/* 题目描述
 * 从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行。
 *     1
 *     2 3
 *     4 5 6
 *     7 8 9 10 
 * 题解：
 *    和层次遍历类似。不同的是，我们需要一个变量来表示当前层中还没有打印的节点数；还需要另一个变量表示下一层节点的数目。
 *                          
 */

#include <vector>
#include <queue>
#include <iostream>

using std::vector;
using std::queue;
using std::cout;
using std::endl;

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
        vector<vector<int> > Print(TreeNode* pRoot) {
            vector<vector<int> > result;
            if (pRoot ==  nullptr) return result;

            queue<TreeNode*> q;
            q.push(pRoot);

            while(!q.empty()) {
              vector<int> temp;
              int low=0, high = q.size()-1;

              while( low++ <= high) {
                TreeNode* node = q.front();
                temp.push_back(node->val);

                q.pop();
                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
              }
              result.push_back(temp);
            } 
            return result;
        }

     vector<int> PrintFromTopToBottom(TreeNode* root) {
        vector<int> result;
        queue<TreeNode*> q;
        TreeNode* tmp;
        int nextLevel=0,toBePrinted;  // nextLevel 表示下一层需要打印的节点数，toBePrinted表示本层尚未打印的节点数，每一个toBePrinted都是下一层节点在队列中的个数

        if (!root)
            return result;
        q.push(root);
        toBePrinted = 1;// 第一次需要打印的节点数量

        while(!q.empty()) {
          tmp = q.front();
          q.pop();
          cout << tmp->val << " \t" << endl;
          --toBePrinted;

          if (tmp->left) {
            q.push(tmp->left);
            nextLevel++;
          }

          if (tmp->right) {
            q.push(tmp->right);
            nextLevel++;
          }

          if (toBePrinted == 0) {
            cout << "\n";   // 注意换行符
            toBePrinted = nextLevel;
            nextLevel = 0;
          }
        }
        
     }
};

