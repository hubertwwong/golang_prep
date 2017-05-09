package main

import "fmt"

func permNext(vals []int) ([]int, bool) {
  // Find longest non-increasing suffix
  i := len(vals) - 1
  for ; i > 0 && vals[i - 1] >= vals[i] ; i-- {}
      
  // Now i is the head index of the suffix
  
  // Are we at the last permutation already?
  if i <= 0 { 
    return nil, false
  }

  // Let vals[i - 1] be the pivot
  // Find rightmost element that exceeds the pivot
  j := len(vals) - 1
  for ; vals[j] <= vals[i - 1] ; j-- {}

  // Now the value vals[j] will become the new pivot
  // Assertion: j >= i
  
  // Swap the pivot with j
  vals[i - 1], vals[j] = vals[j], vals[i - 1]
  
  // Reverse the suffix
  j = len(vals) - 1
  for ; i < j ; {
    vals[i], vals[j] = vals[j], vals[i]
    i++
    j--
  }
  
  // Successfully computed the next permutation
  return vals, true
}

func main() {
  s := []int{1,2,4,3}
  fmt.Println(permNext(s))
}

/*
Stole this...

https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
*/