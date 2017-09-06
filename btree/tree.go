package btree
import ("fmt")
type Comparator func(a,b interface{}) int

type entry struct {
	key interface{}
	value interface{}
}

type node struct {
	 parent *node
	 children []*node
	 entries  []*entry 
}

type BTree struct {
	root *node
	order int
	compare Comparator 
}

func New(cmp Comparator,order int) *BTree {
	return &BTree{order:order,compare:cmp}
} 

func(t *BTree) findkey(node *node,key interface{}) (bool,int) {
	low , high := 0,len(node.entries) - 1
	var mid int 
	for low <= high {
		mid = low + ((high-low)>>1)
		r := t.compare(key,node.entries[mid].key)
		switch{
			case r > 0:
				low = mid + 1
			case r < 0:
				high = mid - 1
			case r == 0:
				return true,mid
		}
	}
	return false,low 
}

func (n *node) isLeaf() bool {
	return len(n.children) == 0
}

func (t* BTree) insert_internal(node *node,e *entry) (bool,*node) {
  found,index := t.findkey(node,e.key)
  if node.isLeaf() {
	if found {
		node.entries[index] = e
		return false,node 
	} else {
	
	  node.entries = append(node.entries,nil)
	  copy(node.entries[index+1:],node.entries[index:])
	  
	  node.entries[index] = e
      return true,node
	}
  }
  return t.insert_internal(node.children[index],e)
}

func (t *BTree) need_split(node *node) bool {
	return len(node.entries) == t.order
}

func setparent(nodes []*node , parent *node) {
	for _,n := range nodes {
		n.parent = parent
	}
}

func (t *BTree) splitnonroot(n *node) {
 
	if !t.need_split(n) { 
		return 
	}
	p := n.parent
	middle := len(n.entries) / 2 // left leaning
	left  := &node{entries:append( []*entry(nil),n.entries[:middle]...),parent:p}
	right := &node{entries:append( []*entry(nil),n.entries[middle+1:]...),parent:p}
	
	if !n.isLeaf() {
		left.children = append([]*node(nil), n.children[:middle+1]...)
		right.children = append([]*node(nil), n.children[middle+1:]...)
		setparent(left.children,left)
		setparent(right.children,right)
	}

	_, index := t.findkey(p,n.entries[middle].key)
	p.entries = append(p.entries,nil)
	copy(p.entries[index+1:] ,p.entries[index:])
	p.entries[index] = n.entries[middle]

	p.children[index] = left
	p.children = append(p.children,nil)
	copy(p.children[index+2:],p.children[index+1:])
	p.children[index+1] = right	

	// recursive split parent 
	t.splitnonroot(p)
}

func(t *BTree) splitroot() {
	current_root := t.root
	if !t.need_split(current_root) {
		return 
	}	
	middle := len(current_root.entries)/2

	left  := &node{entries:append( []*entry(nil),current_root.entries[:middle]...),parent:current_root}
	right := &node{entries:append( []*entry(nil),current_root.entries[middle+1:]...),parent:current_root}
	

	if len(current_root.children) > 0 {
		left.children = append([]*node(nil), current_root.children[:middle+1]...)
		right.children = append([]*node(nil), current_root.children[middle+1:]...)
		setparent(left.children,left)
		setparent(right.children,right)	
	}
	
	new_root := &node{ entries:[]*entry{current_root.entries[middle] }, children:[]*node{left,right}}
	fmt.Println(new_root.entries[0].value)
	//fmt.Println(new_root.children[0].entries[0].value)
	//fmt.Println(new_root.children[1].entries[0].value)
	t.root = new_root	
}

func(t *BTree) splitnode(node *node) {
	if (node == t.root) {
		t.splitroot()
	} else {
		t.splitnonroot(node)
	}
}

func (t *BTree) Add(key interface{},value interface{}) bool {
   e :=  &entry{key:key,value:value}
   if t.root == nil {
	   t.root = &node{entries:[]*entry{e},children:[]* node{}}
	   return true
   }
   inserted , n := t.insert_internal(t.root,e)
   t.splitnode(n)
   return inserted
}



func (t* BTree) ToArray() []interface{} {
   var arr []interface{}
   var q []*node
   q = append(q,t.root)

   for len(q) > 0  {
	node := q[0]
	q = q[1:]
	if node.isLeaf() {
		for i:=0;i<len(node.entries);i++ {
			arr = append(arr,node.entries[i].value)
		}
	} else {
		for i:= 0;i<len(node.children);i++ {
			q  = append(q,node.children[i])
		}
	}
   }
   return arr
}

