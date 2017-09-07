package bintree

import("testing"
		"../util")


func TestBinTree(t *testing.T) {
	btree := New(util.IntComparer)
	btree.Add(100,100)
	btree.Add(50,50)
	btree.Add(200,200)	
}