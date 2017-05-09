package main

import "fmt"

func mergeSort(src []int) ([]int) {
  if len(src) == 0 {
    return nil
  } else if len(src) == 1 {
    return src
  } else if len(src) == 2 {
    //fmt.Println(">2")
    // base sort case
    if src[0] > src[1] {
      // 0 item is larger that 1 item.
      // swap and return it sorted.
      result := []int{src[1], src[0]}
      return result
    } else {
      return src
    }
  } else {
    // general case
    mid := len(src)/2
    odd := len(src)%2  // if odd, i'm adding the extra element to the left side.
    //fmt.Println(">s", src)
    //fmt.Println(">i", mid, odd, src[0:mid+odd], src[2:3])

    // recursion goes here.
    left := mergeSort(src[0:mid+odd])
    right := mergeSort(src[mid+odd:len(src)])

    //fmt.Println(left, "", right)

    // final list
    finalLen := len(left)+len(right)
    finalList := make([]int, finalLen)
    
    // counter for left and right list.
    li:=0
    lLen:=len(left) - 1
    ri:=0
    rLen:=len(right) - 1
    for i:=0 ; i<finalLen ; i++ {
      if li > lLen {
        // left list is done.
        finalList[i] = right[ri]
        ri++
      } else if ri > rLen {
        // right list is done.
        finalList[i] = left[li]
        li++
      } else {
        if left[li] < right[ri] {
          finalList[i] = left[li]
          li++
        } else {
          finalList[i] = right[ri]
          ri++
        }
      }
    }

    return finalList
  }
}

func consec(src []int) ([]int) {
  // guards
  if len(src) == 0 || len(src) == 1 {
    return src
  }

  // sorted src.
  sortedSrc := mergeSort(src)
  //fmt.Println(sortedSrc)

  // init the max length to the first item.
  mLen := 0
  mStart := 0
  mEnd := 0

  for i:=0; i<len(sortedSrc) ; i++ {
    // length of current sequence.
    cLen := 1
    cStart := i
    cEnd := i+1
    
    // increment max list test
    for j:=i+1 ; j<len(sortedSrc) ; j++ {
      if sortedSrc[j] == sortedSrc[j-1] + 1 {
        cLen++
        cEnd++
      } else {
        break
      }
    }

    // if this is longest chain test.
    if cLen > mLen {
      //fmt.Println(cLen, cStart, cEnd)
      mLen = cLen
      mEnd = cEnd
      // handling an edge case of if the cStart is at the start of the list.
      // in all other cases, you want to compare the first 2 items in the list.
      if cStart > 0 {
        mStart = cStart-1
      }
    }

    // shift the main counter to the possible next sequence.
    i += cLen
  }

  // construct the final list.
  dest := sortedSrc[mStart:mEnd]
  return dest
}

func main() {
  // l1 := []int{6,10,3,2,1,2}
  // result1 := mergeSort(l1)
  // fmt.Println(result1)

  l2 := []int{6,10,3,1,2}
  result2 := consec(l2)
  fmt.Println(result2)

  l3 := []int{93,94,1000,2,92,1001}
  result3 := consec(l3)
  fmt.Println(result3)

  l4 := []int{1,2,3,4,5,6,7,8,9,10}
  result4 := consec(l4)
  fmt.Println(result4)

  l5 := []int{2,1}
  result5 := consec(l5)
  fmt.Println(result5)


  l6 := []int{7}
  result6 := consec(l6)
  fmt.Println(result6)

  result7 := consec(nil)
  fmt.Println(result7)
}

/*

https://www.careercup.com/question?id=5656053464170496

Given an array of ints
return the largest consecutive sequence.

input
93,94,1000,2,92,1001

output
92,93,94

*/