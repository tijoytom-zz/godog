package tree

type Comparator func(a,b interface{}) int 

type Tree struct {
	root *Node
	cmp Comparator 
	size int
	order int
}

type Node struct{
	parent *Node
	items []*KV
	children []*Node
}

type KV struct {
   Key interface{}
   Value interface{}
}
func(n *Node) isLeaf() bool {
	return len( n.children) == 0
}

func (t* Tree) find(node* Node , kv *KV) (int,bool) {
	  
	low ,high := 0 ,len(node.items) - 1
	var mid int 
	for low <= high {
		mid = low + ((high - low ) >> 1)
		result := t.cmp(kv.Key,node.items[mid])
		switch  {
			case result > 0 :
				low = mid +1
			case result < 0 :
				high = mid - 1
			case result == 0:
				return mid,true
				}
	}
	return low,false
}

func(t* Tree) insertToLeaf(node* Node,kv *KV) {		
	i , found  := t.find(node,kv) 
	if found {
		node.items[i] = kv
		return 
	}
	node.items = append(node.items,nil)
	copy(node.items[i+1:] , node.items[i:])
	node.items[i] = kv
	// split 	
} 

func (t *Tree) insert(node *Node, kv *KV) bool {

	if node.isLeaf() {

	} else {

	}
}

func (t *Tree) Add(key interface{},value interface{} ) {
	kv := &KV{ Key : key ,Value :value }
	if t.root == nil {
		t.root = &Node{ parent:nil , items : []*KV{kv} , children:[]*Node{} }
		t.size++
	} else {
		
	}
}
