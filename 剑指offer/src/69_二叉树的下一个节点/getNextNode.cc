// // 题目描述：给定一颗二叉树和其中的一个节点，如何找出中序遍历序列的下一个节点？树中的节点除了有两个分别指向左右节点的指针外，
//            还有一个指向父节点的指针
// 题解：
//     分情况讨论： 1. 如果一个节点有右子树，那么它的下一个节点就是右子树的最左节点，也就是说，从右节点出发一直沿着指向左字节点的指针，
//                   即可得它的下一个节点。
//                2. 如果一个节点没有右子树：
//                     2.1 如果节点是它父节点的左字节点，那么它的下一个节点就是它的父节点。
//                     2.2 如果节点是它父节点的右字节点，那么我们可以沿着指向父节点的指针一直向上遍历（意味着有一个指向父节点的指针），
//                          直到找到一个是他父节点的左字节点的节点。如果这样的节点存在，那么这个节点的父节点就是我们要找的下一个节点。

struct TreeNode
{
        int val;
        TreeNode *left;
        TreeNode *right;
        TreeNode *parent;
        TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 };

TreeNode* GetNext(TreeNode* root) {
    if (root == nullptr) return nullptr;
    TreeNode* next = nullptr;
    if (root->right != nullptr) { // 如果右子树非空，那么就往它的右子树的左边找
        TreeNode* pRight = root->right;
        while(pRight->left != nullptr) {
            pRight = pRight->left;
        }
        next = pRight;
    }else if (root->parent !=  nullptr) { // 当前节点右子树为空
        TreeNode* current = root;
        TreeNode* pParent = root->parent;
        while(pParent != nullptr && current == pParent->right){ //如果当前节点是它的父节点的右子节点，那么向上找到是它父节点的左字节点的节点
            current = pParent;
            pParent = pParent->parent;
        }
        next = pParent; // 如果节点是它父节点的左字节点，那么他的父节点就是下一个节点
    }
    return next;
}