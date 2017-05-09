package main

import (
  "fmt"
  "math"
)

func rabinKarp(longS, shortS string) int {
  // guards
  if len(shortS) > len(longS) {
    return -1
  } else if len(shortS) == 0 || len(longS) == 0 {
    return -1
  }

  prime := 101
  shortHash := 0
  longHash := 0

  // initial has computation
  for i:=0 ; i<len(shortS) ; i++ {
    //fmt.Println(i, int(shortS[i]), int(longS[i]), int(math.Pow(float64(prime), float64(i))))
    shortHash = shortHash + int(shortS[i]) * int(math.Pow(float64(prime), float64(i)))
    longHash = longHash + int(longS[i]) * int(math.Pow(float64(prime), float64(i)))
    //fmt.Println(shortHash, longHash)
  }
  if shortHash == longHash && subStrMatch(longS, shortS, 0) {
    return 0
  }

  // main loop
  for i:=1 ; i+len(shortS)-1<len(longS) ; i++ {
    // debug
    
    // update hash
    longHash = longHash - int(longS[i-1])
    longHash = longHash / prime
    longHash = longHash + int(longS[i+len(shortS)-1]) * int(math.Pow(float64(prime), float64(len(shortS)-1)))
    //fmt.Println("lh update", longHash)

    // check
    if shortHash == longHash && subStrMatch(longS, shortS, i) {
      return i
    }
  }

  return -1
}

func subStrMatch(longS, shortS string, pos int) bool {
  if len(shortS) > len(longS) {
    return false
  } else if pos < 0 || pos + len(shortS) > len(longS) {
    return false
  }

  // this is wrong...
  for i,j:=pos,0 ; i<pos+len(shortS) ; i,j = i+1, j+1 {
    if shortS[j] != longS[i] {
      return false
    }
  }

  return true
}

func main() {
  fmt.Println(rabinKarp("lorem ipsum is the best thing ever", "is "))
}


/*
mostly good...
the algo is off...
types have to match
math.Pow is from the math lib.

so substringMatch has an error...
on the update func... the wrong i is being used....
the big thing is what i did the offset off...
when you len something, you want  the last character....

convert everything to ints.

*/