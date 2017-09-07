package btree 

import ("testing"
        "fmt"
        )

func stringcomparer(a,b interface{}) int {
  s1 := a.(string)
  s2 := b.(string)
  var min int 
  if len(s1) < len(s2) {
	  min = len(s1)
  } else {
	  min = len(s2)
  }

  d := 0
  for i:=0;i<min && d == 0;i++ {
	  d += int(s1[i]) - int(s2[i])
  } 
  if d == 0 {
    d = len(s1)-len(s2)
  }
  if d < 0 {
	  return -1
  }
  if d > 0 {
	  return 1
  }
  return 0
}

func intcomparer(a,b interface{}) int {
  s1 := a.(int)
  s2 := b.(int)
  
  if s1 < s2 {
    return -1
  }
  if s1>s2 {
    return 1
  }
  return 0  
}


func TestBTreeSplit(t *testing.T) {
  tree := New(intcomparer,3)
  tree.Add(10,10)	
  tree.Add(200,200)	
  tree.Add(100,100)
  
  arr := tree.ToArray()
  if arr[0] != 10 {   
	 
    t.Error("root is not 100")
  }
}

func TestBTree(t *testing.T) {
	tree := New(stringcomparer,3)
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
