package main

import "fmt"

func isNonRepeat(myStr string) (bool) {
  myStrLen := len(myStr)

  // or return some error.
  if myStrLen <= 0 {
    return true
  }

  // technically, rune is just int32
  m := make(map[rune]bool)

  for _, c := range myStr {
    // shoving it in that range...
    var mapI rune
    mapI = -1

    if c >= 65 && c <= 91 {
      mapI = c-65
    } else if c >= 97 && c <= 123 {
      mapI = c-97
    }

    // checks for mapI was set before makes any other check.
    if mapI != -1 {
      if m[mapI] == true {
        return false
      } else {
        m[mapI] = true
      }
    }
  }

  return true
}

func testRun(myStr string) {
  result := isNonRepeat(myStr)
  fmt.Println(">", myStr, result)
}

func main() {
  testRun("hello")
  testRun("helo")
  testRun("helo 123123123231132321321")
  testRun("Heloh")
  testRun("")
}

/*
  mixing runes and ints using the lazy init.
  you need to keep types consistent.
*/