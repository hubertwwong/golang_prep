package main

import "fmt"

func deleteBOrAC(src []string) ([]string) {
  dest := make([]string, len(src))

  id := 0
  aState := false

  for _, v := range src {

    if aState == true {
      //fmt.Println("> aState is true")
      if v == "a" {
        //fmt.Println("> aState is true > a")
        // you see another a.
        // basically a no op.
      } else if v == "b" {
        //fmt.Println("> aState is true > b. a does not folloc c. add that.")
        aState = false
        dest[id] = "a"
        id++
      } else if v == "c" {
        //fmt.Println("> aState is true > c")
        // you see both A & C. RESET FLAG.
        aState = false  
      } else {
        //fmt.Println("> aState is true > else")
        // no AC.
        // you want to push both letters in the system.
        aState = false
        dest[id] = "a"
        id++
        dest[id] = v
        id++
      }
    } else if v != "b" && v != "a" {
      dest[id] = v
      id++
    } else if v == "b" {
      //fmt.Println("> b")
      // do nothing.
    } else if v == "a" {
      //fmt.Println("> a")
      aState = true
      // do nothing.
    }

    //fmt.Println(i, "", v)
  }

  // slice off the end.
  // We allocated the full size of the array initially.
  // This does a cleanup.
  dest = dest[0:id]

  return dest
}

func main() {
  inputS1 := []string{"z", "a", "c", "d", "e"}
  outputS1 := deleteBOrAC(inputS1)
  fmt.Println(inputS1, outputS1, len(outputS1))

  //inputS2 := []string{"z", "a", "c", "d", "e"}
  outputS2 := deleteBOrAC(nil)
  fmt.Println(outputS2, len(outputS2))

  inputS3 := []string{"a", "c"}
  outputS3 := deleteBOrAC(inputS3)
  fmt.Println(inputS3, outputS3, len(outputS3))

  inputS4 := []string{"a", "c", "f"}
  outputS4 := deleteBOrAC(inputS4)
  fmt.Println(inputS4, outputS4, len(outputS4))

  inputS5 := []string{"a", "b", "c", "d", "e"}
  outputS5 := deleteBOrAC(inputS5)
  fmt.Println(inputS5, outputS5, len(outputS5))

  inputS6 := []string{"a", "d", "e", "f", "g"}
  outputS6 := deleteBOrAC(inputS6)
  fmt.Println(inputS6, outputS6, len(outputS6))

  inputS7 := []string{"f", "c", "a", "b", "e"}
  outputS7 := deleteBOrAC(inputS7)
  fmt.Println(inputS7, outputS7, len(outputS7))
}

/*
  https://www.careercup.com/question?id=18460667

  Eliminate all ‘b’ and ‘ac’ in an array of characters,
  you have to replace them in-place,
  and you are only allowed to iterate over the char array once.

  e.g:

  1: abc -> ac
  2: ac->''
  3: react->rt
*/
