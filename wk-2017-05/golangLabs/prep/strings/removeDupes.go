package main

import "fmt"
import "log"
import "time"

func printInt(vals [][]int) {
   for i := 0 ; i < len(vals) ; i++ {
       fmt.Println(vals[i])
   } 
}

// for debug...
func printBool(vals [][]bool) {
   for i := 0 ; i < len(vals) ; i++ {
       fmt.Println(vals[i])
   } 
}

func timeTrack(start time.Time, name string) {
  elapsed := time.Since(start)
  log.Printf("%s took %s", name, elapsed)
}

// returns the next permutation.
func permNext(vals []int) ([]int, bool) {
  // Find longest non-increasing suffix
  i := len(vals) - 1
  for ; i > 0 && vals[i - 1] >= vals[i] ; i-- {}
      
  // Now i is the head index of the suffix
  
  // Are we at the last permutation already?
  if i <= 0 { 
    return nil, false
  }

  // Let vals[i - 1] be the pivot
  // Find rightmost element that exceeds the pivot
  j := len(vals) - 1
  for ; vals[j] <= vals[i - 1] ; j-- {}

  // Now the value vals[j] will become the new pivot
  // Assertion: j >= i
  
  // Swap the pivot with j
  vals[i - 1], vals[j] = vals[j], vals[i - 1]
  
  // Reverse the suffix
  j = len(vals) - 1
  for ; i < j ; {
    vals[i], vals[j] = vals[j], vals[i]
    i++
    j--
  }
  
  // Successfully computed the next permutation
  return vals, true
}

// helper function that returns the first position that isn't -1
// or -1 if its false
// using the golang idion.
func firstPos(vals []int, offset int) (int, bool) {
    for i := offset ; i < len(vals) ; i++ {
        if vals[i] != -1 {
            return i, true
        }
    }
    return -1, false
}

func removeDuplicateLetters(s string) string {
  // guard
  if len(s) <= 1 {
      return s
  }
  
  // array to store first chars
  allCharA := make([][]int, 26)
  
  // 1st pass.
  // 1. count first possible position of the character.
  // 2. count all possible position of the character.
  for i := 0 ; i < len(s) ; i++ {
    charPosInArray := s[i] - "a"[0]
    
    // store the poisition of the character in array.
    // this is per letter.
    // e.g.
    // a is 1 -> 5 -> 9
    // b is 2 -> 3
    // and so on...
    if len(allCharA[charPosInArray]) == 0 {
      allCharA[charPosInArray] = make([]int, 0)
    }
    allCharA[charPosInArray] = append(allCharA[charPosInArray], i)
  }
  
  // construct a list of possible characters.
  possChars := make([]int, 0)
  for i := 0 ; i < len(allCharA) ; i++ {
    if len(allCharA[i]) != 0 {
      possChars = append(possChars, i)
    }
  }
  
  // compute the permutations. figure out which one is valid.
  // using possChars to store each permutation.
  // using permNextTo calculate the next permutation.
  for nextPermExist := true ; nextPermExist ; possChars, nextPermExist = permNext(possChars) {
    stringI := 0
    numCharMatched := 0
    //fmt.Println("\non perm", possChars)

    // scaning a specific permutation.
    for i := 0 ; i < len(possChars) ; i++ {
      curCharOnPerm := possChars[i]
      //fmt.Println("curCharOnPerm", curCharOnPerm)

      // scan the allCharA
      for j := 0 ; j < len(allCharA[curCharOnPerm]) ; j++ {
        curAllChar := allCharA[curCharOnPerm][j]
        //fmt.Println("matching > j", j, "c", allCharA[curCharOnPerm][j], "string I", stringI)
        if curAllChar >= stringI {
          //fmt.Println("match found")
          stringI = curAllChar
          numCharMatched++
          break
        }
      } 
    }

    // if the number of characters match is the lenght of the permutation,
    // you are done. you found the min permutation.
    if numCharMatched == len(possChars) {
      break
    }
  }

  //fmt.Println("> allCharA")
  //printInt(allCharA)

  resultAsStr := make([]byte, 0)
  for i := 0 ; i < len(possChars) ; i++ {
    resultAsStr = append(resultAsStr, byte(possChars[i]) + "a"[0])
  }

  return string(resultAsStr)
}



func main() {
  fmt.Println(">>>>>")
  var s string
  //s = "cbacdcbc"
  s = "baab"
  //s = "cbacdcb"
  s = "thesqtitxyetpxloeevdeqifkz"
  timeTrack(time.Now(), "Main") 
  res := removeDuplicateLetters(s)
  timeTrack(time.Now(), "MainEnd") 
  fmt.Println(res)
  //printInt(res)
}

/*
  remove duplicate letters from a string.
  return them in a lex minimal ordering.

  Trying an optimization so you are not calcuating every perutation before you start.
  Permutation are n! calcuation so you probably want to avoid this.

  need to cut down of the space of the problem.
  or come up with a better solution

  https://leetcode.com/problems/remove-duplicate-letters/#/description
*/