/*	题目描述
 *		从上往下打印出二叉树的每个节点，同层节点从左至右打印。
 * 
 * 题解：
 * 		二叉树的层次遍历，需要借助一个队列来实现。
 * 
 *
 */
#include <vector>
#include <queue>

using namespace std;

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
    vector<int> PrintFromTopToBottom(TreeNode* root) {
	    vector<int> result;
	    queue<TreeNode*> q;
	    TreeNode *tmp;
	    if (!root)
		    return result;
	    q.push(root);
	    while(!q.empty()) {
		    tmp = q.front();
		    q.pop();
		    result.push_back(tmp->val);

		    if (tmp->left)
			q.push(tmp->left);
		    if (tmp->right)
			q.push(tmp->right);
	    }
	    return result;
    }
};