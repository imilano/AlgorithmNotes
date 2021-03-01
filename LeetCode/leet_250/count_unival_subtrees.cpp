
// 采用后序遍历的方式

class Solution {
    public:
        int countUnivalSubtrees(TreeNode* root) {
            int res = 0;
            isUnival(root,-1,res);
            return res;
        }

        bool isUnival(TreeNode* root, int val, int& res) {
            if (!root) return true;
            if (!isUnival(root->left, root->val, res) | !isUnival(root->right, root->val, res))   // 如果当前节点值与其左节点值不相等并且与右节点曳步相等，那么返回false
                return false;

            res++;   // 当前节点值与左节点值相等或者与右节点值相等或者与二者都相等
            return root.val == val;  // 当前根节点值是否与其父节点值相等。
        }
}