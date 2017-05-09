package main

import "fmt"

// Using slices
// assuming its an array.
func toepliz(src [][]string) (bool) {
  if src == nil {
    return false
  }

  // is mxn check.
  if isMxn(src) == false {
    return false
  }

  // outer loop. m+n-1.
  // you start 1 in and finish 1 early.

  lenY := len(src)
  lenX := len(src[0])
  
  startX := 0
  startY := lenY - 1
  //fmt.Println(">",startX, startY, lenX, lenY)
  for i:=0 ; i < (lenX + lenY) - 1 ; i++ {
    //fmt.Println("> o")
    // So we have the initial character.
    initChar := src[startY][startX]
    // set the cur array poisiton
    curX := startX
    curY := startY
    for j:=0 ; j < (lenX + lenY) - 1 ; j++ {
      //fmt.Println("i s >", curX, curY)
      
      // increment or continue on if you hit the end of the cur loop.
      if curX < (lenX-1) && curY < (lenY-1) {
        //fmt.Println("i if", curX, curY)
        curX++
        curY++
      } else {
        //fmt.Println("i else", curX, curY)
        break
      }

      // character is same or diff.
      //fmt.Println("i c >", curX, curY)
      curChar := src[curY][curX] 
      if initChar != curChar {
        return false
      }
    }

    // increment x y start points.
    // mentally, we are moving up and then across
    if startY > 0 {
      startY--
    } else {
      startX++
    }
  }

  return true
}

/*
  technically, slices are array of array so you can the internal rows to different sizes.
  so there should be a check that the rows are are all of the same length.
*/

func isMxn(src [][]string) (bool) {
  isFull := false
  initialLen := len(src[0])

  for i:=1 ; i<len(src) ; i++ {
    if len(src[i]) != initialLen {
      return isFull
    }
  }
  
  isFull = true
  return isFull
}

func print2dArray(src [][]string) {
  for i:=0 ; i<len(src) ; i++ {
    fmt.Println(src[i])
  }
}



func main() {
  src1 := make([][]string, 4)
  src1[0] = []string{"6", "7", "8", "9", "2"}
  src1[1] = []string{"4", "6", "7", "8", "9"}
  src1[2] = []string{"1", "4", "6", "7", "8"}
  src1[3] = []string{"0", "1", "4", "6", "7"}
  out1 := toepliz(src1)
  print2dArray(src1)
  fmt.Println(out1)

  src2 := make([][]string, 4)
  src2[0] = []string{"6", "7", "8", "9", "2"}
  src2[1] = []string{"4", "6", "7", "8", "9"}
  src2[2] = []string{"1", "4", "6", "7", "8"}
  src2[3] = []string{"0", "2", "4", "6", "7"}
  out2 := toepliz(src2)
  print2dArray(src2)
  fmt.Println(out2)

  src3 := make([][]string, 4)
  src3[0] = []string{"6", "7", "0", "9", "2"}
  src3[1] = []string{"4", "6", "7", "8", "9"}
  src3[2] = []string{"1", "4", "6", "7", "8"}
  src3[3] = []string{"0", "1", "4", "6", "7"}
  out3 := toepliz(src3)
  print2dArray(src3)
  fmt.Println(out3)

  out4 := toepliz(nil)
  fmt.Println(out4)

  src5 := make([][]string, 2)
  src5[0] = []string{"6", "7", "8", "9", "2"}
  src5[1] = []string{"4", "6", "7", "8", "9"}
  out5 := toepliz(src5)
  print2dArray(src5)
  fmt.Println(out5)

  src6 := make([][]string, 1)
  src6[0] = []string{"6"}
  out6 := toepliz(src6)
  print2dArray(src6)
  fmt.Println(out6)

  src7 := make([][]string, 2)
  src7[0] = []string{"6", "7", "8", "9", "2"}
  src7[1] = []string{"4", "6", "7", "8", "9"}
  out7 := toepliz(src7)
  print2dArray(src7)
  fmt.Println(out7)

  src8 := make([][]string, 2)
  src8[0] = []string{"1", "4"}
  src8[1] = []string{"2", "3"}
  out8 := toepliz(src8)
  print2dArray(src8)
  fmt.Println(out8)
}

/**

https://www.careercup.com/question?id=5763139615326208

A matrix is "Toepliz" if each descending diagonal from left to right is constant. 
Given an M x N matrix write the method isToepliz to determine if a matrix is Toepliz.

input:
67892
46789
14678
01467

output:
true

*/