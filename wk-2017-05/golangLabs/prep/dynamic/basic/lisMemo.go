package main

import "fmt"

func licMemo(vals []int) int {
  // guard... if there are less than 1 item, just return that.
  if len(vals) <= 1 {
    return len(vals)
  }
  // construct the memo table.
  memo := make([]int, len(vals))
  for i:=0 ; i<len(memo) ; i++ {
    memo[i] = 1
  }

  // double loops.
  for i:=1 ; i<len(vals) ; i++ {
    for j:=0 ; j<i ; j++ {
      // we only care if the intial value is lower than the current max.
      if vals[i] > vals[j] {
        // if existing solution + 1 is better than what we have.
        if memo[j] + 1 > memo[i] {
          memo[i] = memo[j] + 1
        }
      }
    }
  }

  // scan the memo table for the largest value
  max := 0
  for i:=0 ; i<len(memo) ; i++ {
    if memo[i] > max {
      max = memo[i]
    }
  }
  fmt.Println(memo)

  return max
}

func main() {
  i := []int{1}
  
  // i = []int{1}
  // fmt.Println(i, licNaive(i))
  
  // i = []int {5,2,1,4,5}
  // fmt.Println(i, licNaive(i))

  i = []int {3,4,-1,0,6,2,3}
  fmt.Println(i, licMemo(i))
}

/*

04/12/2017

Take another shot at this..
02:58p
use tushar youtube guide and got this solution.
This was done using his example and tracing through the solution and comming up with the logic.

This is the n2 version....
The big thing with this vs. the naive version is that this is n2 and the naive version is n!

This is the top down version and not the memoization version...



04/11/2017

*[] and []* is different.... Say it out to figure out the difference.
First one is a pointer to an array. Second one is an array of pointers.

So the recursion goes first....


04/10/2017

Naive implementation of the LIS function.
At least I got that right...

http://www.geeksforgeeks.org/dynamic-programming-set-3-longest-increasing-subsequence/
Try to come with a good solution first before staring at the answer.....

You ahve a hit that its a list...

*/