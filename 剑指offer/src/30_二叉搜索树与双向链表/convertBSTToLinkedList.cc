/*  题目描述:
 *      输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
 *
 *  题解:
 *         采用中序遍历的方式遍历BST,这样就能构造出一颗有序的子树.
 *          1. 将左子树钟序遍历,使其构成有序链表,并使用pre指向该链表的最大节点
 *          2. 如果左子树链表不为空,将当前root追加到pre之后,让pre指向根节点,再次构成有序链表
 *          3. 将右子树构成双向链表,并将其连接到pre之后,再把pre指向右子树的最右节点,再次构成一个双向有序链表
 *          4. 递归上述过程即可.
 */    


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
    TreeNode* Convert(TreeNode* pRootOfTree)
    {
        if(pRootOfTree == nullptr) return nullptr;
        TreeNode* pre = nullptr;
         
        convertHelper(pRootOfTree, pre); // 构造有序双向链表
         
        TreeNode* res = pRootOfTree;
        while(res ->left) // 遍历到双向链表的头节点并返回
            res = res ->left;
        return res;
    }
     
     // 中序遍历
    void convertHelper(TreeNode* cur, TreeNode*& pre)
    {
        if(cur == nullptr) return;
         
        convertHelper(cur ->left, pre);
         
        cur ->left = pre;
        if(pre) pre ->right = cur;
        pre = cur;
         
        convertHelper(cur ->right, pre);       
    }
};