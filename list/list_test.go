package list 

import ("testing")

func TestListRemove(t *testing.T) {

 l := New(2)
 for i:=0 ;i<100;i++ {
	l.Add(i)
  }
  l.Remove(99);

  if l.Size() != 99 {
  	t.Error("Failed")
  }

  l.Remove(50)
  v,_ := l.GetAt(50)

  if v != 51  {
  	t.Error("Failed")
  }
}

func TestListAdd(t *testing.T) {

	l := New(2)
	for i:=0 ;i<100;i++{
		l.Add(i)
	}

	if l.Size() != 100 {
		t.Error("Failed")
	}

	for i:= 0;i<100;i++ {
		v,_ := l.GetAt(i)
		if v != i {
			t.Error(i)
		}
	}
}