/*  题目描述：
 *    输入一颗二叉树的根节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
 *    路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。(注意: 在返回值的list中，数组长度大的数组靠前)
 * 
 *  题解：
 *      由于路径是从根节点出发到叶节点，也就说路径总是以根节点作为起始点，因此 我们首先需要遍历根节点。在树的三种遍历序列中，只有前序序列是首先访问根节点的。
 *      用前序遍历的方式访问到某一节点时，我们把该节点添加到路径上，并累加该节点的值。
 *          - 如果该节点是叶节点，并且路径中节点值的和刚好等于输入的整数，则当前路径符合要求，将他打印出来。
 *          - 如果该节点不是叶节点，则继续访问子的子节点。
 *          - 如果当前节点不是叶节点，则继续访问它的子节点。
 *          - 当前节点访问结束后，递归函数将自动回到它的父节点。因此，我们在函数退出之前要在路径上删除当前节点并减去当前节点的值，以确保返回父节点时路径刚好是从根节点到父节点
 */


#include <vector>

using std::vector;

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
   vector<int> path;
   vector<vector<int> > allPath;
   
   vector<vector<int> > FindPath(TreeNode* root, int expectNumber) {
      findEachPath(root, expectNumber);
      return allPath;
 }

    void findEachPath(TreeNode* root,int expectNumber) {
      if (!root) return;
      path.push_back(root->val);
      if (root->left == nullptr && root->right==nullptr && root->val == expectNumber)
        allPath.push_back(path);

      if (!root->left) findEachPath(root->left, expectNumber-root->val);
      if (!root->right) findEachPath(root->right, expectNumber-root->val);
      path.pop_back();  // 每次遍历完一个节点都要删除当前节点
    }
};
