package main

import "fmt"

func nonRepeatChar(inputStr string) (bool) {
  il := len(inputStr)
  if il <= 0 {
    return true
  }

  // 97 is "a". 
  // 65 is "A"
  m := make(map[rune]bool)
  //fmt.Println(m[97])
  //m[97]=true
  //fmt.Println(m[97])
  for _, c := range inputStr {
    //fmt.Println("inserting", c, string(c))
    if m[c] == true {
      //fmt.Println("dupe found")
      return false
    } else {
      m[c] = true
    }
  }

  return true
}

func testRun(inputStr string) {
  result := nonRepeatChar(inputStr)
  fmt.Println(">", inputStr, result)
}

func main() {
  testRun("abcdefg")
  testRun("abcdefga")
}

/*
https://www.careercup.com/question?id=4569644446777344

not sure
02:16p

*/