package main

import "fmt"

func mask1s() (uint64) {
  var x uint64
  for i:=0 ; i<64 ; i++ {
    x = x << 1
    x = x +1
  }
  return x
}

func clearMask(startPos, endPos int) (uint64) {
  var x uint64
  for i:=0 ; i<64 ; i++ {
    // fmt.Println((64-endPos))
    // fmt.Println((64-startPos))
    // fmt.Println(">",i)
    if i >= (64-endPos) && i <= (64-startPos) {
      x = x << 1
      //x = x + 0
    } else {
      x = x << 1
      x = x + 1
    }
  }
  return x
}

func setMask(val uint64, startPos uint) (uint64) {
  return val << startPos
}

func printBits(x uint64) {
  var m uint64
  m = 1
  m = m << 63

  for i,j:= 0, uint(63) ; i<64 ; i,j = i+1, j-1 {
    myBit := x&m
    myBit = myBit >> j
    fmt.Print(myBit)
    m = m >> 1
  }
  fmt.Println("")

  for i:=0 ; i<64 ; i++ {
    fmt.Print((64-i)%10)
  }
  fmt.Println("")
}

func insertMIntoN(m, n uint64, startPos, endPos int) (uint64) {
  finalN := n
  //fmt.Println(">finalN")
  printBits(finalN)
  
  cm := clearMask(startPos, endPos)
  //fmt.Println(">CM")
  printBits(cm)
  
  finalN &= cm
  //fmt.Println(">AFTERCM")
  printBits(finalN)
  
  sm := setMask(m, uint(startPos))
  finalN |= sm
  
  return finalN
}

func main() {
  // x1 := uint64(8)
  // printBits(x1)

  // fmt.Println(">>>>")

  // x2 := clearMask(3,5)
  // printBits(x2)

  // fmt.Println(">>>>")

  // x3 := setMask(x1, 1)
  // printBits(x3)

  x4 := insertMIntoN(uint64(19), uint64(1024), 2, 6)
  fmt.Println(">FINAL")
  printBits(x4)
}

/*
This problem was done with compiler..
bit take aways...
1 you have to be super careful on the iterators...
units are need for shifting. golang will coerce the raw values for you...
you can shift and add the values...
to get a mask of 1.
if you shift 0, you can more zeros. this is an easy way to get zero mask...

didn't read the question
they actually specified 32 bits

indicies seem to be inclusive...


*/