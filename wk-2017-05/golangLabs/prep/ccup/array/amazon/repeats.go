package main

import "fmt"

func repeats(src []int) {
  for i:=0 ; i<len(src) ; i++ {
    // some initial setup.
    numRepeats := 1
    curNum := src[i]
    newI := i // this is the offset.
  
    // count the repeats.
    for j:=i+1 ; j<len(src) ; j++ {
      //fmt.Println(j)
      if src[j] == curNum {
        numRepeats++
      } else {
        break
      }
    }

    // check if num repeats >1
    if numRepeats > 1 {
      fmt.Println(curNum, "repeats", numRepeats, "times.")
      newI = newI + numRepeats
    }

    // update the i counter
    i = newI
    //fmt.Println(">", i)
  }
}

func main() {
  l1 := []int{1,2,3,4,5,6,6}
  fmt.Println(l1)
  repeats(l1)

  l2 := []int{1,2,2,3,4,5,6,6}
  fmt.Println(l2)
  repeats(l2)

  l3 := []int{1}
  fmt.Println(l3)
  repeats(l3)

  l4 := []int{1,1,1,1,1}
  fmt.Println(l4)
  repeats(l4)
}

/*

https://www.careercup.com/question?id=5697297964859392

Find the repeats and list them out.
You are assuming that the list is sorted.

*/