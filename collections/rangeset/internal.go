package rangeset

//// define nodes

type node interface {
	Begin() int64
	Parent() node
}

type nodeBase struct {
	parent      node
	parentIndex int
}

func (nb *nodeBase) Parent() node {
	return nb.parent
}

/*
branchNode records the index of the ranges.

*/
type branchNode struct {
	nodeBase
	begList  []int64 // 1 < len(begList) < maxChildrenCount
	nodeList []node  // for i in 1...len(begList), begList[i] == nodeList[i].Begin()
}

func (b *branchNode) Begin() int64 {
	return b.begList[0]
}

/*
leafNode records the ranges
*/
type leafNode struct {
	nodeBase
	next  *leafNode
	front *leafNode

	beg int64
	end int64
}

func (l *leafNode) Begin() int64 {
	return l.beg
}

/// define BPTree

type rangeBPTree struct {
	maxChildrenCount int

	root node
	head *leafNode
	tail *leafNode
}

type pair struct {
	beg int64
	end int64
}

func (t *rangeBPTree) newBranchNode() *branchNode {
	return &branchNode{
		begList:  make([]int64, 0, t.maxChildrenCount),
		nodeList: make([]node, 0, t.maxChildrenCount),
	}
}

func newLeafNode(p *pair) *leafNode {
	return &leafNode{
		beg: p.beg,
		end: p.end,
	}
}

/*
AddRange applies a range to the rangeBPTree.

requires:
	1. beg < end
*/
func (t *rangeBPTree) AddRange(p pair) {
	if t.root == nil {
		newLeaf := newLeafNode(&p)

		t.root = newLeaf
		t.head = newLeaf
		t.tail = newLeaf

		return
	}

	switch rootNode := t.root.(type) {
	case *branchNode:
		if p.beg < rootNode.Begin() {
			t.addMinRangeToSubTree(rootNode, &p)
		} else {
			t.addNormalRangeToSubTree(rootNode, &p)
		}
	case *leafNode:
		t.addRangeToSingleLeafTree(rootNode, &p)
	}
}

func (t *rangeBPTree) addRangeToSingleLeafTree(root *leafNode, p *pair) {
	if p.beg < root.beg {
		if p.end >= root.beg {
			root.beg = p.beg

			if p.end > root.end {
				root.end = p.end
			}

			return
		}

		// p.beg < p.end < root.beg

		// insert a new leafNode as the head leafNode

		newRoot := t.newBranchNode()
		newLeaf := newLeafNode(p)

		// insert into link-list of ranges
		t.head = newLeaf
		newLeaf.next = root
		root.front = newLeaf

		// build index in shape of tree
		t.root = newRoot
		root.parent = newRoot
		root.parentIndex = 1
		newLeaf.parent = newRoot
		newLeaf.parentIndex = 0
		newRoot.begList = append(newRoot.begList, p.beg, root.beg)
		newRoot.nodeList = append(newRoot.nodeList, newLeaf, root)

		return
	}

	// p.beg >= root.beg

	if p.end <= root.end {
		return
	}

	// p.end > root.end

	if p.beg <= root.end {
		root.end = p.end
		return
	}

	// root.end < p.beg < p.end

	// insert a new leafNode as the tail leafNode

	newRoot := t.newBranchNode()
	newLeaf := newLeafNode(p)

	// insert into link-list of ranges
	t.tail = newLeaf
	newLeaf.front = root
	root.next = newLeaf

	// build index in shape of tree
	t.root = newRoot
	root.parent = newRoot
	root.parentIndex = 0
	newLeaf.parent = newRoot
	newLeaf.parentIndex = 1
	newRoot.begList = append(newRoot.begList, root.beg, p.beg)
	newRoot.nodeList = append(newRoot.nodeList, root, newLeaf)
}

/*
addNormalRangeToSubTree applies a normal range to the subtree of rangeBPTree.

requires:
	1. root.Begin() <= p.beg
*/
func (t *rangeBPTree) addNormalRangeToSubTree(root *branchNode, p *pair) {
	if DEBUG {
		if root.Begin() > p.beg {

		}
	}
}

/*
addMinRangeToSubTree applies a special range to the subtree of rangeBPTree.

requires:
	1. p.beg < root.Begin()
	2. root.parent == nil or root.parent.Begin() == p.beg
*/
func (t *rangeBPTree) addMinRangeToSubTree(root *branchNode, p *pair) {

}

func (t *rangeBPTree) mergeLeaf(l *leafNode) {

}

func (t *rangeBPTree) delNodeFromBranch(b *branchNode, index int) {

}

/*
























 */
