package main

import (
  "fmt"
  "strconv"
)

func factorial(x int) (int) {
  if x == 0 {
    return 1
  }
  return x*factorial(x-1)
}

func factSum(x int) (int) {
  if x == 0 {
    return 1
  }
  return x+factSum(x-1)
}

func printLists(src [][]string) {
  for i:=0 ; i<len(src) ; i++ {
    fmt.Println(src[i])
  }
}

func strToStrArray(srcStr string) ([]string) {
  if len(srcStr) == 0 {
    return []string{""}
  }

  result := make([]string, len(srcStr))

  for i:=0 ; i<len(srcStr) ; i++ {
    cChar := string(srcStr[i])
    result[i] = cChar  
  }

  return result
}

// assumes 3 characters?
func i18nStr(srcStr string, startPos int, endPos int) ([]string) {
  if len(srcStr) < 3 {
    return []string{""}
  } else if startPos > len(srcStr) || startPos < 0 {
    return []string{""}
  } else if endPos > len(srcStr) || endPos < 0 {
    return []string{""}
  }
  // } else if endPos <= startPos {
  //   return []string{""}
  // }

  newStrLen := len(srcStr) - (endPos-startPos)
  result := make([]string, newStrLen)
  numDiff := 1 + endPos - startPos
  numDiffStr := strconv.Itoa(numDiff)

  for i, cPos:=0, 0 ; i<len(srcStr) ; i++ {
    cChar := string(srcStr[i])
    //fmt.Println(cChar, ">i", i, ">cc", cChar, ">cp",cPos)
    if i == startPos {
      result[cPos] = numDiffStr
      i = (i+numDiff)-1
    } else {
      result[cPos] = cChar
    }

    cPos++
  }

  return result
}

func i18n(srcStr string) ([][]string) {
  // guards
  if len(srcStr) == 0 {
    // returning an empty string
    return [][]string{{""}}
  } else if len(srcStr) < 3 {
    // if its less than 3 characters.
    // there is nothing to iterate.
    // just stash the src string.
    return [][]string{{srcStr}}
  }

  finalLen := factSum(len(srcStr)-2)
  //fmt.Println(">fl", finalLen, len(srcStr)-2)
  results := make([][]string, finalLen)
  curResultI := 0
  
  // outer loop controls number of characters.
  // starting at 1. going up to the length of string -2
  // but since we are starting at 1 we remove 1 at the check.
  for i:=1 ; i<len(srcStr)-1 ; i++ {
    numBetween := len(srcStr)-(2+i)
    //curArrayLen := len(srcStr) - (numBetween+1)
    //fmt.Println(numBetween, curArrayLen)

    // num of internal number to iterate.
    // 2 index j. and curNewPos.
    // j is the source index.
    // curNewPos is the string.
    for startPos:=1 ; startPos<len(srcStr)-(numBetween+1) ; startPos++ {
      endPos := startPos + numBetween
      //fmt.Println(">i", i, ">sp", startPos, ">e", endPos, ">cri", curResultI)
      results[curResultI] = i18nStr(srcStr, startPos, endPos)
      curResultI++
    }
  }

  results[curResultI] = strToStrArray(srcStr)
  return results
}



func main() {
  // s1 := "abcde"
  // r1 := i18nStr(s1, 1, 1)
  // fmt.Println(r1)
  // printLists(r1)

  // s2 := "abcde"
  // r2 := i18n(s2)
  // printLists(r2)

  s3 := "careercup"
  r3 := i18n(s3)
  printLists(r3)

  s4 := "cz"
  r4 := i18n(s4)
  printLists(r4)
}

/*

https://www.careercup.com/question?id=5733696185303040

Generate i18n strings

input:
careercup

output:
c7p
ca6p
c6ap
car6p

I mis understood the problem.
Its literall adding the number.

1st one
7 is numBetween
1 is numPosition.

2nd one
6 is numBetween
2 is numPositions
9-(6+1)

You have to compute what to delete.
in 1 case.

compute
- start pos. (this is by the j variable)
- end pos. (j+numBetween)
j=0

remember to offset by 1. you always start 1

startPos=0
endPos=7
exit condition is the last character.

if it does not hit the start pos, insert the character.
if it does hit the start post. insert the number.
skip to the endposition.

startPos=0
endPos=6

I think this is the algo. or close to it.
init the array. insert first and last characters.
length = len - numbetween+1.
startPos=1
endPost=7
j=0. insert 1st character
j=1. hits the start pos. insert number.
set j=7
j gets incremented.
j<len.


*/