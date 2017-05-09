package main

import "fmt"

func spiral(vals [][]int) {
  jWidth := len(vals)-1
  jOffset := 0
  iWidth := len(vals[0])-1
  iOffset := 0

  //fmt.Println(jWidth, jOffset, iWidth, iOffset)

  for  {
    // top
    for i := iOffset ; i <= iWidth ; i++ {
      //fmt.Println(">t", iOffset, i)
      fmt.Println(vals[iOffset][i])
    }

    // right
    for i := jOffset+1 ; i <= jWidth ; i++ {
      //fmt.Println(">r ", i, jWidth)
      fmt.Println(vals[i][jWidth])
    }

    // bottom
    for i := iWidth-1 ; i >= iOffset ; i-- {
      //fmt.Println("> b", iWidth, i)
      fmt.Println(vals[iWidth][i])
    }

    // left
    for i := jWidth-1 ; i > jOffset ; i-- {
      //fmt.Println(">l", i, jOffset)
      fmt.Println(vals[i][jOffset])
    }

    // middle item.
    // if iWidth == 1 && jWidth == 1 {
    //   //fmt.Println(">middle item")
    //   fmt.Println(vals[len(vals)/2][len(vals[0])/2])
    // }

    // middle item 2 width
    // if iWidth == 2 && jWidth == 2 {
    //   fmt.Println(">middle item")
    //   fmt.Println(vals[len(vals)/2][len(vals[0])/2])
    // }

    // exit condition
    //fmt.Println(">b", iOffset, iWidth, jOffset, jWidth)
    iWidth -= 1
    iOffset += 1
    jWidth -= 1
    jOffset += 1
    //fmt.Println(">a", iOffset, iWidth, jOffset, jWidth)
    
    if iWidth <= 0 || jWidth <= 0 {
      break
    }
  }
}

func main() {
  // grid := [][]int{
  //   []int{1,2,3},
  //   []int{8,9,4},
  //   []int{7,6,5},
  // }

  // grid := [][]int{
  //   []int{1,2,3,4},
  //   []int{12,13,14,5},
  //   []int{11,16,15,6},
  //   []int{10,9,8,7},
  // }

  // grid := [][]int{
  //   []int{1,2},
  //   []int{4,3},
  // }

  // grid := [][]int{
  //   []int{1,2,3},
  //   []int{6,5,4},
  // }

  spiral(grid)
}

/*

Create a print funciton that prints a 2d array in a spiral.
This was asked a few times in asana.

10:03 start
notes...

Guessing you need a count var...
A level count.
Some way to compute the levels of rings.
use the width and 4 turns
sub 2 on each level.
1 and 2 are the smallest levels.

Lets try the square case first..
then work on the rectangle..

10:21
have a clue on how to do this..

Did a square one...
There is still an issue of rectangle ones...

*/