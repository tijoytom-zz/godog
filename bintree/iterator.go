package bintree
import ("../common")

type Iterator struct {
	tree *BinTree
	node *Node
	current position
}

type position byte

const (
	begin,between,end position = 0 ,1,2
)

func (t *BinTree) Iterator() common.ReverseIterator {
	return &Iterator{tree:t,node:nil,current:begin}
}

func (it *Iterator) Begin()  {
	it.current = begin 
	it.node = nil	
}

func (it *Iterator) Next() bool  {
 switch it.current {
	 case begin :
		it.current = between
		it.node = it.tree.Left()
	 case between:		
		it.node = it.node.Next()

	 }

	 if it.node == nil {
		 it.current = end
		 return false
	 }
	 return true
}

func (it *Iterator) Value() interface{} {
	return it.node.value
}

func (it *Iterator) Key() interface{} {
	return it.node.key
}

func (it *Iterator) First() bool {
	it.Begin()
	return it.Next()
}

func (it *Iterator) End()  {
	
	it.node = nil
	it.current = end
}

func (it *Iterator) Last() bool {
	it.End()
	return it.Prev()
}


func (it *Iterator) Prev() bool {
	switch it.current {
	case end :
	   it.current = between
	   it.node = it.tree.Right()
	case between:		
	   it.node = it.node.Prev()
	}

	if it.node == nil {
		it.current = begin
		return false
	}
	return true
}











