package ch13_red_black_tree

import "fmt"

//红黑树的性质：
//	- 每个节点要么是红色，要么是黑色
//	- 根节点是黑色
//	- 叶节点是黑色（叶子节点是NIL节点）
//	- 红节点的两个子节点一定是黑色（不能出现两个红色节点相连的情况）
//	- 任一节点到其每个叶子节点的黑高一致
//可以出现两个黑色节点连在一起的情况。红黑树可以在O(lgn)时间内完成查找、插入和删除操作，n为树中元素数目。
//
//红黑树如何达到平衡：自旋（类似于AVL树）和着色
//
//插入的节点总是预着色为红节点（这样的插入仅只可能违背性质4，故而需要进行调整），红黑树的插入分分8种情景（实际旋转过程与AVL树的旋转过程差距不大）：
//	- 如果红黑树是空树
//		直接将插入节点作为根节点，同时将color设为黑。
//	- 当前节点的key已存在
//		已平衡，更新当前节点的值为插入节点的值
//	- 插入节点的父节点为黑节点
//		直接插入，不违反性质五
//	- 插入节点的父节点为红节点
//		根据红黑树性质，根节点总是黑色，而父节点为红色，所以父节点不可能为根节点，所以插入节点总是存在祖父节点。根据插入节点的叔叔节点及祖父节点的颜色，又需要分情况讨论：
//			-	叔叔节点存在并且为红色。
//				-	如果祖父节点为黑色，叔叔节点和父节点为红色，那么将叔叔节点和父节点变为黑色，祖父节点变为红色。如果祖父节点不是根节点，那么祖父节点的父节点就是红色，
//					此时出现了两个红节点相连的情况，此时将祖父节点视为新插入节点，继续向上进行调整；
//				-	如果祖父节点就是根节点，那么根据根节点必须为黑色的性质，直接将祖父节点
//					设为黑色，此时红黑树的黑高增加。
//			-	叔叔节点不存在或为黑节点，并且插入节点的父节点是祖父节点的左子节点，分情况讨论
//				-	插入节点是父节点的左子节点，对父节点进行右旋，然后父节点作为当前黑色根节点，插入节点和父节点的祖父节点作为父节点的两个红色子节点。
//				-	插入节点是父节点的右子节点，对父节点先左旋，再右旋
//			-	叔叔节点不存在或为黑节点，且插入节点的父节点是祖父节点的右子节点。情况与上一种对称。
//
//
//红黑树的删除分为两步，首先查找待删除节点，其次在删除之后要自平衡。删除节点之后还要找节点来替代删除节点的位置，以免断链。找替代节点有三种情形：
//	-	若删除节点无子节点，直接删除
//	-	若删除节点只要一个子节点，用子节点替换删除节点
//	-	若删除节点有两个子节点，用后继节点（大于删除节点的最小节点，也就是删除节点右子树的最左节点，也可以是删除节点左子树的最右节点）替换删除节点。
//
//关于删除节点的一个重要思路是：删除节点被替代后，在不考虑键值对的情况下，对于树来说，可以认为删除的是替代节点。删除中情况二，如果删除节点的子节点没有子节点了，那么在值替换后，
//相当于直接删除子节点；如果删除节点的子节点还有子节点，那么在替换后，就相当于删除一个存在左右子节点的节点，从而转化为情景三，情景三根据删除的思路，又可以转换为情景1.
//
//删除可以分为以下几种情景：
//	-	替换节点是红色节点
//		只需把替换节点换到删除节点的位置，将替换节点的颜色设为删除节点的颜色即可
//	-	替换节点是黑色
//		-	替换节点是其父节点的左子节点
//			-	替换节点的兄弟节点是红节点
//				将替换节点的兄弟节点设为黑色，然后将该兄弟节点的父节点设为红色，然后对该兄弟节点进行左旋。
//			- 	替换节点的兄弟节点是黑节点
//				此时无法确定该兄弟节点的父节点和子节点的颜色，需要分情况讨论
//				-	替换节点R的兄弟节点S的右子节点SR是红节点，左子节点SL颜色任意
//					将S设为父节点P的颜色，然后将SR设为红色，对S进行左旋
//				-	替换节点R的兄弟节点S的右子树是黑节点，左子树是红节点
//					将S设为红色，SL右旋。
//				-	替换节点R的兄弟节点S的左右子节点SL、SR都是黑节点
//					把S设为红色，把S的父节点P设为删除节点。
//		-	替换节点R是其父节点P的右子节点
//			-	替换节点R的兄弟节点S是红色
//				将S设为黑色，将S的父节点P设为红色，然后对S进行右旋
//			-	替换节点的兄弟节点S是黑节点
//				-	替换节点R的兄弟节点S的左子节点SL是红节点，右子节点SR颜色任意
//					将S设为父节点P的颜色，然后将SL设为黑色，对S右旋
//				-	替换节点R的兄弟节点S的左子节点是黑节点，右子节点是红节点
//					将S设为红色，SR设为黑色，对SR左旋
//				-	替换节点的兄弟节点的子节点都为黑节点
//					将S设为红色，把S的父节点P设为删除节点

type color bool

const (
	black, red color = true, false
)

