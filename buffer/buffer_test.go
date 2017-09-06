package buffer

import ("testing")

func TestBufferAdd(t *testing.T) {
  b := New(4)
  for i:=0;i<4;i++{
	  b.Push(i+1)
  }
  added := b.Push(10)
  if added {
	  t.Error("bounds error")
  }
  _,v := b.Pop()
  _,v = b.Pop()
  if v != 2 {
	  t.Error("popped value not 1")
  }
}