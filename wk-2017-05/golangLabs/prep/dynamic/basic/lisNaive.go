package main

import "fmt"

func licNaive(vals []int) int {
  fmt.Println(">", vals)
  if len(vals) == 0 {
    return 0
  }

  lis := 1

  curMax := vals[0]
  for i:=1 ; i<len(vals) ; i++ {
    if vals[i] > curMax {
      curMax = vals[i]
      lis++
    }
  }

  // recursion to compute the lis of the next character over.
  lisR := licNaive(vals[1:])

  fmt.Println(lis, lisR)

  if lisR > lis {
    return lisR
  } else {
    return lis
  }
}

func main() {
  i := []int{1}
  
  // i = []int{1}
  // fmt.Println(i, licNaive(i))
  
  // i = []int {5,2,1,4,5}
  // fmt.Println(i, licNaive(i))

  i = []int {9,10,1,2,3,4,5}
  fmt.Println(i, licNaive(i))
}

/*

Naive implementation of the LIS function.
At least I got that right...

http://www.geeksforgeeks.org/dynamic-programming-set-3-longest-increasing-subsequence/
Try to come with a good solution first before staring at the answer.....

You ahve a hit that its a list...

*/