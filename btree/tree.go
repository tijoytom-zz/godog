package btree
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

func (t* BTree) insert_internal(node *node,e *entry) bool {
  found,index := t.findkey(node,e.key)
  if node.isLeaf() {
	if found {
		node.entries[index] = e
		return false 
	} else {
	
	  node.entries = append(node.entries,nil)
	  copy(node.entries[index+1:],node.entries[index:])
	  
	  node.entries[index] = e
      return true
	}
  }
  return t.insert_internal(node.children[index],e)
}

func (t* BTree) Add(key interface{},value interface{}) bool {
   e :=  &entry{key:key,value:value}
   if t.root == nil {
	   t.root = &node{entries:[]*entry{e},children:[]* node{}}
	   return true
   }
   inserted := t.insert_internal(t.root,e)
   return inserted
}



func (t* BTree) ToArray() []interface{} {
   var arr []interface{}
   var q []*node

   node := t.root
   for node != nil || len(q) > 0  {
	if node.isLeaf() {
		for i:=0;i<len(node.entries);i++ {
			arr = append(arr,node.entries[i].value)
		}
	} else {
		for i:= 0;i<len(node.children);i++ {
			q  = append(q,node.children[i])
		}
	}
	if len(q) == 0 {
		break
	}
	node , q = q[0],q[1:]
   }
   return arr
}