// Tree holds elements of the red-black tree
type Tree struct {
	Root *Node
	size int
	//Comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	Key    interface{}
	Value  interface{}
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Put(key interface{}, value interface{}) {
	var insertedNode *Node
	if tree.Root == nil {
		// Assert key is of comparator's type for initial tree
		//tree.Comparator(key, key)
		tree.Root = &Node{Key: key, Value: value, color: red}
		insertedNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		for loop {
			// compare := tree.Comparator(key, node.Key)
			var compare int
			if key.(int) < node.Key.(int) {
				compare = -1
			} else if key.(int) == node.Key.(int) {
				compare = 0
			} else {
				compare = 1
			}
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node{Key: key, Value: value, color: red}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node{Key: key, Value: value, color: red}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
func (tree *Tree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

// Remove remove the node from the tree by key.
func (tree *Tree) Remove(key interface{}) {
	var child *Node
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size--
}

// Empty 判断数是否为空树
func (tree *Tree) Empty() bool {
	return tree.size == 0
}

// Size	返回树中的节点数 .
func (tree *Tree) Size() int {
	return tree.size
}

// Keys returns all keys in-order
//func (tree *Tree) Keys() []interface{} {
//	keys := make([]interface{}, tree.size)
//	it := tree.Iterator()
//	for i := 0; it.Next(); i++ {
//		keys[i] = it.Key()
//	}
//	return keys
//}

// Values returns all values in-order based on the key.
//func (tree *Tree) Values() []interface{} {
//	values := make([]interface{}, tree.size)
//	it := tree.Iterator()
//	for i := 0; it.Next(); i++ {
//		values[i] = it.Value()
//	}
//	return values
//}

// Left 返回树的最左（小）节点.
func (tree *Tree) Left() *Node {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

// Right 返回树的最右(最大)节点.
func (tree *Tree) Right() *Node {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
// Floor 返回小于等于当前节点的最小节点
func (tree *Tree) Floor(key interface{}) (floor *Node, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		// compare := tree.Comparator(key, node.Key)
		var compare int
		if key.(int) == node.Key.(int) {
			compare = 0
		} else if key.(int) < node.Key.(int) {
			compare = -1
		} else {
			compare = 1
		}
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
// 返回大于等于该节点的最小节点
func (tree *Tree) Ceiling(key interface{}) (ceiling *Node, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		//compare := tree.Comparator(key, node.Key)
		var compare int
		if key.(int) == node.Key.(int) {
			compare = 0
		} else if key.(int) < node.Key.(int) {
			compare = -1
		} else {
			compare = 1
		}
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

// Clear 删除树的所有节点
func (tree *Tree) Clear() {
	tree.Root = nil
	tree.size = 0
}

// String 返回树的字符形式表示
func (tree *Tree) String() string {
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.Key)
}

// 输出树
func output(node *Node, prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

// 查找key
func (tree *Tree) lookup(key interface{}) *Node {
	node := tree.Root
	for node != nil {
		// compare := tree.Comparator(key, node.Key)
		var compare int
		if key.(int) == node.Key.(int) {
			compare = 0
		} else if key.(int) < node.Key.(int) {
			compare = -1
		} else {
			compare = 1
		}

		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

// 返回祖父节点
func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

// 返回叔叔节点
func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

// 返回兄弟节点
func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

// 左旋
func (tree *Tree) rotateLeft(node *Node) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

// 右旋
func (tree *Tree) rotateRight(node *Node) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

// 节点替换
func (tree *Tree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

// 若树是空树，直接插入，并将节点颜色设为黑色；如果非空树，转为case2
func (tree *Tree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

// 如果插入节点的父节点是黑节点，不破坏性质，插入成功，直接返回；如果插入节点的父节点是红节点，进入case3
func (tree *Tree) insertCase2(node *Node) {
	if nodeColor(node.Parent) == black {
		return
	}
	tree.insertCase3(node)
}

// 需要根据祖父节点和叔叔节点的颜色对树进行调整
// 	-	如果叔叔节点和父节点都是红色，由于树之前已经平衡，故而祖父节点必然为黑色。此时将父节点和叔叔节点设为黑色，祖父节点设为红色，
// 		根据祖父节点与其父节点的情况，又需要分情况进行讨论，此时将祖父节点看做插入节点，转换到case1。
// 	-	如果叔叔节点是黑色，父节点是红色，进入case4
func (tree *Tree) insertCase3(node *Node) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

// 如果叔叔节点是黑色或空，父节点是红色（交错型）：
// 		-	如果当前节点是其父节点的右子节点，而其父节点是祖父节点的左子节点，则对插入节点左旋，然后对插入节点的父节点进入case5
//		-	如果当前节点是其父节点的左子节点，而其父节点是祖父节点的右子节点，则对插入节点右旋，然后对插入节点的父节点进入case5
func (tree *Tree) insertCase4(node *Node) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

// 父节点为红色，叔叔节点黑色或空：
// 	-	新节点是父节点的左节点，而父节点是祖父节点的左节点，对父节点右旋，并将父节点转换为黑色，插入节点和子节点作为父节点的子节点作为红色
//	-	新节点是父节点的右节点，而父节点是祖父接地那的右节点，对父节点左旋，并将插入阶段转为黑色，插入节点和子节点作为父节点的子节点作为红色
func (tree *Tree) insertCase5(node *Node) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (tree *Tree) deleteCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		tree.rotateRight(node.Parent)
	}
}

func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}
