package list


type List struct {
	arr []interface{}
	size int
}

const (
 growthFactor =  float32(2.0)
 shrinkFactor = float32(0.25)
)

func (l *List) makeroom(nz int) {

  capacity := cap (l.arr)
  if nz >= capacity {
  	newcapacity := int(float32(capacity) * growthFactor)
  	newarr := make([]interface{},newcapacity,newcapacity)
  	copy(newarr,l.arr)
  	l.arr = newarr
  }
}

func New(capacity int) *List {
	l := List{}
	l.arr = make([]interface{},capacity,capacity)
	return &l
}


func (l *List) Add(values  ...interface{}) {

	nz := len(values) + l.size
	l.makeroom(nz)

	for _,v := range values {
	l.arr[l.size] = v
	l.size++
 }
}

func (l *List) Size() int {
 return l.size
} 

func (l *List) GetAt(idx int) (interface{} ,bool) {
  
  if idx < l.size {
  	return l.arr[idx] ,true 
  }
  return nil,false
}


func (l *List) Remove(idx int) bool {

	if idx >= l.size {
		return false
	}
	for i:= idx ;i<l.size - 1;i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.size--

	return true	
}
