package main

import "fmt"
import "strings"

func extract(s string, pos int) string {
  if pos < 0 || pos > len(s) {
    return ""
  }

  result := strings.Join([]string{s[:pos], s[(pos+1):]}, "")
  return result
}

func main() {
  fmt.Println(extract("Heloba", 2))
}