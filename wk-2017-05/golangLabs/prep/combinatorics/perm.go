package main

import "fmt"

func printInt(vals [][]int) {
   for i := 0 ; i < len(vals) ; i++ {
       fmt.Println(vals[i])
   } 
}

// return a list of permutations...
func perm(vals []int) [][]int {
  // base case
  if len(vals) <= 1 {
    res := make([][]int, 0)
    res = append(res, vals)
    return res
  }

  result := make([][]int, 0)
  for i := 0 ; i < len(vals) ; i++ {
    // construct the subsequence to pass in to the recursion.
    curSubSequence := make([]int, 0)
    if i+1 == len(vals) {
      curSubSequence = vals[:i]
    } else if i == 0 {
      curSubSequence = vals[i+1:]
    } else {
      left := vals[0:i]
      right := vals[i+1:]
      //  append stuff together
      for i := 0 ; i < len(left) ; i++ {
        curSubSequence = append(curSubSequence, left[i])
      }
      for i := 0 ; i < len(right) ; i++ {
        curSubSequence = append(curSubSequence, right[i])
      }
    }
    
    // passing in curSubSequence
    resA := perm(curSubSequence)

    // append on result to mega results
    for j := 0 ; j < len(resA) ; j++ {
      curRow := make([]int, 0)
      curRow = append(curRow, vals[i])          // pass in the sliced off value
      curRow = append(curRow, resA[j]...)       // pass in recursion result
      result = append(result, curRow)
    }
  }

  return result
}

func main() {
  res := perm([]int{1,2,3,4})
  printInt(res)
}