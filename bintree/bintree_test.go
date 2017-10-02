package bintree

import("testing"
		"../util")

func TestBinTreePrev(t *testing.T) {
	btree := New(util.IntComparer)
	btree.Add(2,100)
	btree.Add(3,50)
	btree.Add(1,200)	
	it := btree.Iterator()
	it.End()
	count := 3
	for it.Prev() {
	 if count != it.Key() {
		t.Error("Unexpected value")
	 }
	 count--
	}
} 

func TestBinTreeNext(t *testing.T) {
	btree := New(util.IntComparer)
	btree.Add(2,100)
	btree.Add(3,50)
	btree.Add(1,200)	
	it := btree.Iterator()
	it.Begin()
	
	var count int
	for it.Next() {
	  count++
	  if count != it.Key() {
		  t.Error("Unexpected value")
	  }
	}
}