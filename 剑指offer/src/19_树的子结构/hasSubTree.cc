// 题目描述
// 	输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）

// 题解：
// 	现在tree1上找到和tree2根节点值相同的节点，若找到，则判断tree1的左右子数和tree2的左右子树的关系，如果二者的左右子树都相等，那么返回true，如果不等，返回false；如果根节点不等，则返回false

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
//         bool hasSubTree2(TreeNode* r1,TreeNode* r2) {
//                 if (r1 == nullptrptr)
//                         return false;
//                 if (r2 == nullptrptr)
//                         return true;
//                 if (r1->val != r2->val)
//                         return false;
//                 return hasSubTree2(r1->left, r2->left) && hasSubTree2(r1->right, r2->right);
//         }
//     bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
//     {
//         bool hassubtree = false;
//         if (pRoot2 == nullptrptr)
//                 return false;
//         if (pRoot2 == nullptrptr)
//                 return false;

        
//         hassubtree = hasSubTree2(pRoot1, pRoot2);
//         if(!hassubtree)
//                 hassubtree = HasSubtree(pRoot1->left, pRoot2);
//         if(!hassubtree)
//                 hassubtree = HasSubtree(pRoot1->right, pRoot2);


//         return hassubtree;
//     }


// 短路求值特性
  bool HasSubtree(TreeNode* pRoot1, TreeNode* pRoot2)
  {
		if(pRoot2 == nullptr || pRoot1 == nullptr )
				return false;
		return isSubtree(pRoot1, pRoot2)|| HasSubtree(pRoot1->left,pRoot2) || HasSubtree(pRoot1->right,pRoot2);
  }

  bool isSubtree(TreeNode* pRoot1 , TreeNode* pRoot2){
		if(pRoot2 == nullptr)
				return true;
		if(pRoot1 == nullptr)
				return false;
		return pRoot1->val == pRoot2->val && isSubtree(pRoot1->left,pRoot2->left) && isSubtree(pRoot1->right,pRoot2->right);
  	}
  };