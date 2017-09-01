package queue
import ("testing")

func TestQueue(t *testing.T) {
	q := New()
	for i:=0;i<100;i++ {	
		q.Enqueue(i)
	}

	for i:=0;i<99;i++ {	
		q.Dequeue()
	}
	
	if item,ok :=  q.Dequeue();!ok || item != 99 {
		t.Error("Dequeue failed:", item)
	}
}
