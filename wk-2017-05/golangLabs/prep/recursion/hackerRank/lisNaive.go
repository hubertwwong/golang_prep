package main

import "fmt"

func lisNaive(vals []int) (int) {
  valsLen := len(vals)
  if valsLen == 0 {
    return 0
  }

  max := 1
  for i:=0 ; i<valsLen ; i++ {
    curMax := 1
    curMaxVal := vals[i]

    // increment the sequence.
    // see if there is an increasing sequence.
    for j:=i+1 ; j<valsLen ; j++ {
      if vals[j] > curMaxVal {
        curMaxVal = vals[j]
        curMax++
      }
      fmt.Println(">valsj", vals[j], "cmv", curMaxVal, "j", j, "curMax", curMax)
    }
    fmt.Println();

    // check to see if we have a new max.
    if curMax > max {
      max = curMax
    }
  }

  return max
}

func main() {
  //fmt.Println(lisNaive([]int{5, 4, 1, 2, 3}))
  //fmt.Println(lisNaive([]int{4, 2, 4, 5, 3, 7}))
  //fmt.Println(lisNaive([]int{5, 2, 7, 4, 3, 8}))
  fmt.Println(lisNaive([]int{3,4,-1,0,6,2,3}))
}

/*

lens vals..

http://www.lintcode.com/en/problem/longest-increasing-subsequence/
03:23 start
03:32 end...

NOTE.... THIS IS WRONG>>>>>>

you have to pick or skip the issue...
you can do a quick comparision......

*/