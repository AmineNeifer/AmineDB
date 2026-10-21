package storage

const minChildren, maxChildren = 3, 5;

type BPTree struct {
	root *node
}

type node struct {
	e element
	children [maxChildren]*node
	crowd int // how many elements currently in the tree
	isLeaf bool // leaf when it doesnt have children
}

type element struct {
	key int
	value string
}


func NewNode(e element) *node {
	return &node{e: e, isLeaf: true}
}

func (b *BPTree) Insert(e element) {
	new := NewNode(e)
	if b == nil || b.root == nil {
		b = &BPTree{root: new}
	}

	b.root.Insert(new)
}

func (n *node) Insert(new *node) {
	if n.isLeaf && n.crowd < maxChildren { 
		mid := (maxChildren + 1) / 2
		midChild := n.children[mid]
		
		if midChild == nil {
			midChild = new
			return
		}

		if new.e.key > midChild.e.key {
			for i := mid+1; i < maxChildren; i++ {
				if new.e.key < n.children[i].e.key {
					copy(n.children[i+1:], n.children[i:maxChildren-2])
					n.children[i] = new
					return
				}
				if n.children[i] == nil {
					n.children[i] = new
					return
				}
			}
		}

		if new.e.key < midChild.e.key {
			for i:= mid-1; i >= 0; i-- {
				if new.e.key > n.children[i].e.key {
					copy(n.children[0:i-1], n.children[1:i])
					n.children[i] = new
					return
				}
				if n.children[i] == nil {
					n.children[i] = new
					return
				}
			}
		}

	}
}
