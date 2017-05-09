package main

import "fmt"

// you probably want to do the reverse in place.
func reverseSent(src []string) ([]string) {
  if len(src) == 0 {
    return []string{""}
  } else if len(src) == 1 {
    return []string{src[0]}
  }

  // use i for startPos
  for i:=0 ; i<len(src) ; i++ {
    startPos := i
    endPos := len(src)-1  // default endPos to end of the list.
    
    // figure out the loc of the space. if can be the end of the list.  
    for j:=i+1 ; j<len(src) ; j++ {
      if src[j] == " " {
        endPos = j - 1  // back up 1 for the space.
        break
      }
    }

    // startPos and endPos
    //fmt.Println(startPos, endPos)

    // reverse the list.
    // you were swapping everything.
    midPos := endPos-((endPos-startPos)/2)
    //fmt.Println(">MP", midPos)
    j := startPos
    k := endPos
    for ; j<midPos ; {
      swap := src[j]
      //fmt.Println(">S", j, swap, endPos, endPos-j,k)
      src[j] = src[k]
      src[k] = swap
      
      // increment here to be consistent.
      // can't seem to have multiple increment in the loop construct.
      j++
      k--
    }

    // check the end pos.
    if endPos == len(src)-1 {
      break
    }

    // update positions.
    i = endPos + 1
  }

  return src
}

func reverseWord(src string) (string) {
  if len(src) == 0 {
    return ""
  } else if len(src) == 1 {
    return src
  }

  var dest string
  for i:=0 ; i<len(src) ; i++ {
    dest = dest + string(src[len(src)-(i+1)])
  }

  return dest
}



func main() {
  // s1 := "hello world"
  // r1 := reverseWord(s1)
  // fmt.Println(r1)

  s2 := []string{"h", "e", "l", "l", "o", " ", "w", "o", "r", "l", "d"}
  r2 := reverseSent(s2)
  fmt.Println(r2)

  s3 := []string{"h", "e", "l", " ", "l", "o", " ", "w", "o", "r", "l", "d"}
  r3 := reverseSent(s3)
  fmt.Println(r3)

  s4 := []string{"h", "e"}
  r4 := reverseSent(s4)
  fmt.Println(r4)

  s5 := []string{"h"}
  r5 := reverseSent(s5)
  fmt.Println(r5)
}

/*
  https://www.careercup.com/question?id=5697358784364544

  Reverse each work in a string.

  input
  I am a human

  output
  I ma a namuh

  probably a few ways to do this.
  figure out how t concat a character to a string.
*/