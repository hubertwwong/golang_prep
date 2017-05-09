package main

import "fmt"

func oneCharDiff(srcStr string, list []string) (bool) {
  // guards
  if len(srcStr)==0 {
    return false
  } else if list == nil || len(list) == 0 {
    return false
  }

  // Compute number of items that need to match.
  oneCharDiff := len(srcStr) - 1
  // loop for the list.
  for i:=0 ; i<len(list) ; i++ {
    //fmt.Println(list[i])
    // if the string length don't match, go on to the next item.
    if len(list[i]) != len(srcStr) {
      continue
    }

    // loop thru each string in the list.
    curNumMatch := 0
    for j:=0 ; j<len(srcStr) ; j++ {
      if srcStr[j] == list[i][j] {
        curNumMatch++
      }
    }

    // if the count is 1 off from the length from the srcStr,
    // return true. we are done.
    if curNumMatch == oneCharDiff {
      return true
    }
  }

  return false
}



func main() {
  l1 := []string{"bana", "apple", "banaba", "bonaza", "banamf"}
  s1 := "banana"
  r1 := oneCharDiff(s1, l1)
  fmt.Println("\n>problem example")
  fmt.Println(r1)

  l2 := []string{"bar", "baz"}
  s2 := "zzz"
  r2 := oneCharDiff(s2, l2)
  fmt.Println("\n>non matching example")
  fmt.Println(r2)

  l3 := []string{"bar", "baz"}
  s3 := ""
  r3 := oneCharDiff(s3, l3)
  fmt.Println("\n>empty string")
  fmt.Println(r3)

  //l4 := []string{"bar", "baz"}
  s4 := "zzz"
  r4 := oneCharDiff(s4, nil)
  fmt.Println("\n>nil array")
  fmt.Println(r4)

  l5 := []string{"a", "b", "c"}
  s5 := "z"
  r5 := oneCharDiff(s5, l5)
  fmt.Println("\n>one character for everything")
  fmt.Println(r5)

  l6 := []string{"caa", "zaa"}
  s6 := "aaa"
  r6 := oneCharDiff(s6, l6)
  fmt.Println("\n>another matching example")
  fmt.Println(r6)
}



/*

https://www.careercup.com/question?id=5760697411567616

Given a string an array
Figure out if there is a word in the array that is a single character off from the string.
Can the letter be blank?
Exactly 1..

trees...
- probably want to think on this....
- figure out what to do....

e.g.

banana

[bana, apple, banaba, bonaza, banamf]



*/