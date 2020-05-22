


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

 void convertNode(TreeNode* root, TreeNode*& last) {
   if (root == nullptr) return;

   TreeNode* current =  root;
   if (current->left != nullptr) 
        convertNode(current->left, last);
   current->left = last;
   if (last != nullptr) last->right = current;
   last = current;

   if (current->right != nullptr) 
        convertNode(current->right, last);
 }

 TreeNode* Convert(TreeNode* pRootOfTree) {
   if (pRootOfTree == nullptr) return nullptr;
   if (pRootOfTree->left == nullptr && pRootOfTree->right == nullptr)
     return pRootOfTree;

   TreeNode* lastNodeInList = nullptr;
   convertNode(pRootOfTree, lastNodeInList);

   TreeNode* head = lastNodeInList;
   while (head != nullptr && head->left != nullptr) {
     head = head->left;
        }

  return head;
 }
};