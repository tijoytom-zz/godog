package btree 

import ("testing"
        "fmt"
        "../util"
        )

func TestBTreeSplit(t *testing.T) {
  tree := New(IntComparer,3)
  tree.Add(10,10)	
  tree.Add(200,200)	
  tree.Add(100,100)
  
  arr := tree.ToArray()
  if arr[0] != 10 {   
	 
    t.Error("root is not 100")
  }
}

func TestBTree(t *testing.T) {
	tree := New(StringComparer,3)
	if tree == nil {
		t.Error("failed to create btree")
	}
	tree.Add("key1",10)
	
	tree.Add("key2",200)
	
	tree.Add("key3",100)
	arr := tree.ToArray()
	fmt.Println(arr)
	if len(arr) != 3 {
		t.Error("unexpected length")
	}
}
