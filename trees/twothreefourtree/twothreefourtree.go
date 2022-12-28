package twothreefourtree

type Node struct {
	cap                  int
	keys                 []int
	l, m1, m2, r, parent *Node
	tree                 *Tree
}

type Tree struct {
	root *Node
	cap  int
}

func (node *Node) isLeaf() bool {
	if node.l == nil {
		return true
	}
	return false
}

func (node *Node) hasKey(key int) bool {
	for _, x := range node.keys {
		if x == key {
			return true
		}
	}
	return false
}

func (node *Node) split3Node() (lNode, rNode *Node) {
	if node == node.tree.root {
		lNode.cap = 1
		lNode.keys = []int{node.keys[0]}
		lNode.l = node.l
		lNode.r = node.m1
		lNode.parent = node
		lNode.tree = node.tree

		rNode.cap = 1
		rNode.keys = []int{node.keys[2]}
		rNode.l = node.m2
		rNode.r = node.r
		rNode.parent = node
		rNode.tree = node.tree

		node.cap = 1
		node.keys = []int{node.keys[1]}
		node.l = lNode
		node.r = rNode
		node.m1, node.m2 = nil, nil
	} else {
		parent := node.parent

		lNode.cap = 1
		lNode.keys = []int{node.keys[0]}
		lNode.l, lNode.m1, lNode.m2, lNode.r = node.l, nil, nil, node.m1
		lNode.parent = parent
		lNode.tree = node.tree

		rNode.cap = 1
		rNode.keys = []int{node.keys[2]}
		rNode.l, rNode.m1, rNode.m2, rNode.r = node.m2, nil, nil, node.r
		rNode.parent = parent
		rNode.tree = node.tree

		midKey := node.keys[1]
		if parent.cap == 1 {
			if midKey < parent.keys[0] {
				parent.keys = append([]int{midKey}, parent.keys[0])
			} else {
				parent.keys = append(parent.keys, midKey)
			}
		} else if parent.cap == 2 {
			if midKey < parent.keys[0] {
				parent.keys = append([]int{midKey}, parent.keys...)
			} else if parent.keys[0] < midKey && midKey < parent.keys[1] {
				parent.keys = []int{parent.keys[0], midKey, parent.keys[1]}
			} else {
				parent.keys = append(parent.keys, midKey)
			}
		}
	}
	return
}

func (node *Node) insertKeyToLeaf(key int) {
	if node.hasKey(key) {
		return
	}
	if node.cap == 1 {
		if node.keys[0] < key {
			node.keys = append(node.keys, key)
		} else {
			node.keys = append([]int{key}, node.keys[0])
		}
		node.cap++
		return
	}
	if node.cap == 2 {
		if key < node.keys[0] {
			node.keys = append([]int{key}, node.keys...)
		} else if node.keys[0] < key && key < node.keys[1] {
			node.keys = []int{node.keys[0], key, node.keys[1]}
		} else {
			node.keys = append(node.keys, key)
		}
		node.cap++
		return
	}
	if node.cap == 3 {
		var stack []*Node
		parent := node.parent
		for parent != nil && parent.cap == 3 {
			stack = append(stack, parent)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			stack[i].split3Node()
		}
		mid := node.keys[1]
		l, r := node.split3Node()
		if key < mid {
			if key < l.keys[0] {
				l.keys = []int{key, l.keys[0]}
			} else {
				l.keys = append(l.keys, key)
			}
			l.cap++
		} else {
			if key < r.keys[0] {
				r.keys = []int{key, r.keys[0]}
			} else {
				r.keys = append(r.keys, key)
			}
			r.cap++
		}
	}
}

func (tree *Tree) Contain(key int) bool {
	return false
}

func (tree *Tree) Insert(key int) {
	if tree.cap == 0 {
		node := new(Node)
		node.cap = 1
		node.keys = append([]int{}, key)
	} else {
		node := tree.root
		for !node.isLeaf() {
			// 检查该结点是否已经包含key
			if node.hasKey(key) {
				return
			}
			// node为2结点
			if node.cap == 1 {
				if key < node.keys[0] {
					node = node.l
				} else {
					node = node.r
				}
			} else if node.cap == 2 {
				if key < node.keys[0] {
					node = node.l
				} else if node.keys[0] < key && key < node.keys[1] {
					node = node.m1
				} else {
					node = node.r
				}
			} else if node.cap == 3 {
				if key < node.keys[0] {
					node = node.l
				} else if node.keys[0] < key && key < node.keys[1] {
					node = node.m1
				} else if node.keys[1] < key && key < node.keys[2] {
					node = node.m2
				} else {
					node = node.r
				}
			}
		}
		// now, find the candidate node(leaf) to insert the key
		node.insertKeyToLeaf(key)
	}
	tree.cap++
}

func (tree *Tree) Delete(key int) {

}

func NewTree() *Tree {
	return nil
}

func (tree *Tree) PrettyPrint() {}
