package main

import "fmt"

func pushZeros(src []int) ([]int) {
  // guards. 
  if len(src) < 2 {
    return src
  }

  // using this as a second index to find non zero index.
  nonZeroI := -1
  for i:=0 ; i<len(src)-1 ; i++ {
    // You have to detect a zero.
    // This trigger the swap.
    if src[i] == 0 {
      // intial setup for the nonZeroI
      // this should only fire once.
      if nonZeroI == -1 {
        nonZeroI = i+1
      }

      // scan for non zero element
      for j:=nonZeroI ; j<len(src) ; j++ {
        if src[j] != 0 {
          // swap the two variables.
          swap := src[j]
          src[j] = src[i]
          src[i] = swap

          // you found a non zero element, 
          // set the index to the next value.
          nonZeroI = j + 1
          break
        }
      }
    }
  }

  return src
}



func main() {
  i1 := []int{0,1,2,3,4,5}
  o1 := pushZeros(i1)
  fmt.Println(o1)
}


/*
https://www.careercup.com/question?id=4846025567109120

You are given an array of single digit ints.
Push the zeros to the end of the list
Your solution must be O(n)

0 0 1 2 3 4 5
initial
if 0 pos is zero
swap 0 and 1.
and inc
0 0 1 2 3 4 5
0 1 0 2 3 4 5

scan for non zero.
and swap
you need 2 pointer.

0 0 1 2 3 0 5
1 0 0 2 3 0 5

*/