package buffer

type Buffer struct {
	head , tail int
	size int
	capacity int
	arr []interface{}
}

func New(capacity int) *Buffer {
	return &Buffer{capacity:capacity, arr : make([]interface{},capacity)}
}

func(b *Buffer) Push(val interface{}) bool {

	 if b.size == b.capacity {
		 return false
	 }
	 b.arr[b.head] = val
	 b.head = (b.head + 1) % b.capacity  
	 b.size++
	 return true
}

func (b* Buffer) Pop() (bool,interface{}) {
  if  b.size == 0 {
 	  return false,nil
  }
  v := b.arr[b.tail]
  b.tail = (b.tail + 1) % b.capacity
  return true,v 
}
