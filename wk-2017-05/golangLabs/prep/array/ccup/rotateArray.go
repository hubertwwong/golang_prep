package main

import "fmt"

func posInArray(start, end, arrayLen int) (int) {
  fmt.Println(start, end, arrayLen)
  rawPos := start + end
  rawPos %= arrayLen
  fmt.Println(">", rawPos)
  if rawPos < 0 {
    return arrayLen + 1 + rawPos
  } else {
    return rawPos
  }
}

func rotate(vals []int, pos int) ([]int) {
  valsLen := len(vals)
  if valsLen  <= 0 {
    return nil
  }

  // copy array over.
  //finalVals := make([]int, valsLen)
  //copy(finalVals, vals)

  finalPos := posInArray(0, pos, valsLen)
  for i:=(valsLen-finalPos)+1 ; i>0 ; i-- {
    finalPos = posInArray(i, pos, valsLen)
    fmt.Println(i, finalPos)
    // swap
    temp := vals[i]
    vals[i] = vals[finalPos]
    vals[finalPos] = temp
    fmt.Println(">r",vals)
  }

  return vals
}

func main() {
  r := rotate([]int{1,2,3,4,5,6,7}, 3)
  fmt.Println(r)
}

/*


30 mins to hand write this.
7 mins to type this in.
10:57
exit status 2....
guessing i didn't copy the array and just started manipulating it...
out of index error...
using less than vs. less than equal...
and its wrong...
the end condistion is one of the elements at the end of the list....?
this im not sure about
trying to be too clever...
you have to go backwards to handle roations....

try this again tomorrow..
-but do the dumb solution...
a
a

*/