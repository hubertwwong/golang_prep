package main

import "fmt"

func lisNaive(vals []int) (int) {
  valsLen := len(vals)
  if valsLen == 0 {
    return 0
  } else if valsLen == 1 {
    return 1
  } else if valsLen == 2 {
    // is this a base case..
  }

  // how do you pick and recur
}

func main() {
  fmt.Println(lisNaive([]int{3,4,-1,0,6,2,3}))
}

/*
attempt number 2 on the problem..
*/