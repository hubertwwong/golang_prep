package main

import "fmt"

func lcsNaive(s, t string) string {
  res := make([]byte, 0)

  curTi := 0
  for si := 0 ; si < len(s) ; si++ {
    for ti := curTi ; ti < len(t) ; ti++ {
      if s[si] == t[ti] {
        //fmt.Println(si, ti)
        res = append(res, s[si])
        curTi = ti+1
        break
      }
    }
  }

  return string(res)
}

func main() {
  s1 := ""
  s2 := ""

  s1 = "bbbbbb"
  s2 = "aaaaaa"
  fmt.Println(s1, s2, "> lcs > [",lcsNaive(s1, s2), "]")
}