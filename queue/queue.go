package queue

import ("../list")

type Queue struct {
	sz int
	l *list.List 
}

func New() *Queue {
	return  &Queue{l : list.New(2) , sz : 0 }
}

func(q *Queue) Enqueue(item interface{}) bool {	
	q.l.Add(item)	
	return true
}

func (q *Queue) Dequeue() (interface{}, bool) {  	
	 // In-efficient , but works 
	 item ,ok :=  q.l.GetAt(0)	 
	 q.l.Remove(0)
	 return item,ok
}


