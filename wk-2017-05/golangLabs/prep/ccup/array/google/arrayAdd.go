package main

import "fmt"

func add2Array(a, b []int) ([]int) {
  // guards
  if a == nil && b == nil {
    return nil
  } else if a == nil {
    return b
  } else if b == nil {
    return a
  } else if len(a) == 0 {
    return b
  } else if len(b) == 0 {
    return a
  }

  // setting result array.
  // pick the largest array + 1
  // figuring out the offsets
  cLen := 0
  aOff := 0
  bOff := 0
  if len(a) < len(b) {
    cLen = len(b) + 1
    aOff = len(b) - len(a)
  } else {
    cLen = len(a) + 1
    bOff = len(a) - len(b)
  }
  c := make([]int, cLen)

  remainder := 0
  // adding in.
  for i := 0 ; i<(len(c)-1) ; i++ {
    // iPOS is the number character position.
    iPos := (len(c)-i)-2
    aPos := iPos - aOff
    bPos := iPos - bOff
    cPos := iPos + 1

    //fmt.Println(iPos, i, len(a), len(b), aOff, bOff)
    //fmt.Println(iPos, ">>>", aPos, bPos, cPos)
    if aPos >= 0 && bPos >= 0 {
      c[cPos] = a[aPos] + b[bPos] + remainder
      //fmt.Println("> d", c[cPos], a[aPos], b[bPos])
      
      // remainder
      if c[cPos] >= 10 {
        remainder = 1
        c[cPos] %= 10
      } else {
        remainder = 0
      }
    } else if aPos >= 0 {
      //fmt.Println("> a", a[aPos])
      c[cPos] = a[aPos]
    } else if bPos >= 0 {
      //fmt.Println("> b", b[bPos])
      c[cPos] = b[bPos]
    }
  }

  // adding the last digit
  c[0] = c[0] + remainder
  
  // probably can do some cleanup.
  // to slice out the remainder.
  // and return that.

  return c
}



func main() {
  a1 := []int{1,2,3}
  b1 := []int{1,2,3}
  r1 := add2Array(a1, b1)
  fmt.Println("> 2 equal")
  fmt.Println(a1)
  fmt.Println(b1)
  fmt.Println(r1)

  a2 := []int{1,2,3,5}
  b2 := []int{1,2,3,5}
  r2 := add2Array(a2, b2)
  fmt.Println("> 2 equal. >= 10 for a digit")
  fmt.Println(a2)
  fmt.Println(b2)
  fmt.Println(r2)

  a3 := []int{1,2,3}
  b3 := []int{1,2,3,5}
  r3 := add2Array(a3, b3)
  fmt.Println("> 2 unequal.")
  fmt.Println(a3)
  fmt.Println(b3)
  fmt.Println(r3)

  //a4 := []int{1,2,3}
  //b4 := []int{1,2,3,5}
  r4 := add2Array(nil, nil)
  fmt.Println("> nils")
  //fmt.Println(a4)
  //fmt.Println(b4)
  fmt.Println(r4)

  a5 := []int{1,2,3}
  //b5 := []int{1,2,3,5}
  r5 := add2Array(a5, nil)
  fmt.Println("> 1 side nil.")
  fmt.Println(a5)
  //fmt.Println(b5)
  fmt.Println(r5)

  //a6 := []int{1,2,3}
  b6 := []int{1,2,3,5}
  r6 := add2Array(nil, b6)
  fmt.Println("> other side nil.")
  //fmt.Println(a6)
  fmt.Println(b6)
  fmt.Println(r6)

  a7 := []int  {1,2,4,5}
  b7 := []int{1,2,3,5,5}
  r7 := add2Array(a7, b7)
  fmt.Println("> 2 equal unequal. mutliple carries")
  fmt.Println(a7)
  fmt.Println(b7)
  fmt.Println(r7)
}

/*
  https://www.careercup.com/question?id=5631950045839360

  Arrays as digits
  Given 2 Arrays, add the digits.
  Assume a single digit.
  If both values are > 10, you need to carry the values over.

  SO

  [1, 2, 3]
  [2, 3, 5, 5]
  ============
  [2, 4, 7, 8]
*/