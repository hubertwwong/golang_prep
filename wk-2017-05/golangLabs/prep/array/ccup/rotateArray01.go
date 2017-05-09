package main

import "fmt"

// dumb version
func rotate(vals []int, numPos int) ([]int) {
  valsLen := len(vals)
  if valsLen == 0 {
    return nil
  }
  if numPos == 0 {
    return vals
  }

  // absolute number of the number of position you are shifting.
  posNumPos := numPos
  if numPos < 0 {
    posNumPos = numPos * -1
  }

  for i := 0 ; i<posNumPos ; i++ {
    if numPos > 0 {
      lastItem := vals[valsLen-1]
      for j := valsLen-2 ; j>=0 ; j-- {
        vals[j+1] = vals[j]
      }
      vals[0] = lastItem
    } else {
      firstItem := vals[0]
      for j := 0 ; j<valsLen-1 ; j++ {
        vals[j] = vals[j+1]
      }
      vals[valsLen-1] = firstItem
    } 
  }

  return vals
}

func main() {
  fmt.Println(rotate([]int{1,2,3,4}, 2))
  fmt.Println(rotate([]int{1,2,3,4}, -1))
  fmt.Println(rotate([]int{1,2,3,4}, 0))
  fmt.Println(rotate(nil, 2)) 
  fmt.Println(rotate([]int{1}, 2))
  fmt.Println(rotate([]int{1,2}, 1))  
}