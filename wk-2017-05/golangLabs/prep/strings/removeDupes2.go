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

// return a list of permutations.
// assuming its sorted
func perm(vals []int) [][]int {
  timeTrack(time.Now(), "PermStart") 
  
  // base case
  if len(vals) <= 1 {
    res := make([][]int, 0)
    res = append(res, vals)
    return res
  }

  result := make([][]int, 0)
  for i := 0 ; i < len(vals) ; i++ {
    // construct the subsequence to pass in to the recursion.
    curSubSequence := make([]int, 0)
    if i+1 == len(vals) {
      curSubSequence = vals[:i]
    } else if i == 0 {
      curSubSequence = vals[i+1:]
    } else {
      left := vals[0:i]
      right := vals[i+1:]
      //  append stuff together
      for i := 0 ; i < len(left) ; i++ {
        curSubSequence = append(curSubSequence, left[i])
      }
      for i := 0 ; i < len(right) ; i++ {
        curSubSequence = append(curSubSequence, right[i])
      }
    }
    
    // passing in curSubSequence
    resA := perm(curSubSequence)

    // append on result to mega results
    for j := 0 ; j < len(resA) ; j++ {
      curRow := make([]int, 0)
      curRow = append(curRow, vals[i])          // pass in the sliced off value
      curRow = append(curRow, resA[j]...)       // pass in recursion result
      result = append(result, curRow)
    }
  }
  timeTrack(time.Now(), "PermEnd") 
  return result
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
  possPerms := perm(possChars)

  // compute the permutations. figure out which one is valid.
  permNo := 0
  for ; permNo < len(possPerms) ; permNo++ {
    stringI := 0
    numCharMatched := 0
    fmt.Println("\non perm", possPerms[permNo])

    // scaning a specific permutation.
    for i := 0 ; i < len(possPerms[permNo]) ; i++ {
      curCharOnPerm := possPerms[permNo][i]
      fmt.Println("curCharOnPerm", curCharOnPerm)

      // scan the allCharA
      for j := 0 ; j < len(allCharA[curCharOnPerm]) ; j++ {
        curAllChar := allCharA[curCharOnPerm][j]
        fmt.Println("matching > j", j, "c", allCharA[curCharOnPerm][j], "string I", stringI)
        if curAllChar >= stringI {
          fmt.Println("match found")
          stringI = curAllChar
          numCharMatched++
          break
        }
      } 
    }

    // if the number of characters match is the lenght of the permutation,
    // you are done. you found the min permutation.
    if numCharMatched == len(possPerms[0]) {
      break
    }
  }

  fmt.Println("> allCharA")
  printInt(allCharA)

  // convert int to letters.
  fmt.Println("> permutations")
  printInt(possPerms)
  
  fmt.Println("> result >>>>>>>>>>")
  if permNo < len(possPerms) {
    fmt.Println(permNo, possPerms[permNo])
  } else {
    fmt.Println("> out of range", permNo)
  }

  resultAsStr := make([]byte, 0)
  for i := 0 ; i < len(possPerms[permNo]) ; i++ {
    resultAsStr = append(resultAsStr, byte(possPerms[permNo][i]) + "a"[0])
  }

  return string(resultAsStr)
}




func main() {
  var s string
  //s = "cbacdcbc"
  //s = "abc"
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

  This works in small cases...
  The permutation calucation that a huge amount of time and its causing a TLE error.....

  https://leetcode.com/problems/remove-duplicate-letters/#/description
*/