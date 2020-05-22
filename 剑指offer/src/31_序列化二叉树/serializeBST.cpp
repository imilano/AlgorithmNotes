/*  题目描述
 *     请实现两个函数，分别用来序列化和反序列化二叉树
 *     二叉树的序列化是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树遍历方式来进行修改，序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#），以 ！ 表示一个结点值的结束（value!）。
 *     二叉树的反序列化是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。
 * 
 * 题解:
 *    采用前序遍历的方法来序列化二叉树，对于每一个节点的状态都要加以标志，比如空节点以 $ 这样的特殊符号表示.
 *    之所以用前序遍历的方法,是因为在反序列化时,第一个读到的就是根节点,从而可以方便的构造成出二叉树.
 *    
 * 
 *  
 * 
 */
 

 #include <iostream>
 #include <cstring>
 #include <cstdlib>

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
    char* Serialize(TreeNode *root) {
      if (!root) return "#";
      string s = to_string(root->val);
      s.push_back(',');

      char* left = Serialize(root->left);
      char* right = Serialize(root->right);

      char* ret = new char[strlen(left) + strlen(right) + s.size()];
      strcpy(ret, s.c_str());
      strcat(ret, left);
      strcat(ret, right);
      return ret;
    }

    TreeNode* decode(char*& str) {
      if (*str == '#') {
        str++;
        return nullptr;
      }

      int num = 0;
      while (*str != ',') num = num * 10 + (*(str++) - '0');
      str++;
      TreeNode* root = new TreeNode(num);
      root->left = decode(str);
      root->right = decode(str);
      return root;
    }

    TreeNode* Deserialize(char *str) {
      return decode(str);
    }
};



//  void serializeBST(TreeNode* root, ostream& out) {
//          if (root == nullptr) {
//            out << "#";
//            return;
//          }
//          out << root->val << ",";
//          serializeBST(root->left, out);
//          serializeBST(root->right, out);
//  }
//  bool readStream(ostream& stream, int* number) {
//          // 从ostream中读出一个数字或字符$，如果是数字，返回true，否则返false
//  }

//  void deSerializeBST(TreeNode** root, ostream& out) { 
//         int number;
//         if (readStream(out,&number)) {
//           *root = new TreeNode(number);
//           (*root)->left = (*root)->right = nullptr;

//           deSerializeBST(&(*root)->left, out);
//           deSerializeBST(&(*root)->right, out);
//         }
//  }
 