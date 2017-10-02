package bintree

type Comparator func(a,b interface{}) int

type Node struct {
	key interface{}
	value interface{}
	parent *Node
    left *Node 
    right *Node
}

type BinTree struct {
	root *Node
	size int
	compare Comparator
}

func New(cmp Comparator) *BinTree {
	return &BinTree{compare:cmp}
}

func (t *BinTree) Left() *Node {
	if t == nil {
		return nil
	}
	var n *Node
	for c:= t.root;c != nil;c = c.left{
		n = c
	} 
	return n
}


func (t *BinTree) Right() *Node {
	if t == nil {
		return nil
	}
	var n *Node
	for c:= t.root;c != nil;c = c.right{
		n = c
	} 
	return n
}

func(node *Node) Next() *Node {
	
	if node == nil{
		return nil
	}

	// move right 
	n := node
	if n.right !=  nil  {
		n = n.right
		// in-order succesor is the left most node		
		for n.left != nil {
			n = n.left
		}
		return n
	}
	p := n.parent
	for p != nil && p.right == n {
		n = p 
		p= p.parent
	}
	return p
}

func(node *Node) Prev() *Node {
	
	if node == nil{
		return nil
	}

	// move left 
	n := node
	if n.left !=  nil  {
		n = n.left
		// in-order pre is the left most node		
		for n.right != nil {
			n = n.right
		}
		return n
	}
	p := n.parent
	for p != nil && p.left == n {
		n = p 
		p= p.parent
	}
	return p
}

func (t *BinTree) add(n *Node,key interface{},value interface{}) *Node {
	if n == nil {
		n = &Node{key:key,value:value}
	}
	cmp := t.compare(key,n.key)
	if cmp < 0 {
		node1 := t.add(n.left,key,value)
		n.left = node1
		node1.parent = n
	} else if cmp > 0 {
		node2 := t.add(n.right,key,value)
		n.right = node2
		node2.parent = n
	} else {		
		n.value = value
	}
	return n
}

func (t *BinTree) DeleteMin(x *Node) *Node {
	if x == nil {
		return nil
	}
	if x.left == nil {
	  return x.right
	}
	x.left = DeleteMin(x.left)
	return x
}


func (t *BinTree) delete(x *Node,key interface{}) (bool,*Node) {
	cmp := t.compare(key,x.key)
	if cmp < 0 {
		return delete(x.left,key)
	} else if cmp > 0 {
		return delete(x.right,key)
	} else {
		if x.right == nil {
			return true,x.left
		}
		temp := x 
		x.right = t.DeleteMin(x.right)
		x.left = temp.left
		return true ,x
	}
}
func(t *BinTree) Remove(key interface{}) bool {
	removed , x := t.delete(t.root,key)
	return removed
}

func (t *BinTree) Add(key interface{},value interface{}) {
	t.root = t.add(t.root,key,value)
}