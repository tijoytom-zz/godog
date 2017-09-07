package util


func StringComparer(a,b interface{}) int {
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

func IntComparer(a,b interface{}) int {
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
